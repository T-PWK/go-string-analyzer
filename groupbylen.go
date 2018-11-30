package stranlz

import (
	"strconv"
)

// GroupByLen is a multi-value analyzer function returning length of the string s.
// This function can be used to break down analyzed strings by its length.
func GroupByLen(s string) (string, bool) {
	return strconv.Itoa(len(s)), true
}
