package exprs_test

import (
	"testing"

	"ziniki.org/deployer/deployer/internal/parser/exprs"
	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func TestATokenCanBeWrappedInParens(t *testing.T) {
	p, _ := makeParser(t)
	pr := p.(exprs.ParenReduction)
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

func TestAnORBMustBeClosed(t *testing.T) {
	p, h := makeParser(t)
	h.Sink.Expect(1, 1, 26, "", "did not find matching )")
	pr := p.(exprs.ParenReduction)
	world := lexicator.NewStringToken(lineloc, 6, "world")
	_, ok := pr.ReduceParens([]pluggable.Token{orb, world})
	if ok {
		t.Fatalf("Parsing should have failed")
	}
}

func TestACRBMustHaveBeenOpened(t *testing.T) {
	p, h := makeParser(t)
	h.Sink.Expect(1, 1, 26, "", "unexpected close paren: )")
	pr := p.(exprs.ParenReduction)
	world := lexicator.NewStringToken(lineloc, 6, "world")
	_, ok := pr.ReduceParens([]pluggable.Token{world, crb})
	if ok {
		t.Fatalf("Parsing should have failed")
	}
}
