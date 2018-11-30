package stranlz_test

import (
	"sort"
	"strings"
	"testing"

	"github.com/t-pwk/go-string-analyzer"
)

func TestListId(t *testing.T) {
	if id := createSet().ID(); id != "cid" {
		t.Errorf("Invalid collection ID: got %s, want %s", id, "cid")
	}
}

func TestListIDs(t *testing.T) {
	ids := createSet().IDs()

	if len(ids) != 2 || sort.SearchStrings(ids, "foo") < 0 || sort.SearchStrings(ids, "bar") < 0 {
		t.Errorf("Invalid collection ids: got %s, want [foo bar]", ids)
	}
}

func TestSetSize(t *testing.T) {
	if s := createSet().Size(); s != 2 {
		t.Errorf("Invalid collection size: got %d, want %d", s, 2)
	}
}

func TestListSameAnalyzerId(t *testing.T) {
	c := createSet()

	c.AddSimple("foo", func(s string) bool { return true })
	c.AddSimple("foo", func(s string) bool { return true })

	if s := c.Size(); s != 2 {
		t.Errorf("Invalid list size: got %d, want %d", s, 2)
	}
}

func TestSet(t *testing.T) {
	s := createSet()

	analyze(s, "foo", 2)
	analyze(s, "foobar", 2)
	analyze(s, "bar", 2)

	if s.Total() != 6 {
		t.Errorf("Invalid total value, want %d, got %d", 6, s.Total())
	}
}

func TestSetSimpleCounter(t *testing.T) {
	s := createSet()
	s.AddMulti("multi", stranlz.GroupByValue)

	analyze(s, "foo", 2)
	analyze(s, "bar", 2)

	if v, ok := s.SimpleCounter("foo"); v != 2 || !ok {
		t.Errorf("Invalid counter value, got %d, want %d", v, 4)
	}

	if _, ok := s.SimpleCounter("multi"); ok {
		t.Error("Invalid counter check flag, got `true`, want `false")
	}

	if _, ok := s.SimpleCounter("unknown"); ok {
		t.Error("Invalid counter check flag, got `true`, want `false")
	}
}

func TestSetMultiCounters(t *testing.T) {
	s := createSet()
	s.AddMulti("multi", stranlz.GroupByValue)

	analyze(s, "foo", 2)
	analyze(s, "bar", 2)

	if c, ok := s.MultiCounters("multi"); c == nil || !ok || c["foo"] != 2 || c["bar"] != 2 {
		t.Errorf("Invalid multi counters value, got %v, want %v", c, map[string]int{"foo": 2, "bar": 2})
	}

	if _, ok := s.MultiCounters("unknown"); ok {
		t.Error("Invalid counter check flag, got `true`, want `false")
	}

	if _, ok := s.MultiCounters("foo"); ok {
		t.Error("Invalid counter check flag, got `true`, want `false")
	}
}

func TestSettReset(t *testing.T) {
	s := createSet()
	s.AddMulti("multi", stranlz.GroupByValue)

	analyze(s, "foo", 10)

	s.Reset()

	if v, ok := s.SimpleCounter("foo"); v != 0 || !ok {
		t.Errorf("Invalid counter value, got %d, want %d", v, 0)
	}

	if c, ok := s.MultiCounters("multi"); c == nil || !ok || c["foo"] != 0 {
		t.Errorf("Invalid counter value, got %v, want %v", c, map[string]int32{"foo": 0})
	}
}

func createSet() *stranlz.Set {
	s := stranlz.NewSet("cid")
	s.AddSimple("foo", func(s string) bool { return strings.HasPrefix(s, "foo") })
	s.AddSimple("bar", func(s string) bool { return strings.HasPrefix(s, "bar") })

	return s
}

func analyze(c *stranlz.Set, s string, n int) {
	for i := 0; i < n; i++ {
		c.Analyze(s)
	}
}
