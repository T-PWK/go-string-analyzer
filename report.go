package stranlz

import (
	"fmt"
	"io"
	"math"
	"sort"
	"strings"
)

// Report holds analyzers suite details.
type Report struct {
	Pcts       bool
	Totals     bool
	NamesSpace int
	NumsSpace  int
	Header     string

	sects []*section
}

// section holds analyzers set details.
type section struct {
	id    string
	total int32
	recs  []*record
	sects []*section
}

type record struct {
	id    string
	value int32
}

type recordsSort struct {
	r    []*record
	less func(i, j *record) bool
}

func (r recordsSort) Len() int           { return len(r.r) }
func (r recordsSort) Swap(i, j int)      { r.r[i], r.r[j] = r.r[j], r.r[i] }
func (r recordsSort) Less(i, j int) bool { return r.less(r.r[i], r.r[j]) }

func sortRecordByID(i, j *record) bool    { return i.id < j.id }
func sortRecordByValue(i, j *record) bool { return i.value < j.value }

// NewReport generate suite report for the suite g.
func NewReport(s *Suite) *Report {
	report := new(Report)

	for _, set := range s.sets {
		report.sects = append(report.sects, newReportSection(set))
	}

	report.NamesSpace, report.NumsSpace = reportMaxValues(report)

	return report
}

// NewSetReport generate set report for the set s.
func newReportSection(s *Set) *section {

	sec := &section{id: s.id, total: s.total}

	for _, a := range s.analyzers {
		switch v := a.(type) {
		case *SimpleAnalyzer:
			sec.recs = append(sec.recs, &record{id: v.id, value: v.count})
		case *MultiAnalyzer:
			subsec := &section{id: v.id, total: s.total}
			sec.sects = append(sec.sects, subsec)
			for k, v := range v.counts {
				subsec.recs = append(subsec.recs, &record{id: k, value: v})
			}
			sort.Sort(recordsSort{r: subsec.recs, less: sortRecordByValue})
		}
	}

	return sec
}

// Print prints the report to the writer w.
func (r *Report) Print(w io.Writer) {

	if len(r.Header) > 0 {
		fmt.Fprintf(w, "%s:\n\n", r.Header)
	}

	for _, sec := range r.sects {
		fmt.Fprintf(w, " - %s", sec.id)

		if r.Totals {
			fmt.Fprintf(w, " (total: %d)", sec.total)
		}

		fmt.Fprintln(w)
		fmt.Fprintf(w, "   %s\n", strings.Repeat("-", r.NamesSpace+2))
		printSection(w, r, sec)

		if len(sec.sects) > 0 {

			for _, s := range sec.sects {
				fmt.Fprintf(w, "\n   ** %s **\n", s.id)
				printSection(w, r, s)
			}
		}

		fmt.Fprintln(w)
	}
}

func sectionMaxValues(n int, v int, s *section) (int, int) {
	for _, r := range s.recs {
		n = max(n, len(r.id))
		v = max(v, max(int(math.Log10(float64(r.value)))+1, v))

		for _, s := range s.sects {
			n, v = sectionMaxValues(n, v, s)
		}
	}

	return n, v
}

func reportMaxValues(r *Report) (int, int) {

	var n, v int

	for _, s := range r.sects {
		n, v = sectionMaxValues(n, v, s)
	}

	return n, v
}

func max(i int, j int) int {
	if i > j {
		return i
	}
	return j
}

func printSection(w io.Writer, r *Report, sec *section) {
	for _, rec := range sec.recs {
		fmt.Fprintf(w, fmt.Sprintf("   - %%-%ds : %%%dd", r.NamesSpace, r.NumsSpace), rec.id, rec.value)

		if r.Pcts && sec.total > 0 {
			fmt.Fprintf(w, " (%.2f%%)", float32(rec.value)*100.0/float32(sec.total))
		}

		fmt.Fprintln(w)
	}
}
