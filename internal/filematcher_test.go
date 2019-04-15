package internal

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"hash"
	"reflect"
	"testing"

	"golang.org/x/crypto/sha3"
)

func Test_getHasher(t *testing.T) {
	functions := map[string]hash.Hash{
		"md5":        md5.New(),
		"sha224":     sha256.New224(),
		"sha256":     sha256.New(),
		"sha384":     sha512.New384(),
		"sha512":     sha512.New(),
		"sha512/224": sha512.New512_224(),
		"sha512/256": sha512.New512_256(),
		"sha3-224":   sha3.New224(),
		"sha3-256":   sha3.New256(),
		"sha3-384":   sha3.New384(),
		"sha3-512":   sha3.New512(),
		"default":    sha256.New(),
	}

	for key, val := range functions {
		res := getHasher(key)
		if reflect.TypeOf(res) != reflect.TypeOf(val) {
			t.Errorf("get md5 hash = %s want %s", reflect.TypeOf(res), reflect.TypeOf(val))

		}
	}
}

func Test_ValidateResultName(t *testing.T) {
	res := Result{
		name:   true,
		regex:  false,
		size:   false,
		digest: false,
	}

	exp = newExpected([]string{"test"}, []string{}, 0, "", "")

	result := ValidateResult(exp, &res)
	if result != true {
		t.Errorf("valudate name = %t want %t", result, true)
	}
}

func Test_ValidateResultRegex(t *testing.T) {
	res := Result{
		name:   false,
		regex:  true,
		size:   false,
		digest: false,
	}

	exp = newExpected([]string{}, []string{"test"}, 0, "", "")

	result := ValidateResult(exp, &res)
	if result != true {
		t.Errorf("validate regex = %t want %t", result, true)
	}
}

func Test_ValidateResultSize(t *testing.T) {
	res := Result{
		name:   false,
		regex:  false,
		size:   true,
		digest: false,
	}

	exp = newExpected([]string{}, []string{}, 500, "", "")

	result := ValidateResult(exp, &res)
	if result != true {
		t.Errorf("validate size = %t want %t", result, true)
	}
}

func Test_ValidateResultDigest(t *testing.T) {
	res := Result{
		name:   false,
		regex:  false,
		size:   false,
		digest: true,
	}

	exp = newExpected([]string{}, []string{""}, 0, "16e8296ff0f34df33f8ce96610606173", "md5")

	result := ValidateResult(exp, &res)
	if result != true {
		t.Errorf("validate digest = %t want %t", result, true)
	}
}

func Test_ValidateResultFail(t *testing.T) {
	res := Result{
		name:   false,
		regex:  false,
		size:   false,
		digest: false,
	}

	exp = newExpected([]string{"fdff"}, []string{}, 0, "", "")

	result := ValidateResult(exp, &res)
	if false != result {
		t.Errorf("validate name and size = %t want %t", result, true)
	}
}

func Test_ValidateResult(t *testing.T) {
	res := Result{
		name:   true,
		regex:  true,
		size:   true,
		digest: true,
	}

	exp = newExpected([]string{"test1"}, []string{"test2"}, 500, "16e8296ff0f34df33f8ce96610606173", "md5")

	result := ValidateResult(exp, &res)
	if true != result {
		t.Errorf("validate all = %t want %t", result, true)
	}
}
