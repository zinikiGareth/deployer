package parser_test

import (
	"testing"

	"ziniki.org/deployer/deployer/internal/parser"
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func TestASingleIdIsJustThat(t *testing.T) {
	reporter := mockReporter()
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
	reporter := mockReporter()
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
	reporter := mockReporter()
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

func TestTwoIdsCanBeSeparatedByASpace(t *testing.T) {
	reporter := mockReporter()
	lex := parser.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(1, 1, "hello world")
	if len(toks) != 2 {
		t.Fatalf("not two args returned")
	}
	if toks[0].(pluggable.Identifier).Id() != "hello" {
		t.Fatalf("token 0 was not hello")
	}
	if toks[1].(pluggable.Identifier).Id() != "world" {
		t.Fatalf("token 1 was not world")
	}
}

func TestLeadingSpacesCauseAPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
		}
	}()
	reporter := mockReporter()
	lex := parser.NewLineLexicator(reporter, "test")
	lex.BlockedLine(1, 1, " hello")
	t.Fatalf("did not panic")
}

func TestTwoIdsAndASimpleStringCanBeSeparatedBySpaces(t *testing.T) {
	reporter := mockReporter()
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

func mockReporter() errors.ErrorRepI {
	return &errors.ErrorReporter{}
}
