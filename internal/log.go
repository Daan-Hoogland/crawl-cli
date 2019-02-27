package internal

import (
	log "github.com/sirupsen/logrus"
)

// InitLog initializes the log interface with the correct logging level.
func InitLog() {
	if Verbose {
		log.SetLevel(log.InfoLevel)
	} else if Debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)

	}
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
