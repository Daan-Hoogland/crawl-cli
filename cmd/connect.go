package cmd

import (
	"net/http"

	pkg "github.com/daan-hoogland/crawl-cli/pkg"

	flags "github.com/daan-hoogland/crawl-cli/cmd/flags"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	connectCmd = &cobra.Command{
		Use:   "connect",
		Short: "Tests the connection to the web application.",
		Long: `The connect commands tests the connection between the application
and the web application.`,
		Run: func(cmd *cobra.Command, args []string) {
			log.WithField("command", "connect")
			log.Traceln("executing connect command")
			url := pkg.GenerateURL(flags.External, flags.Port, nil, false, "status", "connection")
			log.WithField("category", "url").Debugln(url)
			resp, err := http.Get(url)
			if err != nil || resp.StatusCode != 200 {
				println("Error during connection")
				log.WithFields(log.Fields{"category": "status", "subcategory": "http"}).Debugln("error during connection")
				log.WithFields(log.Fields{"category": "status", "subcategory": "http"}).Errorln(err)
				log.WithFields(log.Fields{"category": "status", "subcategory": "http"}).Infoln(resp.Status)
			} else {
				println("Successful connection")
				log.WithFields(log.Fields{"category": "status", "subcategory": "http"}).Infoln(resp.Status)
			}
		},
	}
)

// init sets the flags for connectCmd and adds it as a subcommand
// of the root command.
func init() {
	flags.InitFlags(connectCmd)
	rootCmd.AddCommand(connectCmd)
}
