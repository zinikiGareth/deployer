package lexicator_test

import (
	"testing"

	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/testhelpers"
)

func Test0IsANumber(t *testing.T) {
	reporter, _ := testhelpers.MockReporter(t)
	tools := pluggable.NewTools(reporter, nil, nil, nil, nil)
	lex := lexicator.NewLineLexicator(tools, "test")
	toks := lex.BlockedLine(lineOf("0"))
	if len(toks) != 1 {
		t.Fatalf("%d args returned, not 1", len(toks))
	}
	if toks[0].(pluggable.Number).Value() != 0 {
		t.Fatalf("val != 0")
	}
}

func TestPiIsANumber(t *testing.T) {
	reporter, _ := testhelpers.MockReporter(t)
	tools := pluggable.NewTools(reporter, nil, nil, nil, nil)
	lex := lexicator.NewLineLexicator(tools, "test")
	toks := lex.BlockedLine(lineOf("3.14"))
	if len(toks) != 1 {
		t.Fatalf("%d args returned, not 1", len(toks))
	}
	if toks[0].(pluggable.Number).Value() != 3.14 {
		t.Fatalf("val != π")
	}
}

func TestASimpleHexNumber(t *testing.T) {
	reporter, _ := testhelpers.MockReporter(t)
	tools := pluggable.NewTools(reporter, nil, nil, nil, nil)
	lex := lexicator.NewLineLexicator(tools, "test")
	toks := lex.BlockedLine(lineOf("0xff"))
	if len(toks) != 1 {
		t.Fatalf("%d args returned, not 1", len(toks))
	}
	if toks[0].(pluggable.Number).Value() != 255 {
		t.Fatalf("val != 255")
	}
}

func TestExponentNumber(t *testing.T) {
	reporter, _ := testhelpers.MockReporter(t)
	tools := pluggable.NewTools(reporter, nil, nil, nil, nil)
	lex := lexicator.NewLineLexicator(tools, "test")
	toks := lex.BlockedLine(lineOf("2.7e-3"))
	if len(toks) != 1 {
		t.Fatalf("%d args returned, not 1", len(toks))
	}
	if toks[0].(pluggable.Number).Value() != 0.0027 {
		t.Fatalf("val != 0.0027")
	}
}

func TestWeCanParse24hours(t *testing.T) {
	reporter, _ := testhelpers.MockReporter(t)
	tools := pluggable.NewTools(reporter, nil, nil, nil, nil)
	lex := lexicator.NewLineLexicator(tools, "test")
	toks := lex.BlockedLine(lineOf("24 hours"))
	if len(toks) != 2 {
		t.Fatalf("%d args returned, not 2", len(toks))
	}
	if toks[0].(pluggable.Number).Value() != 24 {
		t.Fatalf("val != 24")
	}
}
