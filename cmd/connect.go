package cmd

import (
	"net/http"

	pkg "github.com/daan-hoogland/crawl/pkg"

	internal "github.com/daan-hoogland/crawl/internal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Tests the connection to the web application.",
	Long: `The connect commands tests the connection between the application
and the web application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.WithField("command", "scan")
		log.Traceln("Executing connect command")
		url := pkg.GenerateURL(internal.External, internal.Port, nil, false, "status", "connection")
		log.WithField("category", "url").Debugln(url)
		resp, err := http.Get(url)
		if err != nil {
			log.WithFields(log.Fields{"category": "status", "subcategory": "http"}).Warnln("error during connection")
			log.WithFields(log.Fields{"category": "status", "subcategory": "http"}).Errorln(err)
		} else {
			log.WithFields(log.Fields{"category": "status", "subcategory": "http"}).Infoln("successful connection")
			log.WithFields(log.Fields{"category": "status", "subcategory": "http"}).Infoln(resp.Status)
		}
	},
}

// init sets the flags for connectCmd and adds it as a subcommand
// of the root command.
func init() {
	internal.InitFlags(connectCmd)
	rootCmd.AddCommand(connectCmd)
}
