package internal

import (
	"github.com/spf13/cobra"
)

var (
	// External ip/url used to connect to web service.
	External string

	// Port number the web service is reachable at.
	Port int

	// Names of the files/directories to search for.
	Names []string

	// Hash of the file to search for.
	Hash string

	// Size of the file to search for.
	Size int

	// Algorithm used to get the hash value
	Algorithm string
)

// ExternalFlag sets the external flag on the given command.
func externalFlag(c *cobra.Command) {
	c.Flags().StringVarP(&External, "external", "e", "127.0.0.1", "external address pointing to web service")
}

// PortFlag sets the port flag on the given command.
func portFlag(c *cobra.Command) {
	c.Flags().IntVarP(&Port, "port", "p", 9000, "port used to connect to web service")
}

// FileFlags sets all the flags to a command that are in the file category
// (name, size, hash, algorithm)
func fileFlags(c *cobra.Command) {
	c.Flags().StringVar(&Hash, "hash", "", "hash of the file to search for")
	c.Flags().StringVarP(&Algorithm, "algorithm", "a", "md5", "the hash algorithm used")
	c.Flags().IntVarP(&Size, "size", "s", 0, "file size that the target file must match")
	c.Flags().StringSliceVarP(&Names, "name", "n", nil, "name(s) of files to search for")
}

// InitFlags sets the flags on the connectCmd.
func InitFlags(c *cobra.Command) {
	switch c.Name() {
	case "scan":
		fileFlags(c)
		externalFlag(c)
		portFlag(c)
	case "analyse":
		fileFlags(c)
	case "connect":
		externalFlag(c)
		portFlag(c)
	default:
	}
}
