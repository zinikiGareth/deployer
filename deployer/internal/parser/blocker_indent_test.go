package parser_test

import (
	"fmt"
	"log"
	"testing"

	"ziniki.org/deployer/deployer/internal/parser"
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/interpreters"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func TestCommentLinesAreDiscarded(t *testing.T) {
	blockerTest([]line{
		{lineNo: 10, ignore: true, indent: "", text: "hello"},
	})
}

func TestTheFirstIndentedLineIsAlwaysAccepted(t *testing.T) {
	blockerTest([]line{
		{lineNo: 10, lineIndent: 1, indent: "\t", text: "hello"},
	})
}

func TestConsecutiveLinesAtTheSameIndentAreAccepted(t *testing.T) {
	blockerTest([]line{
		{lineNo: 10, lineIndent: 1, indent: "\t", text: "hello"},
		{lineNo: 15, lineIndent: 1, indent: "\t", text: "goodbye"},
	})
}

func TestAnIdentedLineIsPassedToAnInnerScope(t *testing.T) {
	blockerTest([]line{
		{lineNo: 10, lineIndent: 1, indent: "\t", text: "hello",
			inner: innerBlock([]line{
				{lineNo: 15, lineIndent: 2, indent: "\t\t", text: "goodbye"},
			}),
		},
	})
}

type line struct {
	lineNo     int
	ignore     bool
	lineIndent int
	indent     string
	text       string
	seen       bool
	inner      *tmp
}

type tmp struct {
	sink  errors.ErrorSink
	lines []line
}

func (t *tmp) applySink(sink errors.ErrorSink) {
	t.sink = sink
	for _, l := range t.lines {
		if l.inner != nil {
			l.inner.applySink(sink)
		}
	}
}

func (t *tmp) BlockedLine(lineNo, lenIndent int, text string) pluggable.ProvideBlockedLine {
	for ln := range t.lines {
		l := &t.lines[ln]
		if l.lineNo == lineNo {
			if l.ignore {
				panic(fmt.Sprintf("should ignore line %d", l.lineNo))
			} else if l.lineIndent != lenIndent {
				panic(fmt.Sprintf("line %d indent wrong: %d != %d", l.lineNo, lenIndent, l.lineIndent))
			} else if l.text != text {
				panic(fmt.Sprintf("line %d wrong: %s != %s", l.lineNo, text, l.text))
			}
			log.Println(text)
			l.seen = true
			if l.inner != nil {
				return l.inner
			} else {
				return interpreters.DisallowInnerScope(t.sink)
			}
		}
	}
	panic(fmt.Sprintf("line %d was unexpected: %s", lineNo, text))
}

type testSink struct {
}

func (s *testSink) Report(lineNo int, indent int, lineText string, msg string) {

}

func blockerTest(lines []line) {
	sink := &testSink{}
	mock := innerBlock(lines)
	mock.applySink(sink)
	reporter := errors.NewErrorReporter(sink)
	blocker := parser.NewBlocker(reporter, mock)
	for _, b := range mock.lines {
		blocker.HaveLine(b.lineNo, b.indent+b.text)
	}

	for _, l := range lines {
		if l.ignore == l.seen {
			panic(fmt.Sprintf("incorrect seen: %d seen: %v ignore: %v", l.lineNo, l.seen, l.ignore))
		}
	}
}

func innerBlock(lines []line) *tmp {
	atLeast := 0
	for _, l := range lines {
		if l.lineNo <= atLeast {
			panic(fmt.Sprintf("non-monotonic lineNo: %d", l.lineNo))
		}
	}
	return &tmp{lines: lines}
}
