package exprs_test

import (
	"testing"

	"ziniki.org/deployer/deployer/internal/parser/exprs"
	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func TestAStringIsAnExpr(t *testing.T) {
	p, _ := makeParser(t)
	hello := lexicator.NewStringToken(lineloc, 0, "hello")
	expr, ok := p.Parse([]pluggable.Token{hello})
	if !ok {
		t.Fatalf("Parse failed")
	}
	if expr != hello {
		t.Fatalf("returned expr was not hello")
	}
}

func TestANumberIsAnExpr(t *testing.T) {
	p, _ := makeParser(t)
	nbr := lexicator.NewNumberToken(lineloc, 0, 46)
	expr, ok := p.Parse([]pluggable.Token{nbr})
	if !ok {
		t.Fatalf("Parse failed")
	}
	if expr != nbr {
		t.Fatalf("returned expr was not 46")
	}
}

func TestAnUnboundIDIsAnExpr(t *testing.T) {
	p, _ := makeParser(t)
	id := lexicator.NewIdentifierToken(lineloc, 0, "x")
	expr, ok := p.Parse([]pluggable.Token{id})
	if !ok {
		t.Fatalf("Parse failed")
	}
	if !exprs.IsVar(expr, id) {
		t.Fatalf("returned expr was not x")
	}
}

func TestAnIDBoundToAVerbProducesAnExpr(t *testing.T) {
	p, _ := makeParser(t)
	
	recall.things["hello"] = idFunc
	id := lexicator.NewIdentifierToken(lineloc, 0, "hello")
	expr, ok := p.Parse([]pluggable.Token{id})
	if !ok {
		t.Fatalf("Parse failed")
	}
	if exprs.IsVar(expr, id) {
		t.Fatalf("returned expr was the verb")
	}
	if expr != oneString {
		t.Fatalf("returned expr was not the string")
	}
}
