package cmd

import (
	"fmt"
	"math"
	"time"

	flags "github.com/daan-hoogland/crawl/cmd/flags"

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
		log.Debugln(int(math.Ceil(0.2 * float64(flags.MaxProcs))))
		start := time.Now()
		internal.StartJobs(internal.NewExpected(flags.Name, flags.Regex, flags.Size, flags.Hash, flags.Algorithm), flags.MaxProcs, flags.Directory)
		log.Infoln(time.Since(start))
		fmt.Printf("%v\n", internal.Results)
	},
}

// init sets the flags for analyseCmd and adds it as a subcommand
// of the root command.
func init() {
	flags.InitFlags(analyseCmd)
	rootCmd.AddCommand(analyseCmd)
}
