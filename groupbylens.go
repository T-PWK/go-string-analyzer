package stranlz

import (
	"fmt"
	"sort"
)

// GroupByLens is a multi-value analyzer function returning length of the string s.
// This function can be used to break down analyzed strings by its length.
func GroupByLens(areas []int) MultiAnalyzeFunc {

	if len(areas) == 0 || areas[0] == 0 || !sort.IntsAreSorted(areas) {
		panic("Invalid areas: areas must have at least 1 value, the first value must not be zero (0) and all values must be sorted")
	}

	if len(areas) > 1 {
		for i := range areas {
			if i == 0 {
				continue
			}
			if areas[i]-areas[i-1] < 1 {
				panic("Length secions must be apart from each other by at least 2")
			}
		}
	}

	size := len(areas) + 1
	a := make([]int, size)

	copy(a[1:], areas)

	return func(p string) (string, bool) {
		l := len(p)
		i := sort.SearchInts(a, l)

		switch {
		case i < 2:
			return fmt.Sprintf("[0-%d]", a[1]), true
		case i < size:
			return fmt.Sprintf("[%d-%d]", a[i-1]+1, a[i]), true
		case i >= size:
			return fmt.Sprintf("[%d+)", a[size-1]+1), true
		default:
			return "", false
		}
	}
}
