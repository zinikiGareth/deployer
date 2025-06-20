package exprs_test

import (
	"testing"

	"ziniki.org/deployer/driver/internal/parser/exprs"
	"ziniki.org/deployer/driver/internal/parser/lexicator"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

func TestATokenCanBeWrappedInParens(t *testing.T) {
	p, _ := makeParser(t)
	pr := p.(exprs.ParenReduction)
	world := lexicator.NewStringToken(lineloc, 6, "world")
	blocks, ok := pr.ReduceParens([]driverbottom.Token{orb, world, crb})
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
	_, ok := pr.ReduceParens([]driverbottom.Token{orb, world})
	if ok {
		t.Fatalf("Parsing should have failed")
	}
}

func TestACRBMustHaveBeenOpened(t *testing.T) {
	p, h := makeParser(t)
	h.Sink.Expect(1, 1, 26, "", "unexpected punc char: )")
	pr := p.(exprs.ParenReduction)
	world := lexicator.NewStringToken(lineloc, 6, "world")
	_, ok := pr.ReduceParens([]driverbottom.Token{world, crb})
	if ok {
		t.Fatalf("Parsing should have failed")
	}
}

func TestParensCanBeDoubleNested(t *testing.T) {
	p, _ := makeParser(t)
	pr := p.(exprs.ParenReduction)
	world := lexicator.NewStringToken(lineloc, 6, "world")
	blocks, ok := pr.ReduceParens([]driverbottom.Token{orb, orb, world, crb, crb})
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

func TestAnExpressionMayContainAList(t *testing.T) {
	p, _ := makeParser(t)
	hello := lexicator.NewStringToken(lineloc, 6, "hello")
	world := lexicator.NewStringToken(lineloc, 6, "world")
	pr := p.(exprs.ParenReduction)
	blocks, ok := pr.ReduceParens([]driverbottom.Token{orb, osb, hello, comma, world, csb, crb})
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
	list, ok := br.Tokens[1].(exprs.AsList)
	if !ok {
		t.Fatalf("br[1] was not an AsList")
	}
	if len(list.Tokens) != 5 {
		t.Fatalf("list has %d tokens, not 5", len(list.Tokens))
	}
}
