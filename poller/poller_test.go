package poller

import (
	"fmt"
	"os"
	"testing"

	"go.uber.org/zap"

	"github.com/btcsuite/btcd/rpcclient"
	"github.com/test-go/testify/require"
)

func TestNew(t *testing.T) {
	rpcEndpoint := os.Getenv("TEST_BTC_RPC_ENDPOINT")
	if rpcEndpoint == "" {
		t.Skip("TEST_BTC_RPC_ENDPOINT not set")
	}
	connCfg := &rpcclient.ConnConfig{
		Host:         rpcEndpoint,
		DisableAuth:  true,
		HTTPPostMode: true,
		DisableTLS:   true,
	}

	client, err := rpcclient.New(connCfg, nil)
	require.NoError(t, err)

	r := &Poller{
		rpcClient: client,
		logger:    zap.NewNop(),
	}

	headBlock, err := r.GetHeadBlock()
	require.NoError(t, err)
	fmt.Println("Head Block", headBlock.String())

	libBlock, err := r.GetFinalizedBlock(headBlock)
	require.NoError(t, err)
	fmt.Println("Finalize Block", libBlock.String())

}
