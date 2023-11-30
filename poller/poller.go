package poller

import (
	"context"
	"fmt"
	"time"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/streamingfast/bstream"
	pbbstream "github.com/streamingfast/bstream/pb/sf/bstream/v1"
	pbbitcoin "github.com/streamingfast/firehose-bitcoin/pb/sf/bitcoin/type/v1"
	"github.com/streamingfast/firehose-core/blockpoller"
	"github.com/streamingfast/shutter"
	"go.uber.org/zap"
)

type Poller struct {
	*shutter.Shutter
	endpoint             string
	blockFetchRetryCount uint64
	stateStoragePath     string
	startBlockNum        uint64
	ignoreCursor         bool
	blockInterval        time.Duration
	headBlock            bstream.BlockRef
	rpcClient            *rpcclient.Client
	logger               *zap.Logger
}

func New(endpoint string, blockFetchRetryCount uint64, stateStoragePath string, startBlockNum uint64, ignoreCursor bool, logger *zap.Logger) *Poller {
	return &Poller{
		Shutter:              shutter.New(),
		endpoint:             endpoint,
		blockFetchRetryCount: blockFetchRetryCount,
		stateStoragePath:     stateStoragePath,
		startBlockNum:        startBlockNum,
		ignoreCursor:         ignoreCursor,
		blockInterval:        7 * time.Minute,
		logger:               logger.Named("poller"),
	}
}

func (p *Poller) Run(ctx context.Context) error {
	contentType := getContentType()
	p.logger.Info("launching firebtc poller",
		zap.String("endpoint", p.endpoint),
		zap.String("content", contentType),
		zap.Uint64("start_block_num", p.startBlockNum),
	)

	opts := []blockpoller.Option{
		blockpoller.WithLogger(p.logger),
		blockpoller.WithBlockFetchRetryCount(p.blockFetchRetryCount),
		blockpoller.WithStoringState(p.stateStoragePath),
	}

	if p.ignoreCursor {
		opts = append(opts, blockpoller.IgnoreCursor())
	}

	bp := blockpoller.New(p, blockpoller.NewFireBlockHandler(contentType), opts...)
	p.OnTerminating(func(err error) {
		p.logger.Info("shutting down firebtc reader", zap.Error(err))
		bp.Shutdown(nil)
	})

	connCfg := &rpcclient.ConnConfig{
		Host:         p.endpoint,
		DisableAuth:  true,
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		return fmt.Errorf("failed to setup rpc client: %w", err)
	}
	defer client.Shutdown()
	p.rpcClient = client

	p.headBlock, err = p.GetHeadBlock()
	if err != nil {
		return fmt.Errorf("failed to get head block: %w", err)
	}

	finalizedBlk, err := p.GetFinalizedBlock(p.headBlock)
	if err != nil {
		return fmt.Errorf("failed to get finalized block: %w", err)
	}
	p.logger.Info("retrieved finalized and head block",
		zap.Stringer("finalized_block", finalizedBlk),
		zap.Stringer("head_block", p.headBlock),
	)

	return bp.Run(ctx, p.startBlockNum, finalizedBlk)
}

func (p *Poller) GetFinalizedBlock(headBlock bstream.BlockRef) (bstream.BlockRef, error) {
	finalizedBlockNum := int64(headBlock.Num() - pbbitcoin.LibOffset)
	finalizedBlockHash, err := p.rpcClient.GetBlockHash(finalizedBlockNum)
	if err != nil {
		return nil, fmt.Errorf("failed to get finalized block hash: %w", err)
	}

	return bstream.NewBlockRef(finalizedBlockHash.String(), uint64(finalizedBlockNum)), nil
}

func (p *Poller) GetHeadBlock() (bstream.BlockRef, error) {
	bestBlockHash, err := p.rpcClient.GetBestBlockHash()
	if err != nil {
		return nil, fmt.Errorf("unable to get best block hash: %w", err)
	}
	p.logger.Info("found best block hash", zap.Stringer("blockhash", bestBlockHash))

	bestBlock, err := p.rpcClient.GetBlockVerbose(bestBlockHash)
	if err != nil {
		return nil, fmt.Errorf("unable to get best block %s: %w", bestBlockHash.String(), err)
	}

	return bstream.NewBlockRef(bestBlockHash.String(), uint64(bestBlock.Height)), nil
}

func (p *Poller) Fetch(_ context.Context, blkNum uint64) (*pbbstream.Block, error) {
	for p.headBlock.Num() < blkNum {
		var err error
		p.headBlock, err = p.GetHeadBlock()
		if err != nil {
			return nil, fmt.Errorf("failed to get head block: %w", err)
		}

		if p.headBlock.Num() < blkNum {
			p.logger.Info("head block is behind, waiting for it to catch up",
				zap.Stringer("head_block", p.headBlock),
				zap.Uint64("block_num", blkNum),
				zap.Duration("wait_duration", p.blockInterval),
			)
			time.Sleep(p.blockInterval)
			continue
		}
		break
	}

	p.logger.Debug("fetching block", zap.Uint64("block_num", blkNum))
	blkHash, err := p.rpcClient.GetBlockHash(int64(blkNum))
	if err != nil {
		return nil, fmt.Errorf("unable to get block hash for block %d: %w", blkNum, err)
	}

	rpcBlk, err := p.rpcClient.GetBlockVerboseTx(blkHash)
	if err != nil {
		return nil, fmt.Errorf("unable to get block %d (%s): %w", blkNum, blkHash.String(), err)
	}

	p.logger.Debug("found block",
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

			trx.Vin = append(trx.Vin, pbvin)

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
