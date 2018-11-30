package stranlz_test

import (
	"testing"

	"github.com/t-pwk/go-string-analyzer"
)

func TestAndOrNot(t *testing.T) {
	fnTrue := func(s string) bool { return true }
	fnFalse := func(s string) bool { return false }

	if !stranlz.And(fnTrue)("") || stranlz.And(fnFalse)("") || stranlz.And(fnTrue, fnFalse)("") {
		t.Errorf("And function has failed")
	}

	if !stranlz.Or(fnTrue)("") || stranlz.Or(fnFalse)("") || !stranlz.Or(fnTrue, fnFalse)("") {
		t.Errorf("Or function has failed")
	}

	if stranlz.Not(fnTrue)("") || !stranlz.Not(fnFalse)("") {
		t.Errorf("Not function has failed")
	}
}
