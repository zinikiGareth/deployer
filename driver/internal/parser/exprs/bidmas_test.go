package exprs_test

import (
	"testing"

	"ziniki.org/deployer/driver/internal/basicmath"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

func TestSimpleMultiplication(t *testing.T) {
	p, h := makeParser(t)
	basicmath.RegisterAll(h.Tools)
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "2*3")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 3 {
		t.Fatalf("expected three tokens, not %d", len(toks))
	}
	e, ok := p.Parse(nil, toks)
	if !ok {
		t.Fatalf("error parsing %s", line.Text)
	}
	if e.ShortDescription() != "mult [Number[2.000000],Number[3.000000]]" {
		t.Fatalf("incorrect parsing: %s", e.ShortDescription())
	}
}

// NOTE: this is a test that we DO NOT support obvious HOFs
// Do we want that?
func TestMultiplyIsInvalidByItself(t *testing.T) {
	p, h := makeParser(t)
	h.Sink.Expect(1, 1, 0, "*", "* requires left operand")
	basicmath.RegisterAll(h.Tools)
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "*")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 1 {
		t.Fatalf("expected one token, not %d", len(toks))
	}
	_, ok := p.Parse(nil, toks)
	if ok {
		t.Fatalf("parse erroneously reported successful")
	}
}

func TestMultiplyIsInvalidWithLHS(t *testing.T) {
	p, h := makeParser(t)
	h.Sink.Expect(1, 1, 0, "5*", "* requires right operand")
	basicmath.RegisterAll(h.Tools)
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "5*")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 2 {
		t.Fatalf("expected two tokens, not %d", len(toks))
	}
	_, ok := p.Parse(nil, toks)
	if ok {
		t.Fatalf("parse erroneously reported successful")
	}
}

func TestMultiplyIsInvalidWithRHS(t *testing.T) {
	p, h := makeParser(t)
	h.Sink.Expect(1, 1, 0, "*7", "* requires left operand")
	basicmath.RegisterAll(h.Tools)
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "*7")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 2 {
		t.Fatalf("expected two tokens, not %d", len(toks))
	}
	_, ok := p.Parse(nil, toks)
	if ok {
		t.Fatalf("parse erroneously reported successful")
	}
}

func TestMultiplyEval(t *testing.T) {
	p, h := makeParser(t)
	// h.Sink.Expect(1, 1, 0, "*7", "* requires left operand")
	basicmath.RegisterAll(h.Tools)
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "5*7")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 3 {
		t.Fatalf("expected three token, not %d", len(toks))
	}
	e, ok := p.Parse(nil, toks)
	if !ok {
		t.Fatalf("error parsing %s", line.Text)
	}
	if e.ShortDescription() != "mult [Number[5.000000],Number[7.000000]]" {
		t.Fatalf("incorrect parsing: %s", e.ShortDescription())
	}
	val := e.Eval(nil)
	num, ok := val.(float64)
	if !ok {
		t.Fatalf("Eval did not return a float64 but %T", val)
	}
	if num != 35.0 {
		t.Fatalf("5*7 was not 35 but %f", num)
	}
}
