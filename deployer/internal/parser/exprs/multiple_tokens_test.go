package exprs_test

import (
	"testing"

	"ziniki.org/deployer/deployer/internal/parser/exprs"
	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func TestASingleStringGivesOneExpr(t *testing.T) {
	p, _ := makeParser(t)
	hello := lexicator.NewStringToken(lineloc, 0, "hello")
	exprs, ok := p.ParseMultiple([]pluggable.Token{hello})
	if !ok {
		t.Fatalf("Parse failed")
	}
	if len(exprs) != 1 {
		t.Fatalf("one expr was not returned")
	}
	if exprs[0] != hello {
		t.Fatalf("returned expr was not hello")
	}
}

func TestASingleNumberGivesOneExpr(t *testing.T) {
	p, _ := makeParser(t)
	nbr := lexicator.NewNumberToken(lineloc, 0, 46)
	exprs, ok := p.ParseMultiple([]pluggable.Token{nbr})
	if !ok {
		t.Fatalf("Parse failed")
	}
	if len(exprs) != 1 {
		t.Fatalf("one expr was not returned")
	}
	if exprs[0] != nbr {
		t.Fatalf("returned expr was not 46")
	}
}

func TestAnUnboundIDIsOneExpr(t *testing.T) {
	p, _ := makeParser(t)
	id := lexicator.NewIdentifierToken(lineloc, 0, "x")
	es, ok := p.ParseMultiple([]pluggable.Token{id})
	if !ok {
		t.Fatalf("Parse failed")
	}
	if len(es) != 1 {
		t.Fatalf("one expr was not returned")
	}
	if !exprs.IsVar(es[0], id) {
		t.Fatalf("returned expr was not x")
	}
}

func TestAnIDBoundToAVerbProducesASingleExpr(t *testing.T) {
	p, _ := makeParser(t)
	recall.things["hello"] = idFunc
	id := lexicator.NewIdentifierToken(lineloc, 0, "hello")
	es, ok := p.ParseMultiple([]pluggable.Token{id})
	if !ok {
		t.Fatalf("Parse failed")
	}
	if len(es) != 1 {
		t.Fatalf("one expr was not returned")
	}
	if exprs.IsVar(es[0], id) {
		t.Fatalf("returned expr was the verb")
	}
	if es[0] != oneString {
		t.Fatalf("returned expr was not the string")
	}
}
