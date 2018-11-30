package stranlz_test

import (
	"testing"

	stranlz "github.com/t-pwk/go-string-analyzer"
)

func TestAlphaOnly(t *testing.T) {
	fn := stranlz.AlphaOnly
	data := []*data{
		{"asdfghjkl", true},
		{"日本語", true},
		{"Ärger", true},
		{"ärger", true},
		{"asdf0", false},
		{"asdf.", false},
		{"as df", false},
		{"as-df", false},
		{"asdf$", false},
	}

	execute(t, fn, data)
}
