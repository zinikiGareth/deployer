package lexicator_test

import (
	"testing"

	"ziniki.org/deployer/driver/internal/parser/lexicator"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/testhelpers"
)

func TestAnAdverbBeforeAnId(t *testing.T) {
	reporter, _ := testhelpers.MockReporter(t)
	tools := driverbottom.NewTools(reporter, nil, nil, nil, nil)
	lex := lexicator.NewLineLexicator(tools, "test")
	toks := lex.BlockedLine(lineOf("@adverb fast"))
	if len(toks) != 2 {
		t.Fatalf("not exactly one arg returned")
	}
	if toks[0].(driverbottom.Adverb).Name() != "adverb" {
		t.Fatalf("token was not adverb")
	}
	if toks[1].(driverbottom.Identifier).Id() != "fast" {
		t.Fatalf("token was not fast")
	}
}
