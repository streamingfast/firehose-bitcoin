package poller

import (
	"context"
	"fmt"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/streamingfast/bstream"
	pbbstream "github.com/streamingfast/bstream/pb/sf/bstream/v1"
	pbbitcoin "github.com/streamingfast/firehose-bitcoin/pb/sf/bitcoin/type/v1"
	"github.com/streamingfast/firehose-core/blockpoller"
	"github.com/streamingfast/shutter"
	"go.uber.org/zap"
)

type Reader struct {
	*shutter.Shutter
	endpoint             string
	blockFetchRetryCount uint64
	stateStoragePath     string
	startBlockNum        uint64
	rpcClient            *rpcclient.Client
	logger               *zap.Logger
}

func New(endpoint string, blockFetchRetryCount uint64, stateStoragePath string, startBlockNum uint64, logger *zap.Logger) *Reader {
	return &Reader{
		Shutter:              shutter.New(),
		endpoint:             endpoint,
		blockFetchRetryCount: blockFetchRetryCount,
		stateStoragePath:     stateStoragePath,
		startBlockNum:        startBlockNum,
		logger:               logger.Named("poller"),
	}
}

func (r *Reader) Run(ctx context.Context) error {
	contentType := getContentType()
	r.logger.Info("launching firebtc reader",
		zap.String("endpoint", r.endpoint),
		zap.String("content", contentType),
		zap.Uint64("start_block_num", r.startBlockNum),
	)

	bp := blockpoller.New(r, blockpoller.NewFireBlockHandler(contentType),
		blockpoller.WithLogger(r.logger),
		blockpoller.WithBlockFetchRetryCount(r.blockFetchRetryCount),
		blockpoller.WithStoringState(r.stateStoragePath),
	)
	r.OnTerminating(func(err error) {
		r.logger.Info("shutting down firebtc reader", zap.Error(err))
		bp.Shutdown(nil)
	})

	connCfg := &rpcclient.ConnConfig{
		Host:         r.endpoint,
		DisableAuth:  true,
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return fmt.Errorf("failed to setup rpc client: %w", err)
	}
	defer client.Shutdown()
	r.rpcClient = client

	finalizedBlk, err := r.GetFinalizedBlock()
	if err != nil {
		return fmt.Errorf("failed to get finalized block: %w", err)
	}
	r.logger.Info("retrieved fianlized block", zap.Stringer("finalized_block", finalizedBlk))

	return bp.Run(ctx, r.startBlockNum, finalizedBlk)
}

func (r *Reader) GetFinalizedBlock() (bstream.BlockRef, error) {

	bestBlockHash, err := r.rpcClient.GetBestBlockHash()
	if err != nil {
		return nil, fmt.Errorf("unable to get best block hash: %w", err)
	}
	r.logger.Info("found best block hash", zap.Stringer("blockhash", bestBlockHash))

	bestBlock, err := r.rpcClient.GetBlockVerbose(bestBlockHash)
	if err != nil {
		return nil, fmt.Errorf("unable to get best block %s: %w", bestBlockHash.String(), err)
	}

	return bstream.NewBlockRef(bestBlockHash.String(), uint64(bestBlock.Height)), nil
}

func (r *Reader) Fetch(_ context.Context, blkNum uint64) (*pbbstream.Block, error) {
	r.logger.Debug("fetching block", zap.Uint64("block_num", blkNum))
	blkHash, err := r.rpcClient.GetBlockHash(int64(blkNum))
	if err != nil {
		return nil, fmt.Errorf("unable to get block hash for block %d: %w", blkNum, err)
	}

	rpcBlk, err := r.rpcClient.GetBlockVerboseTx(blkHash)
	if err != nil {
		return nil, fmt.Errorf("unable to get block %d (%s): %w", blkNum, blkHash.String(), err)
	}

	r.logger.Debug("found block",
		zap.Int64("block_num", rpcBlk.Height),
		zap.String("block_hash", rpcBlk.Hash),
		zap.String("prev_hash", rpcBlk.PreviousHash),
	)

	blk := &pbbitcoin.Block{
		Hash:         blkHash.String(),
		Size:         rpcBlk.Size,
		StrippedSize: rpcBlk.StrippedSize,
		Weight:       rpcBlk.Weight,
		Height:       rpcBlk.Height,
		Version:      rpcBlk.Version,
		VersionHex:   rpcBlk.VersionHex,
		MerkleRoot:   rpcBlk.MerkleRoot,
		Tx:           nil,
		Time:         rpcBlk.Time,
		// TODO: we need to solve this
		//Mediantime:   0,
		Nonce:      rpcBlk.Nonce,
		Bits:       rpcBlk.Bits,
		Difficulty: rpcBlk.Difficulty,
		// TODO: we need to solve this
		//Chainwork:    rpcBlk,
		NTx:          uint32(len(rpcBlk.Tx)),
		PreviousHash: rpcBlk.PreviousHash,
	}

	for _, tx := range rpcBlk.Tx {
		trx := &pbbitcoin.Transaction{
			Hex:       tx.Hex,
			Txid:      tx.Txid,
			Hash:      tx.Hash,
			Size:      tx.Size,
			Vsize:     tx.Vsize,
			Weight:    tx.Weight,
			Version:   tx.Version,
			Locktime:  tx.LockTime,
			Vin:       nil,
			Vout:      nil,
			Blockhash: tx.BlockHash,
			Blocktime: tx.Blocktime,
		}

		for _, vin := range tx.Vin {
			pbvin := &pbbitcoin.Vin{
				Txid:        vin.Txid,
				Vout:        vin.Vout,
				Sequence:    vin.Sequence,
				Txinwitness: vin.Witness,
				Coinbase:    vin.Coinbase,
			}

			if vin.ScriptSig != nil {
				pbvin.ScriptSig = &pbbitcoin.ScriptSig{
					Asm: vin.ScriptSig.Asm,
					Hex: vin.ScriptSig.Hex,
				}
			}

			trx.Vin = append(trx.Vin)

		}

		for _, vout := range tx.Vout {
			trx.Vout = append(trx.Vout, &pbbitcoin.Vout{
				Value: vout.Value,
				N:     vout.N,
				ScriptPubKey: &pbbitcoin.ScriptPubKey{
					Asm:       vout.ScriptPubKey.Asm,
					Hex:       vout.ScriptPubKey.Hex,
					ReqSigs:   vout.ScriptPubKey.ReqSigs,
					Type:      vout.ScriptPubKey.Type,
					Addresses: vout.ScriptPubKey.Addresses,
				},
			})
		}
		blk.Tx = append(blk.Tx, trx)
	}

	return blk.MustToBstreamBlock(), nil
}

func getContentType() string {
	blk := &pbbitcoin.Block{}
	return string(blk.ProtoReflect().Descriptor().FullName())
}
