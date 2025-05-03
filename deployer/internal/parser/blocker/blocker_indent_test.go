package blocker_test

import (
	"fmt"
	"testing"

	"ziniki.org/deployer/deployer/internal/parser/blocker"
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
		{lineNo: 10, lineIndent: 0, indent: "\t", text: "hello"},
	})
}

func TestConsecutiveLinesAtTheSameIndentAreAccepted(t *testing.T) {
	blockerTest([]line{
		{lineNo: 10, lineIndent: 0, indent: "\t", text: "hello"},
		{lineNo: 15, lineIndent: 0, indent: "\t", text: "goodbye"},
	})
}

func TestAnIdentedLineIsPassedToAnInnerScope(t *testing.T) {
	blockerTest([]line{
		{lineNo: 10, lineIndent: 0, indent: "\t", text: "hello",
			inner: innerBlock([]line{
				{lineNo: 15, lineIndent: 1, indent: "\t\t", text: "goodbye"},
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

func (t *tmp) HaveTokens(_ *pluggable.Tools, toks []pluggable.Token) pluggable.Interpreter {
	tok := toks[0].(*LineToken)
	lineNo := tok.Loc().Line.Line
	lenIndent := tok.Loc().Line.Indent
	text := tok.tx
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
			l.seen = true
			if l.inner != nil {
				return l.inner
			} else {
				return interpreters.DisallowInnerScope()
			}
		}
	}
	panic(fmt.Sprintf("line %d was unexpected: %s", lineNo, text))
}

func (b *tmp) Completed(tools *pluggable.Tools) {
}

type LineToken struct {
	loc *errors.Location
	tx  string
}

func (t *LineToken) Loc() *errors.Location {
	return t.loc
}

func (t LineToken) String() string {
	return "LIIIINEEETOOOOKEN"
}

type testLex struct {
}

func (l *testLex) BlockedLine(line *errors.LineLoc) []pluggable.Token {
	loc := line.Location(0)
	return []pluggable.Token{&LineToken{loc: loc, tx: line.Text}}
}

type testSink struct {
	errorCount int
}

func (s *testSink) Report(loc *errors.Location, msg string) {
	s.errorCount++
}

func (s *testSink) Reportf(loc *errors.Location, format string, args ...any) {
	s.Report(loc, fmt.Sprintf(format, args...))
}

func (s *testSink) HasErrors() bool {
	return s.errorCount > 0
}

func blockerTest(lines []line) {
	mocklex := &testLex{}
	sink := &testSink{}
	mock := innerBlock(lines)
	mock.applySink(sink)
	reporter := errors.NewErrorReporter(sink)
	tools := pluggable.NewTools(reporter, nil, nil)
	blocker := blocker.NewBlocker(tools, mocklex, mock)
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
