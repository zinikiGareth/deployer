package exprs_test

import (
	"testing"

	"ziniki.org/deployer/driver/internal/parser/exprs"
	"ziniki.org/deployer/driver/internal/parser/lexicator"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

func TestAnExpressionWithParensRemovesThem(t *testing.T) {
	p, _ := makeParser(t)
	recall.things["x"] = konstFunc
	x := lexicator.NewIdentifierToken(lineloc, 0, "x")
	expr, ok := p.Parse([]driverbottom.Token{orb, x, crb})
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
