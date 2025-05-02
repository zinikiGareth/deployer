package exprs_test

import (
	"testing"

	"ziniki.org/deployer/deployer/internal/parser/exprs"
	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func TestASimpleFunctionWithOneArg(t *testing.T) {
	p := makeParser()
	recall.funcs["hello"] = konstFunc
	hello := lexicator.NewIdentifierToken(lineloc, 0, "hello")
	world := lexicator.NewStringToken(lineloc, 6, "world")
	expr, ok := p.Parse([]pluggable.Token{hello, world})
	if !ok {
		t.Fatalf("Parse failed")
	}
	a, ok := expr.(exprs.Apply)
	if !ok {
		t.Fatalf("returned expr was not an Apply")
	}
	if a.Func != konstFunc {
		t.Fatalf("returned Apply Func was not konst")
	}
	if len(a.Args) != 1 {
		t.Fatalf("Apply Func had %d args, not 1", len(a.Args))
	}
	if a.Args[0] != world {
		t.Fatalf("arg was not world")
	}
}
