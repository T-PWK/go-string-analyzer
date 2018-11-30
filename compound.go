package stranlz

// And returns an compound analyzer function that returns true
// only if every analyzer function returns true
func And(fns ...AnalyzeFunc) AnalyzeFunc {

	return func(s string) bool {
		for _, fn := range fns {
			if r := fn(s); !r {
				return false
			}
		}
		return true
	}
}

// Or returns an compound analyzer function that returns true
// only if at least one analyzer function returns true
func Or(fns ...AnalyzeFunc) AnalyzeFunc {

	return func(s string) bool {
		for _, fn := range fns {
			if r := fn(s); r {
				return true
			}
		}
		return false
	}
}

// Not returns an compound analyzer function that returns true
// only if the input function fn returned false
func Not(fn AnalyzeFunc) AnalyzeFunc {
	return func(s string) bool {
		return !fn(s)
	}
}
