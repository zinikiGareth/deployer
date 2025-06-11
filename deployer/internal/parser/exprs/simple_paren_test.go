package exprs_test

import (
	"testing"

	"ziniki.org/deployer/deployer/internal/parser/exprs"
	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func TestAnExpressionWithParensRemovesThem(t *testing.T) {
	p, _ := makeParser(t)
	recall.things["x"] = konstFunc
	x := lexicator.NewIdentifierToken(lineloc, 0, "x")
	expr, ok := p.Parse([]pluggable.Token{orb, x, crb})
	if !ok {
		t.Fatalf("could not parse (x)")
	}
	e2, ok := expr.(*exprs.Apply)
	if !ok {
		t.Fatalf("was not an apply")
	}
	if e2.Func != konstFunc {
		t.Fatalf("function was not konstFunc")
	}
	if len(e2.Args) != 0 {
		t.Fatalf("args != 0")
	}
}
