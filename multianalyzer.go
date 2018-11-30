package stranlz

// MultiAnalyzer is a wrapper of analyzing function fn.
// Analyzer contains a counter of the successful analysis
type MultiAnalyzer struct {
	id     string
	total  int32
	counts map[string]int32
	fn     MultiAnalyzeFunc
}

// NewMultiAnalyzer creates a new analyzer with the given
// identifier and the given analyzer function
func NewMultiAnalyzer(id string, fn MultiAnalyzeFunc) *MultiAnalyzer {
	return &MultiAnalyzer{id: id, fn: fn, counts: make(map[string]int32)}
}

// Analyze performs analysis of the given string using
// analyzer function fn. If the input string matches
// analyzer rules i.e. the function fn returns true
// the counter is incremented by one
func (m *MultiAnalyzer) Analyze(s string) {
	if id, ok := m.fn(s); ok {
		m.counts[id]++
	}
}

// ID returns analyzer identifier
func (m *MultiAnalyzer) ID() string {
	return m.id
}

// Reset sets the counter to zero
func (m *MultiAnalyzer) Reset() {
	m.counts = make(map[string]int32)
}

// Counters returns the counter value i.e. number of successful matches
func (m *MultiAnalyzer) Counters() map[string]int32 {
	counters := make(map[string]int32)

	for k, v := range m.counts {
		counters[k] = v
	}

	return counters
}
