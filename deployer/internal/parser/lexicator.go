package parser

import (
	"unicode"
)

type LineLexicator struct {
	interpreter Interpreter
}

func (ll *LineLexicator) HaveLine(n int, txt string) {
	var toks []Token
	from := 0
	runes := []rune(txt)
	for k, r := range runes {
		if unicode.IsSpace(r) {
			toks = ll.token(toks, n, from, runes[from:k])
			from = k + 1
		}
	}
	if len(runes) > from {
		toks = ll.token(toks, n, from, runes[from:])
	}
	ll.interpreter.HaveTokens(toks)
}

func (ll *LineLexicator) token(toks []Token, line, start int, text []rune) []Token {
	tok := NewIdentifierToken(line, start, string(text))
	return append(toks, tok)
}

func NewLineLexicator(i Interpreter) *LineLexicator {
	return &LineLexicator{interpreter: i}
}
