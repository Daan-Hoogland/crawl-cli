package internal

//Expected struct containing values to be expected.
type Expected struct {
	name  []string
	regex []string
	size  int64
	hash  *combinedHash
}

type combinedHash struct {
	digest   string
	function string
}

//NewExpected creates a new struct of type Expected
func NewExpected(name []string, regex []string, size int64, digest string, function string) *Expected {
	return &Expected{
		name:  name,
		regex: regex,
		size:  size,
		hash:  newHash(digest, function),
	}
}

//newHash generates a new struct of type combinedHash
func newHash(digest string, function string) *combinedHash {
	return &combinedHash{
		digest:   digest,
		function: function,
	}
}
