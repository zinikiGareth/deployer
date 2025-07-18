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

func TestSimpleAddition(t *testing.T) {
	p, h := makeParser(t)
	basicmath.RegisterAll(h.Tools)
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "1+6")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 3 {
		t.Fatalf("expected three tokens, not %d", len(toks))
	}
	e, ok := p.Parse(nil, toks)
	if !ok {
		t.Fatalf("error parsing %s", line.Text)
	}
	if e.ShortDescription() != "add [Number[1.000000],Number[6.000000]]" {
		t.Fatalf("incorrect parsing: %s", e.ShortDescription())
	}
}

func TestPlusIsInvalidByItself(t *testing.T) {
	p, h := makeParser(t)
	h.Sink.Expect(1, 1, 0, "+", "+ requires left operand")
	basicmath.RegisterAll(h.Tools)
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "+")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 1 {
		t.Fatalf("expected one token, not %d", len(toks))
	}
	_, ok := p.Parse(nil, toks)
	if ok {
		t.Fatalf("parse erroneously reported successful")
	}
}

func TestPlusIsInvalidWithLHS(t *testing.T) {
	p, h := makeParser(t)
	h.Sink.Expect(1, 1, 0, "5+", "+ requires right operand")
	basicmath.RegisterAll(h.Tools)
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "5+")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 2 {
		t.Fatalf("expected two tokens, not %d", len(toks))
	}
	_, ok := p.Parse(nil, toks)
	if ok {
		t.Fatalf("parse erroneously reported successful")
	}
}

func TestPlusIsInvalidWithRHS(t *testing.T) {
	p, h := makeParser(t)
	h.Sink.Expect(1, 1, 0, "+7", "+ requires left operand")
	basicmath.RegisterAll(h.Tools)
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "+7")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 2 {
		t.Fatalf("expected two tokens, not %d", len(toks))
	}
	_, ok := p.Parse(nil, toks)
	if ok {
		t.Fatalf("parse erroneously reported successful")
	}
}

func TestAdditionEval(t *testing.T) {
	p, h := makeParser(t)
	basicmath.RegisterAll(h.Tools)
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "5+7")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 3 {
		t.Fatalf("expected three token, not %d", len(toks))
	}
	e, ok := p.Parse(nil, toks)
	if !ok {
		t.Fatalf("error parsing %s", line.Text)
	}
	if e.ShortDescription() != "add [Number[5.000000],Number[7.000000]]" {
		t.Fatalf("incorrect parsing: %s", e.ShortDescription())
	}
	val := e.Eval(nil)
	num, ok := val.(float64)
	if !ok {
		t.Fatalf("Eval did not return a float64 but %T", val)
	}
	if num != 12.0 {
		t.Fatalf("5*7 was not 12 but %f", num)
	}
}

func TestMulAddEval(t *testing.T) {
	p, h := makeParser(t)
	basicmath.RegisterAll(h.Tools)
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "3*5+7")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 5 {
		t.Fatalf("expected three token, not %d", len(toks))
	}
	e, ok := p.Parse(nil, toks)
	if !ok {
		t.Fatalf("error parsing %s", line.Text)
	}
	if e.ShortDescription() != "add [mult [Number[3.000000],Number[5.000000]],Number[7.000000]]" {
		t.Fatalf("incorrect parsing: %s", e.ShortDescription())
	}
	val := e.Eval(nil)
	num, ok := val.(float64)
	if !ok {
		t.Fatalf("Eval did not return a float64 but %T", val)
	}
	if num != 22.0 {
		t.Fatalf("3*5+7 was not 22 but %f", num)
	}
}

func TestAddMulEval(t *testing.T) {
	p, h := makeParser(t)
	basicmath.RegisterAll(h.Tools)
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "3+5*7")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 5 {
		t.Fatalf("expected five token, not %d", len(toks))
	}
	e, ok := p.Parse(nil, toks)
	if !ok {
		t.Fatalf("error parsing %s", line.Text)
	}
	if e.ShortDescription() != "add [Number[3.000000],mult [Number[5.000000],Number[7.000000]]]" {
		t.Fatalf("incorrect parsing: %s", e.ShortDescription())
	}
	val := e.Eval(nil)
	num, ok := val.(float64)
	if !ok {
		t.Fatalf("Eval did not return a float64 but %T", val)
	}
	if num != 38.0 {
		t.Fatalf("3+5*7 was not 38 but %f", num)
	}
}

func TestSubAddEval(t *testing.T) {
	p, h := makeParser(t)
	basicmath.RegisterAll(h.Tools)
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "3-5+7")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 5 {
		t.Fatalf("expected five token, not %d", len(toks))
	}
	e, ok := p.Parse(nil, toks)
	if !ok {
		t.Fatalf("error parsing %s", line.Text)
	}
	if e.ShortDescription() != "add [Number[3.000000],mult [Number[5.000000],Number[7.000000]]]" {
		t.Fatalf("incorrect parsing: %s", e.ShortDescription())
	}
	val := e.Eval(nil)
	num, ok := val.(float64)
	if !ok {
		t.Fatalf("Eval did not return a float64 but %T", val)
	}
	if num != 5.0 {
		t.Fatalf("3-5+7 was not 5 but %f", num)
	}
}
