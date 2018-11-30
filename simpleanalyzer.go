package stranlz

// SimpleAnalyzer is a wrapper of analyzing function fn.
// Analyzer contains a counter of the successful analysis
type SimpleAnalyzer struct {
	id    string
	total int32
	count int32
	fn    AnalyzeFunc
}

// NewSimpleAnalyzer creates a new analyzer with the given
// identifier and the given analyzer function
func NewSimpleAnalyzer(id string, fn AnalyzeFunc) *SimpleAnalyzer {
	return &SimpleAnalyzer{id: id, fn: fn}
}

// Analyze performs analysis of the given string using
// analyzer function fn. If the input string matches
// analyzer rules i.e. the function fn returns true
// the counter is incremented by one
func (a *SimpleAnalyzer) Analyze(s string) {
	a.total++
	if a.fn(s) {
		a.count++
	}
}

// ID returns analyzer identifier
func (a *SimpleAnalyzer) ID() string {
	return a.id
}

// Reset sets the counter to zero
func (a *SimpleAnalyzer) Reset() {
	a.count = 0
}

// Counter returns the counter value
// i.e. number of successful matches
func (a *SimpleAnalyzer) Counter() int32 {
	return a.count
}

// Total returns the number of analyzed values
func (a *SimpleAnalyzer) Total() int32 {
	return a.total
}
