package stranlz

import (
	"unicode"
)

// LowerAlphaOnly analyzer function checks whether
// string s is composed of lowercase letters only.
func LowerAlphaOnly(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) || !unicode.IsLower(r) {
			return false
		}
	}
	return true
}
