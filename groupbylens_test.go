package stranlz_test

import (
	"testing"

	stranlz "github.com/t-pwk/go-string-analyzer"
)

func TestGroupByLens3(t *testing.T) {
	fn := stranlz.GroupByLens([]int{5, 10, 15})
	data := []*multidata{
		{"", "[0-5]", true},
		{"123", "[0-5]", true},
		{"12345", "[0-5]", true},
		{"1234567", "[6-10]", true},
		{"1234567891", "[6-10]", true},
		{"12345678901", "[11-15]", true},
		{"123456789012345", "[11-15]", true},
		{"1234567890123456", "[16+)", true},
	}

	executeMulti(t, fn, data)
}

func TestGroupByLens1(t *testing.T) {
	fn := stranlz.GroupByLens([]int{5})
	data := []*multidata{
		{"", "[0-5]", true},
		{"123", "[0-5]", true},
		{"12345", "[0-5]", true},
		{"1234567891", "[6+)", true},
	}

	executeMulti(t, fn, data)
}
