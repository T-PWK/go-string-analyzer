package stranlz_test

import (
	"testing"

	"github.com/t-pwk/go-string-analyzer"
)

func TestRegexp(t *testing.T) {
	fn := stranlz.Regexp("\\d{3}$")

	data := []*data{
		{"asdfghjkl", false},
		{"日本語", false},
		{"123Ärger", false},
		{"är123ger", false},
		{"123", true},
		{"123123123", true},
		{" 000", true},
		{"123 ", false},
	}

	execute(t, fn, data)
}
