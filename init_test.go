package stranlz_test

import (
	"testing"

	stranlz "github.com/t-pwk/go-string-analyzer"
)

type data struct {
	s  string
	ok bool
}

type multidata struct {
	s, v string
	ok   bool
}

func execute(t *testing.T, fn stranlz.AnalyzeFunc, data []*data) {
	for _, d := range data {
		if r := fn(d.s); r != d.ok {
			t.Errorf("Test %s failed for input: `%s`, got: %t, want: %t", t.Name(), d.s, r, d.ok)
		}
	}
}

func executeMulti(t *testing.T, fn stranlz.MultiAnalyzeFunc, data []*multidata) {
	for _, d := range data {
		if v, ok := fn(d.s); v != d.v || ok != d.ok {
			t.Errorf("Test %s failed for input: `%s`, got: (%s, %t), want: (%s, %t)", t.Name(), d.s, v, ok, d.v, d.ok)
		}
	}
}

func fnTrue(s string) bool  { return true }
func fnFalse(s string) bool { return false }
