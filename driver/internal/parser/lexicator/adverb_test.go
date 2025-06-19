package lexicator_test

import (
	"testing"

	"ziniki.org/deployer/driver/internal/parser/lexicator"
	"ziniki.org/deployer/driver/pkg/pluggable"
	"ziniki.org/deployer/driver/pkg/testhelpers"
)

func TestAnAdverbBeforeAnId(t *testing.T) {
	reporter, _ := testhelpers.MockReporter(t)
	tools := pluggable.NewTools(reporter, nil, nil, nil, nil)
	lex := lexicator.NewLineLexicator(tools, "test")
	toks := lex.BlockedLine(lineOf("@teardown preserve"))
	if len(toks) != 2 {
		t.Fatalf("not exactly one arg returned")
	}
	if toks[0].(pluggable.Adverb).Name() != "teardown" {
		t.Fatalf("token was not teardown")
	}
	if toks[1].(pluggable.Identifier).Id() != "preserve" {
		t.Fatalf("token was not preserve")
	}
}
