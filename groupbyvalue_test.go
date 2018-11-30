package stranlz_test

import (
	"testing"

	stranlz "github.com/t-pwk/go-string-analyzer"
)

func TestGroupByValue(t *testing.T) {
	fn := stranlz.GroupByValue
	data := []*multidata{
		{"", "", true},
		{"OK", "OK", true},
		{"123123", "123123", true},
	}

	executeMulti(t, fn, data)
}
