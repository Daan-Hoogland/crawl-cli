package cmd

import (
	internal "github.com/daan-hoogland/crawl/internal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "The scan command scans the filesystem for files or running services.",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Executing connect command")
	},
}

// init sets the flags for the connectCmd and adds it as a subcommand
// of the root command.
func init() {
	initFlags()
	rootCmd.AddCommand(connectCmd)
}

// initFlags sets the flags on the connectCmd.
func initFlags() {
	internal.ExternalFlag(connectCmd)
	internal.PortFlag(connectCmd)
}
