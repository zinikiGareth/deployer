package parser_test

import (
	"testing"

	"ziniki.org/deployer/deployer/internal/parser"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func TestForSingleSlashWithSpaces(t *testing.T) {
	reporter, _ := mockReporter(t)
	lex := parser.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(1, 1, "hello / world")
	if len(toks) != 3 {
		t.Fatalf("%d args returned, not 3", len(toks))
	}
	if !toks[1].(pluggable.Operator).Is("/") {
		t.Fatalf("!op.Is(/)")
	}
	if toks[1].(pluggable.Operator).Op() != "/" {
		t.Fatalf("op was not /")
	}

}

func TestForSingleSlashWithoutSpaces(t *testing.T) {
	reporter, _ := mockReporter(t)
	lex := parser.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(1, 1, "hello/world")
	if len(toks) != 3 {
		t.Fatalf("%d args returned, not 3", len(toks))
	}
	if toks[0].(pluggable.Identifier).Id() != "hello" {
		t.Fatalf("tok0 != hello")
	}
	if !toks[1].(pluggable.Operator).Is("/") {
		t.Fatalf("!op.Is(/)")
	}
	if toks[1].(pluggable.Operator).Op() != "/" {
		t.Fatalf("op was not /")
	}
	if toks[2].(pluggable.Identifier).Id() != "world" {
		t.Fatalf("tok2 != world")
	}
}

func TestForEqualRightArrowWithoutSpaces(t *testing.T) {
	reporter, _ := mockReporter(t)
	lex := parser.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(1, 1, "hello=>world")
	if len(toks) != 3 {
		t.Fatalf("%d args returned, not 3", len(toks))
	}
	if toks[0].(pluggable.Identifier).Id() != "hello" {
		t.Fatalf("tok0 != hello")
	}
	if !toks[1].(pluggable.Operator).Is("=>") {
		t.Fatalf("!op.Is(=>)")
	}
	if toks[1].(pluggable.Operator).Op() != "=>" {
		t.Fatalf("op was not =>")
	}
	if toks[2].(pluggable.Identifier).Id() != "world" {
		t.Fatalf("tok2 != world")
	}
}

func TestForMinusLeftArrowWithoutSpaces(t *testing.T) {
	reporter, _ := mockReporter(t)
	lex := parser.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(1, 1, "hello<-'world'")
	if len(toks) != 3 {
		t.Fatalf("%d args returned, not 3", len(toks))
	}
	if toks[0].(pluggable.Identifier).Id() != "hello" {
		t.Fatalf("tok0 != hello")
	}
	if !toks[1].(pluggable.Operator).Is("<-") {
		t.Fatalf("!op.Is(<-)")
	}
	if toks[1].(pluggable.Operator).Op() != "<-" {
		t.Fatalf("op was not <-")
	}
	if toks[2].(pluggable.String).Text() != "world" {
		t.Fatalf("tok2 != world")
	}
}
