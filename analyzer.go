package stranlz

// AnalyzeFunc type is a function that reports whether the string s
// matches rules of the given implementations
type AnalyzeFunc func(s string) bool

// MultiAnalyzeFunc type is a function that reports whether
// the input string s matches rules of the given implementations.
// It also returns specific key associated with the input string s.
type MultiAnalyzeFunc func(s string) (string, bool)

// Analyzer is an interface for string analyzer components
type Analyzer interface {
	ID() string
	Analyze(s string)
	Reset()
}

type byID []Analyzer

func (a byID) Len() int           { return len(a) }
func (a byID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byID) Less(i, j int) bool { return a[i].ID() < a[j].ID() }
