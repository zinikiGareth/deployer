package exprs_test

import (
	"testing"

	"ziniki.org/deployer/deployer/internal/parser/exprs"
	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

var orb = lexicator.NewPuncToken(lineloc, 0, '(')
var crb = lexicator.NewPuncToken(lineloc, 12, ')')

func TestATokenCanBeWrappedInParens(t *testing.T) {
	p, _ := makeParser(t)
	pr := p.(exprs.ParenReduction)
	lineloc.Text = ""
	world := lexicator.NewStringToken(lineloc, 6, "world")
	blocks, ok := pr.ReduceParens([]pluggable.Token{orb, world, crb})
	if !ok {
		t.Fatalf("Parsing failed")
	}
	if len(blocks) != 1 {
		t.Fatalf("%d blocks returned, not 1", len(blocks))
	}
	br, ok := blocks[0].(exprs.Bracketed)
	if !ok {
		t.Fatalf("block[0] was not a bracketed")
	}
	if len(br.Tokens) != 3 {
		t.Fatalf("block[0] has %d tokens, not 3", len(br.Tokens))
	}
}
