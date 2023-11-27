package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"github.com/streamingfast/cli"
	"github.com/streamingfast/cli/sflags"
	"github.com/streamingfast/firehose-bitcoin/poller"
	"go.uber.org/zap"
)

var readerCmd = &cobra.Command{
	Use:          "reader <start-block>",
	Short:        "launches the RPC reader",
	SilenceUsage: true,
	RunE:         readerRunE,
	Args:         cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(readerCmd)
	readerCmd.Flags().String("rpc-endpoint", "http://localhost:8333", "The bitcoin RPC node")
	readerCmd.Flags().Uint64("block-fetch-retry-count", 3, "The number of times to retry fetching a block before ending in error")
	readerCmd.Flags().String("reader-state-storage-path", "/localdata/", "The local path where the reader state will be stored, if blank no state will be stored")
	readerCmd.Flags().Duration("graceful-shutdown-delay", 0*time.Millisecond, "delay before shutting down, after the health endpoint returns unhealthy")
	readerCmd.Flags().Duration("unready-period-delay", 0*time.Millisecond, "the delay starting the shutdown sequence after the health endpoint returns unhealthy")
}

func readerRunE(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	startBlockNumStr := args[0]
	rpcEndpoint := sflags.MustGetString(cmd, "rpc-endpoint")
	blockFetchRetryCount := sflags.MustGetUint64(cmd, "block-fetch-retry-count")
	readerStateStoragePath := sflags.MustGetString(cmd, "reader-state-storage-path")
	gracefulShutdownDelay := sflags.MustGetDuration(cmd, "graceful-shutdown-delay")
	unreadyPeriodDelay := sflags.MustGetDuration(cmd, "unready-period-delay")

	zlog.Info("launching firebtc reader",
		zap.String("start_block_num", startBlockNumStr),
		zap.String("rpc_endpoint", rpcEndpoint),
		zap.Uint64("block_fetch_retry_count", blockFetchRetryCount),
		zap.String("reader_state_storage_path", readerStateStoragePath),
		zap.Duration("graceful_shutdown_delay", gracefulShutdownDelay),
		zap.Duration("unready_period_delay", unreadyPeriodDelay),
	)

	startBlockNum, err := strconv.ParseUint(startBlockNumStr, 10, 64)
	if err != nil {
		return fmt.Errorf("unable to parse start block number %s: %w", startBlockNumStr, err)
	}

	r := poller.New(rpcEndpoint, blockFetchRetryCount, readerStateStoragePath, startBlockNum, zlog)
	app := cli.NewApplication(ctx)
	app.SuperviseAndStart(r)

	return app.WaitForTermination(zlog, unreadyPeriodDelay, gracefulShutdownDelay)
}
