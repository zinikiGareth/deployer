package lexicator_test

import (
	"fmt"
	"testing"

	"ziniki.org/deployer/deployer/pkg/errors"
)

type errorStruct struct {
	line, ind, offset int
	text, msg         string
}

type mockSink struct {
	t      *testing.T
	errors []errorStruct
}

func (s *mockSink) Expect(line, ind, offset int, text, msg string) {
	s.errors = append(s.errors, errorStruct{line: line, ind: ind, offset: offset, text: text, msg: msg})
}

func (s *mockSink) Report(line *errors.Location, msg string) {
	if len(s.errors) == 0 {
		s.t.Fatalf("unexpected error: " + msg)
	}
	es := s.errors[0]
	s.errors = s.errors[1:]
	if es.msg != msg {
		s.t.Fatalf("msg was '%s' not '%s'", msg, es.msg)
	}
	if es.line != line.Line.Line {
		s.t.Fatalf("was line %d not %d", line.Line.Line, es.line)
	}
	if es.ind != line.Line.Indent {
		s.t.Fatalf("was ind %d not %d", line.Line.Indent, es.ind)
	}
	// TODO: offset??
	if es.text != line.Line.Text {
		s.t.Fatalf("text was '%s' not '%s'", line.Line.Text, es.text)
	}
}

func (s *mockSink) Reportf(loc *errors.Location, format string, args ...any) {
	s.Report(loc, fmt.Sprintf(format, args...))
}

func (s *mockSink) HasErrors() bool {
	return len(s.errors) > 0
}

func mockReporter(t *testing.T) (errors.ErrorRepI, *mockSink) {
	sink := &mockSink{t: t}
	return errors.NewErrorReporter(sink), sink
}
