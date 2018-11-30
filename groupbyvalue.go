package stranlz

// GroupByValue is a multi-value analyzer function returning string s itslf.
// This function can be used to break down analyzed strings by its value.
func GroupByValue(s string) (string, bool) {
	return s, true
}
