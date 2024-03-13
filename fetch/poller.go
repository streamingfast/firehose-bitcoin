package fetch

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

type BlockFetcher struct {
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

	headers    map[string]string
	disableTLS bool
}

func New(
	endpoint string,
	https bool,
	blockFetchRetryCount uint64,
	stateStoragePath string,
	startBlockNum uint64,
	ignoreCursor bool,
	headers map[string]string,
	logger *zap.Logger) *BlockFetcher {
	return &BlockFetcher{
		Shutter:              shutter.New(),
		endpoint:             endpoint,
		blockFetchRetryCount: blockFetchRetryCount,
		stateStoragePath:     stateStoragePath,
		startBlockNum:        startBlockNum,
		ignoreCursor:         ignoreCursor,
		blockInterval:        7 * time.Minute,
		disableTLS:           !https,
		headers:              headers,
		logger:               logger.Named("poller"),
	}
}

func (p *BlockFetcher) Run(ctx context.Context) error {
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
		DisableTLS:   p.disableTLS,
		ExtraHeaders: p.headers,
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
		zap.Stringer("chain_head_block", p.headBlock),
	)

	return bp.Run(ctx, p.startBlockNum, 1)
}

func (p *BlockFetcher) GetFinalizedBlock(headBlock bstream.BlockRef) (bstream.BlockRef, error) {
	finalizedBlockNum := int64(headBlock.Num() - pbbitcoin.LibOffset)
	finalizedBlockHash, err := p.rpcClient.GetBlockHash(finalizedBlockNum)
	if err != nil {
		return nil, fmt.Errorf("failed to get finalized block hash: %w", err)
	}

	return bstream.NewBlockRef(finalizedBlockHash.String(), uint64(finalizedBlockNum)), nil
}

func (p *BlockFetcher) GetHeadBlock() (bstream.BlockRef, error) {
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

func (p *BlockFetcher) IsBlockAvailable(requestedSlot uint64) bool {
	return requestedSlot <= p.headBlock.Num()
}

func (p *BlockFetcher) Fetch(_ context.Context, blkNum uint64) (*pbbstream.Block, bool, error) {
	for p.headBlock.Num() < blkNum {
		headblock, err := p.GetHeadBlock()
		if err != nil {
			return nil, false, fmt.Errorf("failed to get head block: %w", err)
		}

		p.headBlock = headblock
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
	t0 := time.Now()
	blkHash, err := p.rpcClient.GetBlockHash(int64(blkNum))
	if err != nil {
		return nil, false, fmt.Errorf("unable to get block hash for block %d: %w", blkNum, err)
	}
	duration := time.Since(t0)

	t1 := time.Now()
	rpcBlk, err := p.rpcClient.GetBlockVerboseTx(blkHash)
	if err != nil {
		return nil, false, fmt.Errorf("unable to get block %d (%s): %w", blkNum, blkHash.String(), err)
	}

	p.logger.Debug("found block",
		zap.Int64("block_num", rpcBlk.Height),
		zap.String("block_hash", rpcBlk.Hash),
		zap.String("prev_hash", rpcBlk.PreviousHash),
		zap.Duration("blockhash_duration", duration),
		zap.Duration("block_duration", time.Since(t1)),
		zap.Duration("total_duration", time.Since(t0)),
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

	return blk.MustToBstreamBlock(), false, nil
}

func getContentType() string {
	blk := &pbbitcoin.Block{}
	return "type.googleapis.com/" + string(blk.ProtoReflect().Descriptor().FullName())
}
