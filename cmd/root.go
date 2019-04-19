package cmd

import (
	"fmt"
	"os"

	internal "github.com/daan-hoogland/crawl-cli/internal"

	flags "github.com/daan-hoogland/crawl-cli/cmd/flags"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "crawl",
	Version: "0.3",
	Short:   "Crawl is a tool to search through a filesystem or services.",
	Long: `A tool to be combined with the web application with the same name.
The application searches for a file or directory with a name, hash
or size and will report any findings back to the user.`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// initialize logging before every single command
		internal.InitLog(cmd, flags.Debug, flags.Develop, flags.Verbose)
		error := flags.CheckRequiredFlags(cmd)
		if (error) != nil {
			fmt.Println(error)
			cmd.Usage()
			os.Exit(1)
		}
	},
}

func init() {
	flags.InitFlags(rootCmd)
}

// Execute initializes the Cobra commands.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
