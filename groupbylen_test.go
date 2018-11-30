package stranlz_test

import (
	"testing"

	stranlz "github.com/t-pwk/go-string-analyzer"
)

func TestGroupByLen(t *testing.T) {
	fn := stranlz.GroupByLen
	data := []*multidata{
		{"", "0", true},
		{"1234", "4", true},
	}

	executeMulti(t, fn, data)
}
