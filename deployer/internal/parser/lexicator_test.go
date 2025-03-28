package parser_test

import (
	"log"
	"testing"

	"ziniki.org/deployer/deployer/internal/parser"
	"ziniki.org/deployer/deployer/pkg/errors"
)

func TestASingleIdIsJustThat(t *testing.T) {
	reporter := mockReporter()
	lex := parser.NewLineLexicator(reporter, "test")
	toks := lex.BlockedLine(1, 1, "hello")
	if len(toks) != 1 {
		t.Fatalf("not exactly one arg returned")
	}
}

func TestLeadingSpacesCauseAPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("recovered")
		}
	}()
	reporter := mockReporter()
	lex := parser.NewLineLexicator(reporter, "test")
	lex.BlockedLine(1, 1, " hello")
	t.Fatalf("did not panic")
}

func mockReporter() errors.ErrorRepI {
	return &errors.ErrorReporter{}
}
