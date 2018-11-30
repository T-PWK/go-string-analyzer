package stranlz_test

import (
	"bytes"
	"testing"

	"github.com/t-pwk/go-string-analyzer"
)

func TestReport(t *testing.T) {
	suit := stranlz.NewSuite()
	suit.AddSimple("simple", "always-true-1", fnTrue)
	suit.AddSimple("simple", "always-true-2", fnTrue)
	suit.AddMulti("multi", "group-by-value", stranlz.GroupByValue)

	for i := 0; i < 10; i++ {
		suit.Analyze("test")
	}

	report := stranlz.NewReport(suit)

	if report == nil {
		t.Error("Failed to create report")
	}

	report.Header = "Test report"
	report.Totals = true
	report.Pcts = true

	writer := bytes.NewBufferString("")

	report.Print(writer)
}
