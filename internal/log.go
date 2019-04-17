package internal

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// InitLog initializes the log interface with the correct logging level.
func InitLog(cmd *cobra.Command, debug bool, develop bool, verbose bool) {
	if debug {
		log.SetLevel(log.DebugLevel)
	} else if develop {
		log.SetLevel(log.TraceLevel)
	} else if verbose {
		log.SetLevel(log.InfoLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"command", "component", "category", "subcategory"},
	})
}
