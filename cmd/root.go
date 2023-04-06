package cmd

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const CLIVersion float64 = 0.1

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kli",
	Short: "Kubernetes CLI for SRE",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
