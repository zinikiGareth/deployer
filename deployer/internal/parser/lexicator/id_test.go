package lexicator_test

import (
	"testing"

	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/testhelpers"
)

func TestASingleIdIsJustThat(t *testing.T) {
	reporter, _ := testhelpers.MockReporter(t)
	lex := lexicator.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(lineOf("hello"))
	if len(toks) != 1 {
		t.Fatalf("not exactly one arg returned")
	}
	if toks[0].(pluggable.Identifier).Id() != "hello" {
		t.Fatalf("token was not hello")
	}
}

func TestAnIdCanHaveDots(t *testing.T) {
	reporter, _ := testhelpers.MockReporter(t)
	lex := lexicator.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(lineOf("hello.world"))
	if len(toks) != 1 {
		t.Fatalf("not exactly one arg returned")
	}
	if toks[0].(pluggable.Identifier).Id() != "hello.world" {
		t.Fatalf("token was not hello.world")
	}
}

// test of ids/strings/symbols butted up against each other
// id"hello" is an error - must be a space
// id<-x is probably fine, though, because we want x*3, and think about calc in css

func TestTwoIdsCanBeSeparatedByASpace(t *testing.T) {
	reporter, _ := testhelpers.MockReporter(t)
	lex := lexicator.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(lineOf("hello world"))
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
	reporter, _ := testhelpers.MockReporter(t)
	lex := lexicator.NewLineLexicator(reporter, "test")
	lex.BlockedLine(lineOf(" hello"))
	t.Fatalf("did not panic")
}
