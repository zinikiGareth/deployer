package exprs_test

import (
	"testing"

	"ziniki.org/deployer/driver/internal/basicmath"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

func TestSimpleMultiplication(t *testing.T) {
	p, h := makeParser(t)
	basicmath.RegisterAll(h.Tools)
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "2*3")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 3 {
		t.Fatalf("expected three tokens, not %d", len(toks))
	}
	e, ok := p.Parse(nil, toks)
	if !ok {
		t.Fatalf("error parsing %s", line.Text)
	}
	if e.ShortDescription() != "mult [Number[2.000000],Number[3.000000]]" {
		t.Fatalf("incorrect parsing: %s", e.ShortDescription())
	}
}
