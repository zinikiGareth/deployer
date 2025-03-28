package parser

import (
	"unicode"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type Lexicator interface {
	BlockedLine(n, ind int, txt string) []pluggable.Token
}

type LineLexicator struct {
	reporter *errors.ErrorReporter
	file     string
}

func (ll *LineLexicator) BlockedLine(n, ind int, txt string) []pluggable.Token {
	var toks []pluggable.Token
	from := 0
	runes := []rune(txt)
	for k, r := range runes {
		if unicode.IsSpace(r) {
			toks = ll.token(toks, n, ind+from, runes[from:k])
			from = k + 1
		}
	}
	if len(runes) > from {
		toks = ll.token(toks, n, ind+from, runes[from:])
	}
	return toks
}

func (ll *LineLexicator) token(toks []pluggable.Token, line, start int, text []rune) []pluggable.Token {
	tok := NewIdentifierToken(ll.file, line, start, string(text))
	return append(toks, tok)
}

func NewLineLexicator(reporter *errors.ErrorReporter, file string) Lexicator {
	return &LineLexicator{reporter: reporter, file: file}
}
