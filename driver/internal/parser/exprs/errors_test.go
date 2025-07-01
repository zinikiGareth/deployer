package exprs_test

import (
	"testing"

	"ziniki.org/deployer/driver/internal/parser/lexicator"
	"ziniki.org/deployer/driver/pkg/driverbottom"
)

func TestMustHaveAtLeastOneToken(t *testing.T) {
	p, h := makeParser(t)
	lineloc.Text = ""
	h.Sink.Expect(1, 1, 0, "", "no expression found")
	_, ok := p.Parse(nil, []driverbottom.Token{})
	if ok {
		t.Fatalf("No error reported")
	}
}

func TestCannotHaveTwoNouns(t *testing.T) {
	p, h := makeParser(t)
	lineloc.Text = "hello world"
	h.Sink.Expect(1, 1, 0, "hello world", "no function symbol found in this expression")
	hello := lexicator.NewIdentifierToken(lineloc, 0, "hello")
	world := lexicator.NewStringToken(lineloc, 6, "world")
	_, ok := p.Parse(nil, []driverbottom.Token{hello, world})
	if ok {
		t.Fatalf("No error reported")
	}
}
