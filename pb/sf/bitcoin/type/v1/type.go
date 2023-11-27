package pbbtc

import (
	"fmt"
	"time"

	"google.golang.org/protobuf/types/known/anypb"

	pbbstream "github.com/streamingfast/bstream/pb/sf/bstream/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (b *Block) MustToBstreamBlock() *pbbstream.Block {
	blk, err := b.ToBstreamBlock()
	if err != nil {
		panic(fmt.Errorf("unable to convert block to bstream block: %w", err))
	}
	return blk
}

func (b *Block) ToBstreamBlock() (*pbbstream.Block, error) {
	blkNum := uint64(b.Height)
	anyBlk, err := anypb.New(b)
	if err != nil {
		return nil, fmt.Errorf("unable to anypb marshal block: %w", err)
	}
	// modify type URL to

	parentBlockNum := uint64(0)
	if blkNum > 0 {
		parentBlockNum = blkNum - 1
	}

	libNum := uint64(0)
	if blkNum > 6 {
		libNum = blkNum - 6
	}

	return &pbbstream.Block{
		Number:    blkNum,
		Id:        b.Hash,
		ParentId:  b.PreviousHash,
		Timestamp: timestamppb.New(time.Unix(b.Time, 0)),
		LibNum:    libNum,
		ParentNum: parentBlockNum,
		Payload:   anyBlk,
	}, nil
}
