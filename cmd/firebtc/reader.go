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

var pollerCmd = &cobra.Command{
	Use:          "poller <start-block>",
	Short:        "launches the RPC reader",
	SilenceUsage: true,
	RunE:         readerRunE,
	Args:         cobra.ExactArgs(1),
}

func init() {
	rootCmd.AddCommand(pollerCmd)
	pollerCmd.Flags().String("rpc-endpoint", "http://localhost:8333", "The bitcoin RPC node")
	pollerCmd.Flags().Bool("ignore-cursor", false, "When enable it will ignore the cursor and start from the start block num, the cursor will still be saved as the poller progresses")
	pollerCmd.Flags().Uint64("block-fetch-retry-count", 3, "The number of times to retry fetching a block before ending in error")
	pollerCmd.Flags().String("reader-state-storage-path", "/localdata/", "The local path where the reader state will be stored, if blank no state will be stored")
	pollerCmd.Flags().Duration("graceful-shutdown-delay", 0*time.Millisecond, "delay before shutting down, after the health endpoint returns unhealthy")
	pollerCmd.Flags().Duration("unready-period-delay", 0*time.Millisecond, "the delay starting the shutdown sequence after the health endpoint returns unhealthy")
}

func readerRunE(cmd *cobra.Command, args []string) error {
	ctx := cmd.Context()

	startBlockNumStr := args[0]
	rpcEndpoint := sflags.MustGetString(cmd, "rpc-endpoint")
	blockFetchRetryCount := sflags.MustGetUint64(cmd, "block-fetch-retry-count")
	readerStateStoragePath := sflags.MustGetString(cmd, "reader-state-storage-path")
	gracefulShutdownDelay := sflags.MustGetDuration(cmd, "graceful-shutdown-delay")
	unreadyPeriodDelay := sflags.MustGetDuration(cmd, "unready-period-delay")
	ignoreCursor := sflags.MustGetBool(cmd, "ignore-cursor")

	zlog.Info("launching firebtc reader",
		zap.String("start_block_num", startBlockNumStr),
		zap.String("rpc_endpoint", rpcEndpoint),
		zap.Uint64("block_fetch_retry_count", blockFetchRetryCount),
		zap.String("reader_state_storage_path", readerStateStoragePath),
		zap.Duration("graceful_shutdown_delay", gracefulShutdownDelay),
		zap.Duration("unready_period_delay", unreadyPeriodDelay),
		zap.Bool("ignore_cursor", ignoreCursor),
	)

	startBlockNum, err := strconv.ParseUint(startBlockNumStr, 10, 64)
	if err != nil {
		return fmt.Errorf("unable to parse start block number %s: %w", startBlockNumStr, err)
	}

	p := poller.New(rpcEndpoint, blockFetchRetryCount, readerStateStoragePath, startBlockNum, ignoreCursor, zlog)
	app := cli.NewApplication(ctx)
	app.SuperviseAndStart(p)

	return app.WaitForTermination(zlog, unreadyPeriodDelay, gracefulShutdownDelay)
}
