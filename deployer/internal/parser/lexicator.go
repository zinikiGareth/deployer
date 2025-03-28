package parser

import (
	"unicode"

	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type LineLexicator struct {
	reporter    *errors.ErrorReporter
	interpreter pluggable.Interpreter
	file        string
}

func (ll *LineLexicator) BlockedLine(n, ind int, txt string) pluggable.ProvideBlockedLine {
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
	return ll.interpreter.HaveTokens(ll.reporter, toks)
}

func (ll *LineLexicator) token(toks []pluggable.Token, line, start int, text []rune) []pluggable.Token {
	tok := NewIdentifierToken(ll.file, line, start, string(text))
	return append(toks, tok)
}

func NewLineLexicator(reporter *errors.ErrorReporter, i pluggable.Interpreter, file string) *LineLexicator {
	return &LineLexicator{reporter: reporter, interpreter: i, file: file}
}
