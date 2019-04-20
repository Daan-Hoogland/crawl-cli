package internal

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"io"
	"os"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/sha3"
)

//Result object contains results of match checks.
type Result struct {
	name   bool
	regex  bool
	size   bool
	digest bool
}

// MatchFile attempts to match a file according to the values given in program args.
func MatchFile(file os.FileInfo, path string, exp *Expected) *Result {
	var result Result
	// log.WithField("component", path).Debugln("")
	for _, regex := range exp.regex {
		r, _ := regexp.Compile(regex)

		// Despite reassignment, should still be false as long as it doesn't match.
		if !result.digest {
			result.regex = r.MatchString(file.Name())
		}
	}

	for _, n := range exp.name {
		if !result.name {
			result.name = file.Name() == n
		}
	}

	result.size = file.Size() == exp.size

	// Open file for reading
	fileStream, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer fileStream.Close()

	// Create new hasher, which is a writer interface. Default sha256.
	var hasher = getHasher(exp.hash.function)
	log.WithField("component", "func").Debugln(exp.hash.function)
	_, err = io.Copy(hasher, fileStream)
	if err != nil {
		log.Fatal(err)
	}

	log.WithField("component", "res").Debugln(hex.EncodeToString(hasher.Sum(nil)))
	log.WithField("component", "exp").Debugln(exp.hash.digest)

	// Hash and print. Pass nil since
	// the data is not coming in as a slice argument
	// but is coming through the writer interface
	// result.digest = bytes.Equal(hasher.Sum(nil), []byte(exp.hash.digest))
	result.digest = hex.EncodeToString(hasher.Sum(nil)) == exp.hash.digest

	return &result
}

//getHasher returns a new hash.Hash object depending on input string given.
func getHasher(function string) hash.Hash {
	switch strings.ToLower(function) {
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

//ValidateResult compares expected values to the result values.
//todo: needs to check if more than one has to be valid.
func (res *Result) ValidateResult(exp *Expected) bool {
	match := false
	if len(exp.name) > 0 {
		match = res.name
	}

	if len(exp.regex) > 0 {
		if !match {
			match = res.regex
		} else {
			match = match && res.regex
		}
	}

	if exp.size > 0 {
		if !match {
			match = res.size
		} else {
			match = match && res.size
		}
	}

	if len(exp.hash.digest) > 0 {
		if !match {
			match = res.digest
		} else {
			match = match && res.digest
		}
	}

	return match
}
