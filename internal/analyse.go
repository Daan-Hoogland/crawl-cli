package internal

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"io"
	"log"
	"os"
	"regexp"

	"golang.org/x/crypto/sha3"
)

// MatchFile attempts to match a file according to the values given in program args.
func MatchFile(file os.FileInfo, path string) bool {
	regexMatch := false
	for _, regex := range Regex {
		r, _ := regexp.Compile(regex)

		// Despite reassignment, should still be false as long as it doesn't match.
		if !regexMatch {
			regexMatch = r.MatchString(file.Name())
		}
	}

	nameMatch := false
	for _, n := range Name {
		if !nameMatch {
			nameMatch = file.Name() == n
		}
	}

	sizeMatch := file.Size() == Size
	// Open file for reading
	fileStream, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer fileStream.Close()

	// Create new hasher, which is a writer interface. Default sha256.
	var hasher = getHasher(Algorithm)
	_, err = io.Copy(hasher, fileStream)
	if err != nil {
		log.Fatal(err)
	}

	// Hash and print. Pass nil since
	// the data is not coming in as a slice argument
	// but is coming through the writer interface
	hashMatch := bytes.Equal(hasher.Sum(nil), []byte(Hash))

	return false
}

//getHasher returns a new hash.Hash object depending on input string given.
func getHasher(algorithm string) hash.Hash {
	switch algorithm {
	case "md5":
		return md5.New()
	case "sha224":
		return sha256.New224()
	case "sha256":
		return sha256.New()
	case "sha384":
		return sha512.New384()
	case "sha512":
		return sha512.New()
	case "sha512/224":
		return sha512.New512_224()
	case "sha512/256":
		return sha512.New512_256()
	case "sha3-224":
		return sha3.New224()
	case "sha3-256":
		return sha3.New256()
	case "sha3-384":
		return sha3.New384()
	case "sha3-512":
		return sha3.New512()
	default:
		return sha256.New()
	}
}
