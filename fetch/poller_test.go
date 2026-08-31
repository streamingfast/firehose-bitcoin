package fetch

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
	"time"

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

	r := &BlockFetcher{
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

func TestWaitForBlockHeight(t *testing.T) {
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

	r := &BlockFetcher{
		rpcClient:            client,
		headBlockWaitTimeout: 5 * time.Second,
		logger:               zap.NewNop(),
	}

	headBlock, err := r.GetHeadBlock()
	require.NoError(t, err)

	// Already reached: returns the head without waiting.
	reached, err := r.waitForBlockHeight(headBlock.Num())
	require.NoError(t, err)
	require.Equal(t, headBlock.Num(), reached.Num())

	// Not reached: blocks for the timeout, then reports the head the node is still on.
	start := time.Now()
	pending, err := r.waitForBlockHeight(headBlock.Num() + 1)
	require.NoError(t, err)
	require.True(t, time.Since(start) >= r.headBlockWaitTimeout)
	require.True(t, pending.Num() <= headBlock.Num())
	fmt.Println("waited on head", pending.String())
}

// Any waitforblockheight failure must fall back to reading the head block, not fail the fetch:
// an unsupported method, and equally a proxy that cuts the long-lived request.
func TestAwaitBlockHeightFallsBack(t *testing.T) {
	tests := []struct {
		name  string
		reply func(w http.ResponseWriter, id json.RawMessage)
	}{
		{"method not found", func(w http.ResponseWriter, id json.RawMessage) {
			_, _ = fmt.Fprintf(w, `{"result":null,"error":{"code":-32601,"message":"Method not found"},"id":%s}`, id)
		}},
		{"proxy cut the request", func(w http.ResponseWriter, _ json.RawMessage) {
			w.WriteHeader(http.StatusGatewayTimeout)
		}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var methods []string
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				var req struct {
					ID     json.RawMessage `json:"id"`
					Method string          `json:"method"`
				}
				require.NoError(t, json.NewDecoder(r.Body).Decode(&req))
				methods = append(methods, req.Method)
				w.Header().Set("Content-Type", "application/json")
				if req.Method == "waitforblockheight" {
					tc.reply(w, req.ID)
					return
				}
				// The head-block read then fails distinguishably.
				_, _ = fmt.Fprintf(w, `{"result":null,"error":{"code":-1,"message":"nope"},"id":%s}`, req.ID)
			}))
			defer srv.Close()

			client, err := rpcclient.New(&rpcclient.ConnConfig{
				Host:         strings.TrimPrefix(srv.URL, "http://"),
				DisableAuth:  true,
				HTTPPostMode: true,
				DisableTLS:   true,
			}, nil)
			require.NoError(t, err)

			r := &BlockFetcher{
				rpcClient:            client,
				headBlockWaitTimeout: time.Millisecond,
				logger:               zap.NewNop(),
			}

			_, err = r.awaitBlockHeight(100)
			require.Error(t, err)
			require.Contains(t, err.Error(), "failed to get head block")
			require.Equal(t, []string{"waitforblockheight", "getbestblockhash"}, methods)
		})
	}
}
