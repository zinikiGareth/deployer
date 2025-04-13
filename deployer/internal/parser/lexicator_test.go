package parser_test

import (
	"fmt"
	"testing"

	"ziniki.org/deployer/deployer/pkg/errors"
)

type errorStruct struct {
	line, ind int
	text, msg string
}

type mockSink struct {
	t      *testing.T
	errors []errorStruct
}

func (s *mockSink) Expect(line, ind int, text, msg string) {
	s.errors = append(s.errors, errorStruct{line: line, ind: ind, text: text, msg: msg})
}

func (s *mockSink) Report(line, ind int, text, msg string) {
	if len(s.errors) == 0 {
		s.t.Fatalf("unexpected error: " + msg)
	}
	es := s.errors[0]
	s.errors = s.errors[1:]
	if es.msg != msg {
		s.t.Fatalf("msg was '%s' not '%s'", msg, es.msg)
	}
	if es.line != line {
		s.t.Fatalf("was line %d not %d", line, es.line)
	}
	if es.ind != ind {
		s.t.Fatalf("was ind %d not %d", ind, es.ind)
	}
	if es.text != text {
		s.t.Fatalf("text was '%s' not '%s'", text, es.text)
	}
}

func (s *mockSink) Reportf(lineNo int, indent int, lineText string, format string, args ...any) {
	s.Report(lineNo, indent, lineText, fmt.Sprintf(format, args...))
}

func (s *mockSink) HasErrors() bool {
	return len(s.errors) > 0
}

func mockReporter(t *testing.T) (errors.ErrorRepI, *mockSink) {
	sink := &mockSink{t: t}
	return errors.NewErrorReporter(sink), sink
}
