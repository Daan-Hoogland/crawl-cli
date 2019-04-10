package cmd

import (
	"fmt"
	"math"
	"time"

	internal "github.com/daan-hoogland/crawl/internal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var analyseCmd = &cobra.Command{
	Use:   "analyse",
	Short: "The analyse command scans the filesystem for files or running services.",
	Long: `The analyse command scans the filesystem for files or running services.
Unlike the scan command, the analyse command does not send the results
to a running web application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.WithField("command", "analyse")
		log.Traceln("executing analyse command")
		log.Debugln(int(math.Ceil(0.2 * float64(internal.MaxProcs))))
		start := time.Now()
		internal.StartJobs()
		log.Infoln(time.Since(start))
		fmt.Printf("%v\n", internal.Results)
	},
}

// init sets the flags for analyseCmd and adds it as a subcommand
// of the root command.
func init() {
	internal.InitFlags(analyseCmd)
	rootCmd.AddCommand(analyseCmd)
}
