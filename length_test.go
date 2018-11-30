package stranlz_test

import (
	"testing"

	stranlz "github.com/t-pwk/go-string-analyzer"
)

func TestLenEq(t *testing.T) {
	fn := stranlz.LenEq(5)
	data := []*data{
		{"", false},
		{"12345", true},
		{"1234512345", false},
	}

	execute(t, fn, data)
}

func TestLenGe(t *testing.T) {
	fn := stranlz.LenGe(5)
	data := []*data{
		{"", false},
		{"12345", true},
		{"1234512345", true},
	}

	execute(t, fn, data)
}

func TestLenGr(t *testing.T) {
	fn := stranlz.LenGr(5)
	data := []*data{
		{"", false},
		{"12345", false},
		{"1234512345", true},
	}

	execute(t, fn, data)
}

func TestLenLs(t *testing.T) {
	fn := stranlz.LenLs(5)
	data := []*data{
		{"", true},
		{"1", true},
		{"12345", false},
		{"1234512345", false},
	}

	execute(t, fn, data)
}

func TestLenLe(t *testing.T) {
	fn := stranlz.LenLe(5)
	data := []*data{
		{"", true},
		{"1", true},
		{"12345", true},
		{"1234512345", false},
	}

	execute(t, fn, data)
}

func TestLenBe(t *testing.T) {
	fn := stranlz.LenBe(2, 5)
	data := []*data{
		{"", false},
		{"1", false},
		{"12", true},
		{"1234", true},
		{"12345", true},
		{"123456", false},
	}

	execute(t, fn, data)
}
