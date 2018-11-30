package stranlz

import (
	"strings"
)

// Prefix analyzer function checks whether the string s starts with the string p.
func Prefix(p string) AnalyzeFunc {

	return func(s string) bool {
		return strings.HasPrefix(s, p)
	}
}
