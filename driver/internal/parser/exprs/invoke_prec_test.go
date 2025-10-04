package exprs_test

import (
	"testing"

	"ziniki.org/deployer/driver/internal/basicmath"
	"ziniki.org/deployer/driver/internal/methods"
	"ziniki.org/deployer/driver/internal/parser/lexicator"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

func TestParseMultipleWithNoParensGivesError(t *testing.T) {
	p, h := makeParser(t)
	lineloc.Text = "f x->y"
	h.Sink.Expect(1, 1, 26, lineloc.Text, "too many exprs to left of ->")

	// recall.things["x"] = konstFunc
	recall.things["->"] = methods.MakeInvokeFunc(h.Tools)

	f := lexicator.NewIdentifierToken(lineloc, 0, "f")
	x := lexicator.NewIdentifierToken(lineloc, 0, "x")
	y := lexicator.NewIdentifierToken(lineloc, 0, "y")
	_, ok := p.Parse(nil, []driverbottom.Token{f, x, arrow, y})
	if ok {
		t.Fatalf("parsing incorrectly succeeded")
	}
}

func TestParseInsideFnWithNoParensGivesError(t *testing.T) {
	p, h := makeParser(t)
	lineloc.Text = "f + ->y"
	h.Sink.Expect(1, 1, 26, lineloc.Text, "no expr to left of ->")

	recall.things["+"] = basicmath.MakeAddFunc(h.Tools)
	recall.things["->"] = methods.MakeInvokeFunc(h.Tools)

	f := lexicator.NewIdentifierToken(lineloc, 0, "f")
	plus := lexicator.NewOperatorToken(lineloc, 0, "+")
	y := lexicator.NewIdentifierToken(lineloc, 0, "y")
	_, ok := p.Parse(nil, []driverbottom.Token{f, plus, arrow, y})
	if ok {
		t.Fatalf("parsing incorrectly succeeded")
	}
}

func TestNoObjectOnArrowGivesError(t *testing.T) {
	p, h := makeParser(t)
	lineloc.Text = "->y"
	h.Sink.Expect(1, 1, 26, lineloc.Text, "no expr to left of ->")

	recall.things["->"] = methods.MakeInvokeFunc(h.Tools)

	y := lexicator.NewIdentifierToken(lineloc, 0, "y")
	_, ok := p.Parse(nil, []driverbottom.Token{arrow, y})
	if ok {
		t.Fatalf("parsing incorrectly succeeded")
	}
}

func TestNoMethodForArrowGivesError(t *testing.T) {
	p, h := makeParser(t)
	lineloc.Text = "x->"
	h.Sink.Expect(1, 1, 26, lineloc.Text, "no method after ->")

	recall.things["->"] = methods.MakeInvokeFunc(h.Tools)

	x := lexicator.NewIdentifierToken(lineloc, 0, "x")
	_, ok := p.Parse(nil, []driverbottom.Token{x, arrow})
	if ok {
		t.Fatalf("parsing incorrectly succeeded")
	}
}
