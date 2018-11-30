package stranlz

import (
	"regexp"
)

// Regexp analyzer function checks whether the string s starts with the pattern p.
func Regexp(p string) AnalyzeFunc {
	r := regexp.MustCompile(p)

	return func(s string) bool {
		return r.MatchString(s)
	}
}
