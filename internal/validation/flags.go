package validation

import (
	"errors"

	internal "github.com/daan-hoogland/crawl/internal"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// CheckRequiredFlags validates the flags given to command cmd.
func CheckRequiredFlags(cmd *cobra.Command) error {
	switch cmd.Name() {
	case "scan":
		return validateScan()
	case "analyse":
		return validateAnalyse()
	case "connect":
		return validateConnect()
	default:
		return nil
	}
}

func validateScan() error {
	isNameEmpty := false
	if len(internal.Name) == 0 {
		isNameEmpty = true
	}

	isSizeZero := false
	if internal.Size == 0 {
		isSizeZero = true
	}

	isHashEmpty := isEmpty(internal.Hash)
	isAlgorithmEmpty := isEmpty(internal.Algorithm)

	isHashFlagsInvalid := (!isHashEmpty && isAlgorithmEmpty) || (isHashEmpty && !isAlgorithmEmpty)

	if isNameEmpty && isSizeZero && isHashFlagsInvalid {
		return errors.New("Error: one of the following flags is required: \"name\", \"size\" or \"hash\" and \"algorithm\"")
	}

	return nil
}

func validateAnalyse() error {
	log.Errorln("Invalid flags for analyse")
	log.WithFields(internal.LogFields()).Debugln("Debug flags")
	return nil
}

func validateConnect() error {
	log.Errorln("Invalid flags for connect")
	log.WithFields(internal.LogFields()).Debugln("Debug flags")
	return nil
}

func isEmpty(ss ...string) bool {
	for _, s := range ss {
		if s != "" {
			return true
		}
	}
	return false
}
