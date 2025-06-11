package lexicator_test

import (
	"testing"

	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/pluggable"
	"ziniki.org/deployer/deployer/pkg/testhelpers"
)

func TestForPuncIdPunc(t *testing.T) {
	reporter, _ := testhelpers.MockReporter(t)
	tools := pluggable.NewTools(reporter, nil, nil, nil, nil, nil)
	lex := lexicator.NewLineLexicator(tools, "test")
	toks := lex.BlockedLine(lineOf("(hello)"))
	if len(toks) != 3 {
		t.Fatalf("%d args returned, not 3", len(toks))
	}
	if !toks[0].(pluggable.Punc).Is('(') {
		t.Fatalf("first was not ORB")
	}
	if !toks[2].(pluggable.Punc).Is(')') {
		t.Fatalf("last was not )")
	}
}
