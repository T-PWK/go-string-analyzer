package stranlz

import (
	"unicode"
)

// ContainsNumbers analyzer function checks whether
// string s contains any numeric characters.
func ContainsNumbers(s string) bool {
	for _, r := range s {
		if unicode.IsNumber(r) {
			return true
		}
	}
	return false
}
