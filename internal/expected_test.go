package internal

import (
	"reflect"
	"testing"
)

func Test_Creation(t *testing.T) {
	exp := []string{"name"}
	res := NewExpected([]string{"name"}, nil, 0, "", "")
	if reflect.DeepEqual(exp, res.name) {
		t.Errorf("get new expected = %s want %s", res.name, exp)
	}
}
