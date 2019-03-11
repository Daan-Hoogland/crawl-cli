package validation

import (
	"errors"
	"net"

	internal "github.com/daan-hoogland/crawl/internal"
	lgr "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var log = lgr.WithField("component", "validation")

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
	log.Traceln("entering validateScan")
	isNameEmpty := false
	if len(internal.Name) == 0 {
		isNameEmpty = true
	}
	log.WithFields(lgr.Fields{"category": "name", "state": isNameEmpty}).Debugln("name check")

	isSizeZero := false
	if internal.Size == 0 {
		isSizeZero = true
	}
	log.WithFields(lgr.Fields{"category": "size", "state": isSizeZero}).Debugln("size check")

	isHashEmpty := isNotEmpty(internal.Hash)
	log.WithFields(lgr.Fields{"category": "hash", "state": isHashEmpty}).Debugln("hash check")
	isAlgorithmEmpty := isNotEmpty(internal.Algorithm)
	log.WithFields(lgr.Fields{"category": "hash", "subcategory": "algorithm", "state": isAlgorithmEmpty}).Debugln("algorithm check")

	isHashFlagsInvalid := isHashEmpty && isAlgorithmEmpty
	log.WithFields(lgr.Fields{"category": "hash", "subcategory": "combined", "state": isHashFlagsInvalid}).Debugln("combined check")

	if isNameEmpty && isSizeZero && isHashFlagsInvalid {
		log.Traceln("leaving validateScan with error")
		return errors.New("Error: one of the following flags is required: \"name\", \"size\" or \"hash\" and \"algorithm\"")
	}

	log.Traceln("leaving validateScan")
	return nil
}

func validateAnalyse() error {
	log.Traceln("entering validateAnalysis")
	log.Traceln("leaving validateAnalysis")
	return nil
}

func validateConnect() error {
	log.Traceln("entering validateConnect")

	if nil == net.ParseIP(internal.External) {
		log.Traceln("leaving validateConnect with error")
		return errors.New("Error: flag \"external\" does not compile to a valid IP address")
	}
	log.Traceln("leaving validateConnect")
	return nil
}

func isNotEmpty(ss ...string) bool {
	for _, s := range ss {
		if s != "" && len(s) > 0 {
			return false
		}
	}
	return true
}
