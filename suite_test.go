package stranlz_test

import (
	"testing"

	"github.com/t-pwk/go-string-analyzer"
)

func TestNewSuite(t *testing.T) {
	s := stranlz.NewSuite()

	if s == nil {
		t.Fail()
	}
}

func TestAnalyzers(t *testing.T) {
	suite := stranlz.NewSuite()
	suite.AddSimple("set", "simple", fnTrue)
	suite.AddMulti("set", "multi", stranlz.GroupByValue)

	suite.Analyze("foo")
	suite.Analyze("bar")

	s, ok := suite.FindSet("set")

	if !ok {
		t.Error("Failed to fetch set")
	}

	if v, ok := s.SimpleCounter("simple"); v != 2 || !ok {
		t.Errorf("Failed to fetch simple analyzer or correct value, got %d, wanted 2", v)
	}

	if v, ok := s.MultiCounters("multi"); len(v) != 2 || !ok {
		t.Errorf("Failed to fetch multi analyzer or correct value, got %v, wanted %v", v, map[string]int32{"foo": 1, "bar": 1})
	}
}

func TestAnalyzeBySet(t *testing.T) {
	suite := stranlz.NewSuite()
	suite.AddSimple("s", "simple", fnTrue)
	suite.AddMulti("m", "multi", stranlz.GroupByValue)

	suite.AnalyzeBySet("s", "foo")
	suite.AnalyzeBySet("s", "bar")

	s, sok := suite.FindSet("s")
	m, mok := suite.FindSet("m")

	if !sok || !mok {
		t.Error("Failed to fetch sets")
	}

	if v, ok := s.SimpleCounter("simple"); v != 2 || !ok {
		t.Errorf("Failed to fetch simple analyzer or correct value, got %d, wanted 2", v)
	}

	if v, ok := m.MultiCounters("multi"); len(v) != 0 || !ok {
		t.Errorf("Failed to fetch multi analyzer or correct value, got %v, wanted %v", v, map[string]int32{})
	}
}

func TestReset(t *testing.T) {
	suite := stranlz.NewSuite()
	suite.AddSimple("s", "simple", fnTrue)
	suite.AddMulti("s", "multi", stranlz.GroupByValue)

	suite.AnalyzeBySet("s", "foo")
	suite.AnalyzeBySet("s", "bar")
	suite.AnalyzeBySet("s", "adm")

	suite.Reset()

	s, sok := suite.FindSet("s")

	if !sok {
		t.Error("Failed to fetch set")
	}

	if v, ok := s.SimpleCounter("simple"); v != 0 || !ok {
		t.Errorf("Failed to fetch simple analyzer or correct value, got %d, wanted 0", v)
	}

	if v, ok := s.MultiCounters("multi"); len(v) != 0 || !ok {
		t.Errorf("Failed to fetch multi analyzer or correct value, got %v, wanted %v", v, map[string]int32{})
	}
}
