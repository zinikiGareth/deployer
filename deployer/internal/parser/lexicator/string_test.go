package lexicator_test

import (
	"testing"

	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/testhelpers"
)

func TestAStringCanBeFoundBetweenDoubleQuotes(t *testing.T) {
	reporter, _ := testhelpers.MockReporter(t)
	lex := lexicator.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(lineOf("\"hello, world\""))
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
	reporter, _ := testhelpers.MockReporter(t)
	lex := lexicator.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(lineOf("\"hello, \"\"world\"\"\""))
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
	reporter, _ := testhelpers.MockReporter(t)
	lex := lexicator.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(lineOf("'hello, world'"))
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
	reporter, _ := testhelpers.MockReporter(t)
	lex := lexicator.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(lineOf("'hello, ''world'''"))
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
	reporter, sink := testhelpers.MockReporter(t)
	tx := "\"hello, world"
	sink.Expect(1, 1, 0, tx, "unterminated string")
	lex := lexicator.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(lineOf(tx))
	if toks != nil {
		t.Fatalf("expected nil")
	}
}

func TestAStringMustNotEndWithNetedQuote(t *testing.T) {
	reporter, sink := testhelpers.MockReporter(t)
	tx := "\"hello, world\"\""
	sink.Expect(1, 1, 0, tx, "unterminated string")
	lex := lexicator.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(lineOf(tx))
	if toks != nil {
		t.Fatalf("expected nil")
	}
}

func TestThereMustBeASpaceBetweenIDAndAString(t *testing.T) {
	reporter, sink := testhelpers.MockReporter(t)
	tx := "system'hello'"
	lex := lexicator.NewLineLexicator(reporter, "test")
	sink.Expect(1, 1, 6, tx, "space required after identifier before string")
	toks := lex.BlockedLine(lineOf(tx))
	if toks != nil {
		t.Fatalf("expected nil")
	}
}

func TestThereMustBeASpaceBetweenAStringAndAnID(t *testing.T) {
	reporter, sink := testhelpers.MockReporter(t)
	tx := "'hello'system"
	lex := lexicator.NewLineLexicator(reporter, "test")
	sink.Expect(1, 1, 7, tx, "space required after string before identifier")
	toks := lex.BlockedLine(lineOf(tx))
	if toks != nil {
		t.Fatalf("expected nil")
	}
}

func TestTwoIdsAndASimpleStringCanBeSeparatedBySpaces(t *testing.T) {
	reporter, _ := testhelpers.MockReporter(t)
	lex := lexicator.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(lineOf("ensure test.S3.Bucket \"org.ziniki.launch_bucket\""))
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
