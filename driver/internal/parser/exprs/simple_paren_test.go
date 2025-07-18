package exprs_test

import (
	"testing"

	"ziniki.org/deployer/driver/internal/parser/lexicator"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

func TestParseRemovesParensFromAnExpression(t *testing.T) {
	p, _ := makeParser(t)
	recall.things["x"] = konstFunc
	x := lexicator.NewIdentifierToken(lineloc, 0, "x")
	expr, ok := p.Parse(nil, []driverbottom.Token{orb, x, crb})
	if !ok {
		t.Fatalf("could not parse (x)")
	}
	e2, ok := expr.(*konstExpr)
	if !ok {
		t.Fatalf("was not an apply")
	}
	if len(e2.args) != 0 {
		t.Fatalf("args != 0")
	}
}

func TestParseRemovesDoubleParensFromAnExpression(t *testing.T) {
	p, _ := makeParser(t)
	recall.things["x"] = konstFunc
	x := lexicator.NewIdentifierToken(lineloc, 0, "x")
	expr, ok := p.Parse(nil, []driverbottom.Token{orb, orb, x, crb, crb})
	if !ok {
		t.Fatalf("could not parse (x)")
	}
	e2, ok := expr.(*konstExpr)
	if !ok {
		t.Fatalf("was not an apply")
	}
	if len(e2.args) != 0 {
		t.Fatalf("args != 0")
	}
}

func TestParseMultipleRemovesParensFromAnExpression(t *testing.T) {
	p, _ := makeParser(t)
	recall.things["x"] = konstFunc
	x := lexicator.NewIdentifierToken(lineloc, 0, "x")
	es, ok := p.ParseMultiple(nil, []driverbottom.Token{orb, x, crb})
	if !ok {
		t.Fatalf("could not parse (x)")
	}
	if len(es) != 1 {
		t.Fatalf("expected 1 expr not %d", len(es))
	}
	e2, ok := es[0].(*konstExpr)
	if !ok {
		t.Fatalf("was not an apply")
	}
	if len(e2.args) != 0 {
		t.Fatalf("args != 0")
	}
}
