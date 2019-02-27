package cmd

import (
	internal "github.com/daan-hoogland/crawl/internal"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var scanCmd = &cobra.Command{
	Use:   "scan",
	Short: "The scan command scans the filesystem for files or running services.",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
				Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {
		logrus.Info("Executing scan command")
	},
}

func init() {
	internal.FileFlags(scanCmd)
	rootCmd.AddCommand(scanCmd)
}
