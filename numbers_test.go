package stranlz_test

import (
	"testing"

	"github.com/t-pwk/go-string-analyzer"
)

func TestContainsNumbers(t *testing.T) {
	fn := stranlz.ContainsNumbers
	data := []*data{
		{"asdfghjkl", false},
		{"日本語", false},
		{"Ärger", false},
		{"ärger", false},
		{"asdf0", true},
		{"123123123", true},
		{"as df", false},
	}

	execute(t, fn, data)
}
