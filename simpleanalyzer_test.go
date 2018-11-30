package stranlz_test

import (
	"testing"

	stranlz "github.com/t-pwk/go-string-analyzer"
)

func TestSimpleAnalyzerReport(t *testing.T) {
	a := stranlz.NewSimpleAnalyzer("id", func(s string) bool { return true })

	for i := 1; i <= 10; i++ {
		a.Analyze("")

		if c, to := a.Counter(), a.Total(); c != int32(i) || to != int32(i) {
			t.Errorf("Incorrect number of successful analysis, got %d, want %d", c, i)
		}
	}
}
func TestSimpleAnalyzerReset(t *testing.T) {
	a := stranlz.NewSimpleAnalyzer("id", func(s string) bool { return true })

	a.Analyze("")
	a.Analyze("")
	a.Reset()

	if c := a.Counter(); c != 0 {
		t.Errorf("Incorrect number of successful analysis, got %d, want 0", c)
	}
}
