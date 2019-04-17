package cmd

import (
	flags "github.com/daan-hoogland/crawl/cmd/flags"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "The scan command scans the filesystem for files or running services.",
	Long: `The scan commands scans the filesystem for files or running applications
and compares them to the given input. Any matches will be sent
to the connected web application.`,
	SuggestFor: []string{"start"},
	Run: func(cmd *cobra.Command, args []string) {
		log.WithField("command", "scan")
	},
}

// init sets the flags for scanCmd and adds it as a subcommand
// of the root command.
func init() {
	flags.InitFlags(scanCmd)
	rootCmd.AddCommand(scanCmd)
}
