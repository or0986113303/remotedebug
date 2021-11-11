package cmd

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/or0986113303/remotedebug/service/apiman"
	"github.com/spf13/cobra"
)

var (
	daemonCmd = &cobra.Command{
		Use:   "daemon",
		Short: "Run remotedebug as daemon",
		Run:   daemon,
	}
)

func init() {
	rootCmd.AddCommand(daemonCmd)
}

func daemon(cmd *cobra.Command, args []string) {
	server := apiman.New()
	server.StartServer()

	quit := make(chan os.Signal, 2)
	signal.Notify(quit, os.Interrupt)
	signal.Notify(quit, syscall.SIGTERM)
	<-quit

	log.Info("stopping ...")
	defer server.StopServer()
	log.Info("bye bye")
}
