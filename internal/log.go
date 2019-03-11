package internal

import (
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
)

// InitLog initializes the log interface with the correct logging level.
func InitLog() {
	if Verbose {
		log.SetLevel(log.InfoLevel)
	} else if Debug {
		log.SetLevel(log.DebugLevel)
	} else if Develop {
		log.SetLevel(log.TraceLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}

	log.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category", "subcategory"},
	})
}

// LogFields returns a Fields object to log the flag fields.
func LogFields() log.Fields {
	return log.Fields{
		"name":      Name,
		"size":      Size,
		"hash":      Hash,
		"algorithm": Algorithm,
		"external":  External,
		"port":      Port,
		"verbose":   Verbose,
	}
}
