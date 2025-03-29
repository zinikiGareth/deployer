package parser_test

import (
	"testing"

	"ziniki.org/deployer/deployer/internal/parser"
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func TestASingleIdIsJustThat(t *testing.T) {
	reporter, _ := mockReporter(t)
	lex := parser.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(1, 1, "hello")
	if len(toks) != 1 {
		t.Fatalf("not exactly one arg returned")
	}
	if toks[0].(pluggable.Identifier).Id() != "hello" {
		t.Fatalf("token was not hello")
	}
}

func TestAnIdCanHaveDots(t *testing.T) {
	reporter, _ := mockReporter(t)
	lex := parser.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(1, 1, "hello.world")
	if len(toks) != 1 {
		t.Fatalf("not exactly one arg returned")
	}
	if toks[0].(pluggable.Identifier).Id() != "hello.world" {
		t.Fatalf("token was not hello.world")
	}
}

func TestAStringCanBeFoundBetweenDoubleQuotes(t *testing.T) {
	reporter, _ := mockReporter(t)
	lex := parser.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(1, 1, "\"hello, world\"")
	if len(toks) != 1 {
		t.Fatalf("not one arg returned")
	}
	str, ok := toks[0].(pluggable.String)
	if !ok {
		t.Fatalf("token 0 was not a string")
	}
	if str.Text() != "hello, world" {
		t.Fatalf("token 0 was not the string \"hello, world\" but %s", str)
	}
}

func TestAStringCanIncludeANestedDQPairBetweenDoubleQuotes(t *testing.T) {
	reporter, _ := mockReporter(t)
	lex := parser.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(1, 1, "\"hello, \"\"world\"\"\"")
	if len(toks) != 1 {
		t.Fatalf("not one arg returned but %v", toks)
	}
	str, ok := toks[0].(pluggable.String)
	if !ok {
		t.Fatalf("token 0 was not a string")
	}
	if str.Text() != "hello, \"world\"" {
		t.Fatalf("token 0 was not the string \"hello, \"world\"\" but %s", str)
	}
}

func TestAStringCanBeFoundBetweenSingleQuotes(t *testing.T) {
	reporter, _ := mockReporter(t)
	lex := parser.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(1, 1, "'hello, world'")
	if len(toks) != 1 {
		t.Fatalf("not one arg returned")
	}
	str, ok := toks[0].(pluggable.String)
	if !ok {
		t.Fatalf("token 0 was not a string")
	}
	if str.Text() != "hello, world" {
		t.Fatalf("token 0 was not the string \"hello, world\" but %s", str)
	}
}

func TestAStringCanIncludeANestedSQPairBetweenSingleQuotes(t *testing.T) {
	reporter, _ := mockReporter(t)
	lex := parser.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(1, 1, "'hello, ''world'''")
	if len(toks) != 1 {
		t.Fatalf("not one arg returned but %v", toks)
	}
	str, ok := toks[0].(pluggable.String)
	if !ok {
		t.Fatalf("token 0 was not a string")
	}
	if str.Text() != "hello, 'world'" {
		t.Fatalf("token 0 was not the string \"hello, 'world'\" but %s", str)
	}
}

func TestAStringMustBeTerminated(t *testing.T) {
	reporter, sink := mockReporter(t)
	tx := "\"hello, world"
	sink.Expect(1, 1, tx, "unterminated string")
	lex := parser.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(1, 1, tx)
	if toks != nil {
		t.Fatalf("expected nil")
	}
}

func TestAStringMustNotEndWithNetedQuote(t *testing.T) {
	reporter, sink := mockReporter(t)
	tx := "\"hello, world\"\""
	sink.Expect(1, 1, tx, "unterminated string")
	lex := parser.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(1, 1, tx)
	if toks != nil {
		t.Fatalf("expected nil")
	}
}

// test of ids/strings/symbols butted up against each other
// id"hello" is an error - must be a space
// id<-x is probably fine, though, because we want x*3, and think about calc in css

func TestTwoIdsCanBeSeparatedByASpace(t *testing.T) {
	reporter, _ := mockReporter(t)
	lex := parser.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(1, 1, "hello world")
	if len(toks) != 2 {
		t.Fatalf("not two args returned")
	}
	if toks[0].(pluggable.Identifier).Id() != "hello" {
		t.Fatalf("token 0 was not hello, but %s", toks[0])
	}
	if toks[1].(pluggable.Identifier).Id() != "world" {
		t.Fatalf("token 1 was not world, but %s", toks[1])
	}
}

func TestLeadingSpacesCauseAPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	reporter, _ := mockReporter(t)
	lex := parser.NewLineLexicator(reporter, "test")
	lex.BlockedLine(1, 1, " hello")
	t.Fatalf("did not panic")
}

func TestTwoIdsAndASimpleStringCanBeSeparatedBySpaces(t *testing.T) {
	reporter, _ := mockReporter(t)
	lex := parser.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(1, 1, "ensure test.S3.Bucket \"org.ziniki.launch_bucket\"")
	if len(toks) != 3 {
		t.Fatalf("not three args returned")
	}
	if toks[0].(pluggable.Identifier).Id() != "ensure" {
		t.Fatalf("token 0 was not ensure")
	}
	if toks[1].(pluggable.Identifier).Id() != "test.S3.Bucket" {
		t.Fatalf("token 1 was not test.S3.Bucket")
	}
	if toks[2].(pluggable.String).Text() != "org.ziniki.launch_bucket" {
		t.Fatalf("token 2 was not org.ziniki.launch_bucket")
	}
}

// TODO: numbers of all forms
// TODO: symbols

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
	if es.line != line {
		s.t.Fatalf("was line %d not %d", line, es.line)
	}
	if es.ind != ind {
		s.t.Fatalf("was line %d not %d", ind, es.ind)
	}
	if es.text != text {
		s.t.Fatalf("text was '%s' not '%s'", text, es.text)
	}
	if es.msg != msg {
		s.t.Fatalf("msg was '%s' not '%s'", msg, es.msg)
	}
}

func mockReporter(t *testing.T) (errors.ErrorRepI, *mockSink) {
	sink := &mockSink{t: t}
	return errors.NewErrorReporter(sink), sink
}
