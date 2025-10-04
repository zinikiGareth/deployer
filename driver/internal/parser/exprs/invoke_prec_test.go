package exprs_test

import (
	"testing"

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
