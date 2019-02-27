package cmd

import (
	internal "github.com/daan-hoogland/crawl/internal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "The scan command scans the filesystem for files or running services.",
	Long: `The scan commands scans the filesystem for files or running applications
and compares them to the given input. Any matches will be sent
to the connected web application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Executing scan command")
	},
}

// init sets the flags for scanCmd and adds it as a subcommand
// of the root command.
func init() {
	internal.InitFlags(scanCmd)
	rootCmd.AddCommand(scanCmd)
}
