package exprs_test

import (
	"log"
	"testing"

	"ziniki.org/deployer/driver/internal/basicmath"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

// There are a whole range of issues around fixity and we should try
// and catch them all

// For example, if "~$" is a postfix function (say fac), then both
//     ~$ 3
// and
//   2 ~$ 3
// are invalid just because there cannot be an argument after them.

// Likewise, in the condition of
//   2 ~$ * 4
// precedence does not come into it, because the ONLY valid way to parse this is:
//  (2 ~$)* 4

// On the other hand,
//   4 * 2 ~$
// does depend on precedence, because it could be
//  (4 * 2)~$
// or
//   4 *(2 ~$)

// Likewise, for prefix and postfix operators, the question of
// associativity does not occur, because there is only one way
// to parse
//   3 ~$ ~$
// as
//  (3 ~$)~$

// This file tests a whole bunch of those cases

func TestSimplePostfixExpr(t *testing.T) {
	p, h := makeParser(t)
	basicmath.RegisterAll(h.Tools)
	recall.things["~$"] = postFac
	fl := errorsink.FileLoc{File: "testfile.dply"}
	line := fl.AtLine(1, 1, "3 ~$")
	toks := h.Lex.BlockedLine(line)
	if len(toks) != 2 {
		t.Fatalf("expected two tokens, not %d", len(toks))
	}
	e, ok := p.Parse(nil, toks)
	if !ok {
		t.Fatalf("error parsing %s", line.Text)
	}
	log.Printf("%s", e.ShortDescription())
}
