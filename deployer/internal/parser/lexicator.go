package parser

import (
	"unicode"
)

type LineLexicator struct {
	interpreter Interpreter
	file        string
}

func (ll *LineLexicator) BlockedLine(n, ind int, txt string) {
	var toks []Token
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
	ll.interpreter.HaveTokens(toks)
}

func (ll *LineLexicator) token(toks []Token, line, start int, text []rune) []Token {
	tok := NewIdentifierToken(ll.file, line, start, string(text))
	return append(toks, tok)
}

func NewLineLexicator(i Interpreter, file string) *LineLexicator {
	return &LineLexicator{interpreter: i, file: file}
}
