package stranlz

import (
	"unicode"
)

// AlphaOnly analyzer function checks whether string s is composed of letters only.
func AlphaOnly(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
