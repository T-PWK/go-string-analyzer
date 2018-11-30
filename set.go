package stranlz

import (
	"sort"
)

// Set is a collection of analyzers.
// An analyzer set contains identifier, which is used in analysis report.
type Set struct {
	id        string
	total     int32
	analyzers []Analyzer
}

type setByID []*Set

func (a setByID) Len() int           { return len(a) }
func (a setByID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a setByID) Less(i, j int) bool { return a[i].ID() < a[j].ID() }

// NewSet creates a new collection of analyzers with the given identifier
func NewSet(id string) *Set {
	return &Set{id: id}
}

// ID returns the set identifier.
func (s *Set) ID() string {
	return s.id
}

// IDs returns identifiers of the set analyzers
func (s *Set) IDs() []string {
	ids := make([]string, len(s.analyzers))

	for l, i := len(s.analyzers), 0; i < l; i++ {
		ids[i] = s.analyzers[i].ID()
	}

	return ids
}

// Analyze performs analysis of the given string against analyzers of ths set.
func (s *Set) Analyze(p string) {
	s.total++

	for _, a := range s.analyzers {
		a.Analyze(p)
	}
}

// AddSimple adds the given analyzer function fn to this collection
// with the given identifier id.
func (s *Set) AddSimple(id string, fn AnalyzeFunc) {
	s.AddAnalyzer(NewSimpleAnalyzer(id, fn))
}

// AddMulti adds the given analyzer function fn to this collection
// with the given identifier id.
func (s *Set) AddMulti(id string, fn MultiAnalyzeFunc) {
	s.AddAnalyzer(NewMultiAnalyzer(id, fn))
}

// AddAnalyzer adds the given analyzer to the collection of analyzers
func (s *Set) AddAnalyzer(a Analyzer) {
	i := s.findAnalyzerIndex(a.ID())

	if i == -1 {
		s.analyzers = append(s.analyzers, a)
	} else {
		s.analyzers[i] = a
	}

	sort.Sort(byID(s.analyzers))
}

func (s *Set) findAnalyzerIndex(id string) int {
	i := sort.Search(len(s.analyzers), func(i int) bool { return s.analyzers[i].ID() >= id })

	if i < len(s.analyzers) && s.analyzers[i].ID() == id {
		return i
	}
	return -1
}

func (s *Set) findAnalyzer(id string) Analyzer {
	i := s.findAnalyzerIndex(id)

	if i >= 0 {
		return s.analyzers[i]
	}
	return nil
}

// SimpleCounter returns counter value and a check flag of a simple analyzer identified by id.
// If there is no analyzer with the given id or it is not of type SimpleAnalyzer the returned
// value is 0 and the check flag is false.
func (s *Set) SimpleCounter(id string) (int32, bool) {
	anlz := s.findAnalyzer(id)

	switch v := anlz.(type) {
	case *SimpleAnalyzer:
		return v.Counter(), true
	default:
		return 0, false
	}
}

// MultiCounters returns counters value and a check flag of a multi analyzer identified by id.
// If there is no analyzer with the given id or it is not of type MultiAnalyzer the returned
// value is nil and the check flag is false.
func (s *Set) MultiCounters(id string) (map[string]int32, bool) {
	anlz := s.findAnalyzer(id)

	switch v := anlz.(type) {
	case *MultiAnalyzer:
		return v.Counters(), true
	default:
		return nil, false
	}
}

// Reset resets all analyzers
func (s *Set) Reset() {
	for _, a := range s.analyzers {
		a.Reset()
	}
}

// Size returns number of child analyzers
func (s *Set) Size() int {
	return len(s.analyzers)
}

// Total returns number of analysis performed
func (s *Set) Total() int32 {
	return s.total
}
