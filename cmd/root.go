package cmd

import (
	"fmt"
	"os"

	internal "github.com/daan-hoogland/crawl/internal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "crawl",
	Version: "0.1 ALPHA",
	Short:   "Crawl is a tool to search through a filesystem or services.",
	Long: `A tool to be combined with the web application with the same name.
The application searches for a file or directory with a name, hash
or size and will report any findings back to the web application.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// initialize logging before every single command
		internal.InitLog()
		log.Debugln(cmd)
	},
}

// Execute initializes the Cobra commands.
func Execute() {
	internal.InitFlags(rootCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
