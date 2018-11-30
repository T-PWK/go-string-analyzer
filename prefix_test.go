package stranlz_test

import (
	"testing"

	"github.com/t-pwk/go-string-analyzer"
)

func TestPrefix(t *testing.T) {
	fn := stranlz.Prefix("123")

	data := []*data{
		{"asdfghjkl", false},
		{"日本語", false},
		{"Ärger", false},
		{"ärger", false},
		{"123", true},
		{"123123123", true},
		{" 123", false},
		{"12", false},
	}

	execute(t, fn, data)
}
