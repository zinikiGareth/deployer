package exprs_test

import (
	"fmt"
	"testing"

	"ziniki.org/deployer/deployer/internal/parser/exprs"
	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func TestAVerbAndANoun(t *testing.T) {
	p, _ := makeParser(t)
	recall.things["hello"] = konstFunc
	hello := lexicator.NewIdentifierToken(lineloc, 0, "hello")
	world := lexicator.NewStringToken(lineloc, 6, "world")
	exs, ok := p.ParseMultiple([]pluggable.Token{hello, world})
	if !ok {
		t.Fatalf("Parse failed")
	}
	if len(exs) != 2 {
		t.Fatalf("%d args returned, not 2", len(exs))
	}
	fmt.Printf("%v\n", exs[0])
	a, ok := exs[0].(exprs.Apply)
	if !ok {
		t.Fatalf("returned expr was not an Apply")
	}
	if a.Func != konstFunc {
		t.Fatalf("returned Apply Func was not konst")
	}
	if len(a.Args) != 0 {
		t.Fatalf("Apply Func had %d args, not 0", len(a.Args))
	}

	if exs[1] != world {
		t.Fatalf("second was not world")
	}
}
