package cmd

import (
	"math"
	"os"
	"time"

	"github.com/briandowns/spinner"
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
		lgr := log.WithField("command", "analyse")
		lgr.Debugln(int(math.Ceil(0.2 * float64(flags.MaxProcs))))
		start := time.Now()
		s := spinner.New([]string{"\\", "|", "/", "-"}, 150*time.Millisecond)
		s.Suffix = "  Crawling files..."
		s.Writer = os.Stderr
		s.Start()
		internal.StartJobs(internal.NewExpected(flags.Name, flags.Regex, flags.Size, flags.Hash, flags.Algorithm), flags.MaxProcs, flags.Directory)
		s.Stop()
		executionTime := time.Since(start)
		data := internal.ResultTo2DSlice(&internal.Results)
		table := internal.GenerateTable(data)
		table.Render()
		lgr.WithField("component", "execution time").Infoln(executionTime.String())
	},
}

// init sets the flags for analyseCmd and adds it as a subcommand
// of the root command.
func init() {
	flags.InitFlags(analyseCmd)
	rootCmd.AddCommand(analyseCmd)
}
