package parser_test

import (
	"testing"

	"ziniki.org/deployer/deployer/internal/parser"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func Test0IsANumber(t *testing.T) {
	reporter, _ := mockReporter(t)
	lex := parser.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(1, 1, "0")
	if len(toks) != 1 {
		t.Fatalf("%d args returned, not 1", len(toks))
	}
	if toks[0].(pluggable.Number).Value() != 0 {
		t.Fatalf("val != 0")
	}
}


func TestWeCanParse24hours(t *testing.T) {
	reporter, _ := mockReporter(t)
	lex := parser.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(1, 1, "24 hours")
	if len(toks) != 2 {
		t.Fatalf("%d args returned, not 2", len(toks))
	}
	if toks[0].(pluggable.Number).Value() != 24 {
		t.Fatalf("val != 24")
	}
}
