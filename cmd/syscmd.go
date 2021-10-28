package cmd

import (
	"github.com/or0986113303/remotedebug/pkg/system"

	"github.com/spf13/cobra"
)

var (
	syscmd = &cobra.Command{
		Use:   "sys",
		Short: "Route to system command",
		Long:  `Route to system command`,
		Run:   sysinfo,
	}
	cpuuseagecmd = &cobra.Command{
		Use:   "cpu",
		Short: "Route to system command",
		Long:  `Route to system command`,
		Run:   cpuuseage,
	}
)

func init() {
	rootCmd.AddCommand(syscmd)
	rootCmd.AddCommand(cpuuseagecmd)
}

func sysinfo(cmd *cobra.Command, args []string) {
	worker := system.New()
	env, err := worker.Sysinfo()
	if err != nil {
		log.Info(err)
	} else {
		log.Info(env)
	}
}

func cpuuseage(cmd *cobra.Command, args []string) {
	worker := system.New()
	cpuuseage := worker.CPUUseage()
	log.Info(cpuuseage)
}
