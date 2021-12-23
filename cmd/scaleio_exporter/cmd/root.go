package cmd

import "github.com/spf13/cobra"

func GetRootCmd(args []string) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:          "scaleio_exporter",
		Short:        "scaleio_exporter is a Prometheus Exporter that collects some metrics from ScaleIO",
		SilenceUsage: true,
	}

	rootCmd.SetArgs(args)
	rootCmd.AddCommand(serverCmd())

	return rootCmd
}
