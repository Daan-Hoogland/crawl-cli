package internal

import (
	"bytes"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"io"
	"os"
	"regexp"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"
)

//Result object contains results of match checks.
type Result struct {
	Name  bool
	Regex bool
	Size  bool
	Hash  bool
}

// MatchFile attempts to match a file according to the values given in program args.
func MatchFile(file os.FileInfo, path string) Result {
	// log.WithField("component", path).Debugln("")
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

	return Result{
		Name:  nameMatch,
		Regex: regexMatch,
		Size:  sizeMatch,
		Hash:  hashMatch,
	}
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

func ValidateResult(res Result) bool {
	match := false
	if len(Name) > 0 {
		match = res.Name
	}

	if len(Regex) > 0 {
		if !match {
			match = res.Regex
		} else {
			match = match && res.Regex
		}
	}

	if Size > 0 {
		if !match {
			match = res.Size
		} else {
			match = match && res.Size
		}
	}

	if len(Hash) > 0 {
		if !match {
			match = res.Hash
		} else {
			match = match && res.Hash
		}
	}

	return match
}
