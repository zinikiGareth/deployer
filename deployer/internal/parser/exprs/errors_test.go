package exprs_test

import (
	"testing"

	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func TestMustHaveAtLeastOneToken(t *testing.T) {
	p, h := makeParser(t)
	lineloc.Text = ""
	h.Sink.Expect(1, 1, 0, "", "no expression found")
	_, ok := p.Parse([]pluggable.Token{})
	if ok {
		t.Fatalf("No error reported")
	}
}

func TestCannotHaveTwoNouns(t *testing.T) {
	p, h := makeParser(t)
	lineloc.Text = "hello world"
	h.Sink.Expect(1, 1, 0, "hello world", "no function found")
	hello := lexicator.NewIdentifierToken(lineloc, 0, "hello")
	world := lexicator.NewStringToken(lineloc, 6, "world")
	_, ok := p.Parse([]pluggable.Token{hello, world})
	if ok {
		t.Fatalf("No error reported")
	}
}
