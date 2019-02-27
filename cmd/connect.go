package cmd

import (
	internal "github.com/daan-hoogland/crawl/internal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Tests the connection to the web application.",
	Long: `The connect commands tests the connection between the application
and the web application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Executing connect command")
	},
}

// init sets the flags for connectCmd and adds it as a subcommand
// of the root command.
func init() {
	internal.InitFlags(connectCmd)
	rootCmd.AddCommand(connectCmd)
}
