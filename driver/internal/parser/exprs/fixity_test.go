package exprs_test

import (
	"testing"

	"ziniki.org/deployer/driver/internal/basicmath"
	"ziniki.org/deployer/driver/internal/lists"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

// There are a whole range of issues around fixity and we should try
// and catch them all

// For example, if "~$" is a postfix function (say fac), then both
//     ~$ 3
// and
//   2 ~$ 3
// are invalid just because there cannot be an argument after them.

// Likewise, in the condition of
//   2 ~$ * 4
// precedence does not come into it, because the ONLY valid way to parse this is:
//  (2 ~$)* 4

// On the other hand,
//   4 * 2 ~$
// does depend on precedence, because it could be
//  (4 * 2)~$
// or
//   4 *(2 ~$)

// Likewise, for prefix and postfix operators, the question of
// associativity does not occur, because there is only one way
// to parse
//   3 ~$ ~$
// as
//  (3 ~$)~$

// This file tests a whole bunch of those cases

func TestSimplePostfixExpr(t *testing.T) {
	p, h := makeParser(t)
	basicmath.RegisterAll(h.Tools)
	recall.things["~$"] = postFac
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "3 ~$")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 2 {
		t.Fatalf("expected two tokens, not %d", len(toks))
	}
	e, ok := p.Parse(nil, toks)
	if !ok {
		t.Fatalf("error parsing %s", line.Text)
	}
	if e.ShortDescription() != "<fac>(Number[3.000000])" {
		t.Fatalf("incorrect parsing: %s", e.ShortDescription())
	}
	k := e.Eval(nil)
	if k != 6 {
		t.Fatalf("was not 6 but %v", k)
	}
}

func TestPostfixExprCannotBePrefix(t *testing.T) {
	p, h := makeParser(t)
	h.Sink.Expect(1, 1, 0, "~$ 3", "postfix operator cannot have arguments afterwards")
	basicmath.RegisterAll(h.Tools)
	recall.things["~$"] = postFac
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "~$ 3")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 2 {
		t.Fatalf("expected two tokens, not %d", len(toks))
	}
	_, ok := p.Parse(nil, toks)
	if ok {
		t.Fatalf("parse erroneously reported successful")
	}
}

func TestPostfixExprCannotBeInfix(t *testing.T) {
	p, h := makeParser(t)
	h.Sink.Expect(1, 1, 0, "2 ~$ 3", "postfix operator cannot have arguments afterwards")
	basicmath.RegisterAll(h.Tools)
	recall.things["~$"] = postFac
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "2 ~$ 3")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 3 {
		t.Fatalf("expected two tokens, not %d", len(toks))
	}
	_, ok := p.Parse(nil, toks)
	if ok {
		t.Fatalf("parse erroneously reported successful")
	}
}

func TestPostfixOpMustBeDoneFirstIfFirstTerm(t *testing.T) {
	p, h := makeParser(t)
	basicmath.RegisterAll(h.Tools)
	recall.things["~$"] = postFac
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "3 ~$ * 7")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 4 {
		t.Fatalf("expected four tokens, not %d", len(toks))
	}
	e, ok := p.Parse(nil, toks)
	if !ok {
		t.Fatalf("error parsing %s", line.Text)
	}
	if e.ShortDescription() != "mult [<fac>(Number[3.000000]),Number[7.000000]]" {
		t.Fatalf("incorrect parsing: %s", e.ShortDescription())
	}
	k := e.Eval(nil)
	if k != 42.0 {
		t.Fatalf("was not 42 but %v", k)
	}
}

func TestPostfixOpConsidersPrecedence(t *testing.T) {
	p, h := makeParser(t)
	basicmath.RegisterAll(h.Tools)
	recall.things["~$"] = postFac
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "7 * 3 ~$")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 4 {
		t.Fatalf("expected four tokens, not %d", len(toks))
	}
	e, ok := p.Parse(nil, toks)
	if !ok {
		t.Fatalf("error parsing %s", line.Text)
	}
	if e.ShortDescription() != "mult [Number[7.000000],<fac>(Number[3.000000])]" {
		t.Fatalf("incorrect parsing: %s", e.ShortDescription())
	}
	k := e.Eval(nil)
	if k != 42.0 {
		t.Fatalf("was not 42 but %v", k)
	}
}

func TestPostfixOpAssociatesRight(t *testing.T) {
	p, h := makeParser(t)
	basicmath.RegisterAll(h.Tools)
	recall.things["~$"] = postFac
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "3 ~$ ~$")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 3 {
		t.Fatalf("expected three tokens, not %d", len(toks))
	}
	e, ok := p.Parse(nil, toks)
	if !ok {
		t.Fatalf("error parsing %s", line.Text)
	}
	if e.ShortDescription() != "<fac>(<fac>(Number[3.000000]))" {
		t.Fatalf("incorrect parsing: %s", e.ShortDescription())
	}
	k := e.Eval(nil)
	if k != 720 {
		t.Fatalf("was not 42 but %v", k)
	}
}

func TestPostfixOpWithLowPrecedenceGivesABigNumber(t *testing.T) {
	p, h := makeParser(t)
	basicmath.RegisterAll(h.Tools)
	postFac.(*postFacFunc).prec = 2
	recall.things["~$"] = postFac
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "3 * 3 ~$")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 4 {
		t.Fatalf("expected four tokens, not %d", len(toks))
	}
	e, ok := p.Parse(nil, toks)
	if !ok {
		t.Fatalf("error parsing %s", line.Text)
	}
	if e.ShortDescription() != "<fac>(mult [Number[3.000000],Number[3.000000]])" {
		t.Fatalf("incorrect parsing: %s", e.ShortDescription())
	}
	k := e.Eval(nil)
	if k != 362880 { // is an integer because fac() is done second
		t.Fatalf("was not 362880 but %v", k)
	}
}

func TestPostfixExprCannotComeBeforePrefixOp(t *testing.T) {
	p, h := makeParser(t)
	h.Sink.Expect(1, 1, 0, "2 ~$ sum 3", "cannot have postfix operator followed by prefix operator")
	basicmath.RegisterAll(h.Tools)
	recall.things["~$"] = postFac
	recall.things["sum"] = lists.MakeSumFunc(h.Tools)
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "2 ~$ sum 3")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 4 {
		t.Fatalf("expected four tokens, not %d", len(toks))
	}
	_, ok := p.Parse(nil, toks)
	if ok {
		t.Fatalf("parse erroneously reported successful")
	}
}

func TestASuitablyLowPrecedenceFuncCanSuckUpSubsequentPostfixExprThenPrefixOp(t *testing.T) {
	p, h := makeParser(t)
	basicmath.RegisterAll(h.Tools)
	recall.things["collect"] = collectFunc
	recall.things["~$"] = postFac
	recall.things["sum"] = lists.MakeSumFunc(h.Tools)
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "collect 2 ~$ sum 3")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 5 {
		t.Fatalf("expected five tokens, not %d", len(toks))
	}
	_, ok := p.Parse(nil, toks)
	if ok {
		t.Fatalf("parse erroneously reported successful")
	}
}

func TestParensBeatPrePostIssues(t *testing.T) {
	p, h := makeParser(t)
	basicmath.RegisterAll(h.Tools)
	recall.things["collect"] = collectFunc
	recall.things["~$"] = postFac
	recall.things["sum"] = lists.MakeSumFunc(h.Tools)
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "collect (2 ~$) sum 3")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 7 {
		t.Fatalf("expected seven tokens, not %d", len(toks))
	}
	e, ok := p.Parse(nil, toks)
	if !ok {
		t.Fatalf("error parsing %s", line.Text)
	}
	if e.ShortDescription() != "collect(<fac>(Number[2.000000]), sum(Number[3.000000]))" {
		t.Fatalf("incorrect parsing: %s", e.ShortDescription())
	}
	k := e.Eval(nil)
	sle := k.([]interface{})
	if len(sle) != 2 {
		t.Fatalf("length was not 2 but %d", len(sle))
	}
}
