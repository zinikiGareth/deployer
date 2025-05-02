package lexicator_test

import (
	"testing"

	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func TestALineBeginningDoubleSlashSpaceIsIgnored(t *testing.T) {
	reporter, _ := mockReporter(t)
	lex := lexicator.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(lineOf("// hello, world"))
	if len(toks) != 0 {
		t.Fatalf("%d args returned, not 0", len(toks))
	}
}

func TestALineWithDoubleSlashSpaceIsTerminated(t *testing.T) {
	reporter, _ := mockReporter(t)
	lex := lexicator.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(lineOf("hello // , world"))
	if len(toks) != 1 {
		t.Fatalf("%d args returned, not 1", len(toks))
	}
	if toks[0].(pluggable.Identifier).Id() != "hello" {
		t.Fatalf("token was not hello")
	}
}

func TestALineBeginningTripleSlashIsIgnored(t *testing.T) {
	reporter, _ := mockReporter(t)
	lex := lexicator.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(lineOf("/// hello, world"))
	if len(toks) != 0 {
		t.Fatalf("%d args returned, not 0", len(toks))
	}
}
