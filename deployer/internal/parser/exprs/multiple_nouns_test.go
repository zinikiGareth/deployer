package exprs_test

import (
	"testing"

	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func TestZeroTokensMakeZeroExprs(t *testing.T) {
	p, _ := makeParser(t)
	lineloc.Text = ""
	exprs, ok := p.ParseMultiple([]pluggable.Token{})
	if !ok {
		t.Fatalf("Parsing failed")
	}
	if len(exprs) != 0 {
		t.Fatalf("%d exprs returned, not zero", len(exprs))
	}
}

func TestTwoNounsComeBackSeparately(t *testing.T) {
	p, _ := makeParser(t)
	lineloc.Text = "hello world"
	hello := lexicator.NewIdentifierToken(lineloc, 0, "hello")
	world := lexicator.NewStringToken(lineloc, 6, "world")
	exprs, ok := p.ParseMultiple([]pluggable.Token{hello, world})
	if !ok {
		t.Fatalf("Parsing failed")
	}
	if len(exprs) != 2 {
		t.Fatalf("%d exprs returned, not zero", len(exprs))
	}
	if exprs[0] != hello {
		t.Fatalf("first expr was not hello")
	}
	if exprs[1] != world {
		t.Fatalf("second expr was not world")
	}
}

