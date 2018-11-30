package stranlz_test

import (
	"testing"

	stranlz "github.com/t-pwk/go-string-analyzer"
)

func BenchmarkLowerAlphaOnly(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stranlz.LowerAlphaOnly("a")
	}
}

func TestLowerAlphaOnly(t *testing.T) {
	fn := stranlz.LowerAlphaOnly
	data := []*data{
		{"asdfghjkl", true},
		{"日本語", false},
		{"Ärger", false},
		{"ärger", true},
		{"asdf0", false},
		{"asdf.", false},
		{"as df", false},
	}

	execute(t, fn, data)
}
