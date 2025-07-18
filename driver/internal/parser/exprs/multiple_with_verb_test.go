package exprs_test

import (
	"testing"

	"ziniki.org/deployer/driver/internal/parser/lexicator"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

func TestAVerbAndANoun(t *testing.T) {
	p, _ := makeParser(t)
	recall.things["hello"] = konstFunc
	hello := lexicator.NewIdentifierToken(lineloc, 0, "hello")
	world := lexicator.NewStringToken(lineloc, 6, "world")
	exs, ok := p.ParseMultiple(nil, []driverbottom.Token{hello, world})
	if !ok {
		t.Fatalf("Parse failed")
	}
	if len(exs) != 2 {
		t.Fatalf("%d args returned, not 2", len(exs))
	}
	a, ok := exs[0].(*konstExpr)
	if !ok {
		t.Fatalf("returned expr was not an Apply")
	}
	if len(a.args) != 0 {
		t.Fatalf("Apply Func had %d args, not 0", len(a.args))
	}

	if exs[1] != world {
		t.Fatalf("second was not world")
	}
}
