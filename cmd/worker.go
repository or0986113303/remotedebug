package cmd

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var (
	workerCmd = &cobra.Command{
		Use:   "worker",
		Short: "Run worker example",
		Run:   runworker,
	}
)

const (
	ModeUnknown = iota
	ModeNormal
	ModeNormalSafeEnabled
	ModeNormalBootCompleted
	ModeSafe
)

func init() {
	rootCmd.AddCommand(workerCmd)
}

func worker(ctx context.Context, name string) {

	for {
		select {
		case <-ctx.Done():
			log.Info("Done for %s \n", name)
			return
		default:
			time.Sleep(1 * time.Second)
			log.Info("Keep working for %s \n", name)
		}
	}
}

func runworker(cmd *cobra.Command, args []string) {
	testctx, testcancel := context.WithCancel(context.Background())
	go worker(testctx, "worker1")
	go worker(testctx, "worker2")
	go worker(testctx, "worker3")

	log.Info(ModeNormal)
	log.Info(ModeNormalSafeEnabled)

	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit

	log.Info("stopping ...")
	defer testcancel()
	log.Info("bye bye")
}
