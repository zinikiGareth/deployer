package parser_test

import (
	"testing"

	"ziniki.org/deployer/deployer/internal/parser"
)

func TestABlankLineReturnsEmptyIndent(t *testing.T) {
	ind, line := parser.Split("")
	if ind != "" {
		t.Fatalf("indent not empty")
	}
	if line != "" {
		t.Fatalf("line not empty")
	}
}

func TestACommentLineReturnsAnEmptyIndent(t *testing.T) {
	tx := "This is a comment"
	ind, line := parser.Split(tx)
	if ind != "" {
		t.Fatalf("indent not empty")
	}
	if line != tx {
		t.Fatalf("line not correct: %s != %s", line, tx)
	}
}

func TestALineStartingWithATabReturnsT(t *testing.T) {
	tx := "\ttarget hello"
	ind, line := parser.Split(tx)
	if ind != "T" {
		t.Fatalf("indent not T")
	}
	if line != tx[1:] {
		t.Fatalf("line not correct: %s != %s", line, tx[1:])
	}
}

func TestALineStartingWithTwoSpacesReturnsSS(t *testing.T) {
	tx := "  target hello"
	ind, line := parser.Split(tx)
	if ind != "SS" {
		t.Fatalf("indent not SS")
	}
	if line != tx[2:] {
		t.Fatalf("line not correct: %s != %s", line, tx[2:])
	}
}

func TestALineStartingWithATabAndTwoSpacesReturnsTSS(t *testing.T) {
	tx := "\t  target hello"
	ind, line := parser.Split(tx)
	if ind != "TSS" {
		t.Fatalf("indent not TSS")
	}
	if line != tx[3:] {
		t.Fatalf("line not correct: %s != %s", line, tx[3:])
	}
}

func TestALineWithUnicodeSpaceReturnsU(t *testing.T) {
	tx := "\u00a0target hello"
	ind, line := parser.Split(tx)
	if ind != "U" {
		t.Fatalf("indent not U")
	}
	if line != tx[2:] {
		t.Fatalf("line not correct: %s != %s", line, tx[2:])
	}
}
