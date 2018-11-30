package stranlz

// LenEq analyzer function checks whether `len(p) == l`
func LenEq(l int) AnalyzeFunc {
	return func(p string) bool { return len(p) == l }
}

// LenGr analyzer function checks whether `len(p) > l`
func LenGr(l int) AnalyzeFunc {
	return func(p string) bool { return len(p) > l }
}

// LenGe analyzer function checks whether `len(p) >= l`
func LenGe(l int) AnalyzeFunc {
	return func(p string) bool { return len(p) >= l }
}

// LenLs analyzer function checks whether `len(p) < l`
func LenLs(l int) AnalyzeFunc {
	return func(p string) bool { return len(p) < l }
}

// LenLe analyzer function checks whether `len(p) <= l`
func LenLe(l int) AnalyzeFunc {
	return func(p string) bool { return len(p) <= l }
}

// LenBe analyzer function checks whether the lenght
// of the input string s meets the following condition `min <= len(s) <= max`
func LenBe(min int, max int) AnalyzeFunc {
	return func(s string) bool {
		if l := len(s); l >= min && l <= max {
			return true
		}
		return false
	}
}
