package stranlz

import (
	"sort"
)

// Suite contains a collection of analyzer sets.
type Suite struct {
	total int32
	sets  []*Set
}

// NewSuite creates a new test suite
func NewSuite() *Suite {
	return new(Suite)
}

// AddSimple adds simple analyzer with identifier id to the set with sid identifier.
func (s *Suite) AddSimple(sid string, id string, fn AnalyzeFunc) {
	s.AddAnalyzer(sid, NewSimpleAnalyzer(id, fn))
}

// AddMulti adds multi-value analyzer with identifier id to the set with sid identifier.
func (s *Suite) AddMulti(sid string, id string, fn MultiAnalyzeFunc) {
	s.AddAnalyzer(sid, NewMultiAnalyzer(id, fn))
}

// AddAnalyzer adds the analyzer a to the set identified by id.
func (s *Suite) AddAnalyzer(id string, a Analyzer) {
	set, ok := s.FindSet(id)

	if !ok {
		set = NewSet(id)
		s.sets = append(s.sets, set)
		sort.Sort(setByID(s.sets))
	}

	set.AddAnalyzer(a)
}

// Analyze performs analysis of the string v through all analyzers.
func (s *Suite) Analyze(v string) {
	s.total++

	for _, c := range s.sets {
		c.Analyze(v)
	}
}

// Reset resets all sets
func (s *Suite) Reset() {
	for _, set := range s.sets {
		set.Reset()
	}
}

// AnalyzeBySet performs analysis of the string v through all analyzers in the set identified by id.
func (s *Suite) AnalyzeBySet(id string, v string) {
	set, ok := s.FindSet(id)

	if ok {
		s.total++
		set.Analyze(v)
	}
}

// FindSet returns a set with the given identifier.
// The ok flag is set to true if an the analyzer is found; otherwise false.
func (s *Suite) FindSet(id string) (*Set, bool) {
	i := sort.Search(len(s.sets), func(i int) bool { return s.sets[i].ID() >= id })

	if i < len(s.sets) && s.sets[i].ID() == id {
		return s.sets[i], true
	}
	return nil, false
}
