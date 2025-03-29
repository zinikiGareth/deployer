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
	reporter errors.ErrorRepI
	file     string
}

type lexmode int

const (
	starting lexmode = iota
	inIdentifier
	inString
	stringEnding
)

func (ll *LineLexicator) BlockedLine(n, ind int, txt string) []pluggable.Token {
	var toks []pluggable.Token
	from := 0
	runes := []rune(txt)
	var quoteRune rune
	mode := starting
	var tok []rune
	for k, r := range runes {
		switch mode {
		case starting:
			tok = []rune{}
			if unicode.IsSpace(r) {
				if k == 0 {
					panic("cannot have leading spaces on a line")
				}
			} else if r == '"' || r == '\'' {
				from = k + 1
				mode = inString
				quoteRune = r
			} else { // TODO: numbers, symbols
				from = k
				mode = inIdentifier
				tok = append(tok, r)
			}
		case inIdentifier:
			if unicode.IsSpace(r) {
				toks = ll.token(toks, n, ind+from, tok)
				tok = []rune{}
				mode = starting
			} else { // TODO: stop on non-valid identifier char
				tok = append(tok, r)
			}
		case inString:
			if r == quoteRune {
				mode = stringEnding
			} else {
				tok = append(tok, r)
			}
		case stringEnding:
			if r == quoteRune { // it was "" in the middle of the string, add one of them
				tok = append(tok, quoteRune)
				mode = inString
			} else {
				toks = ll.strtok(toks, n, ind+from, tok)
				tok = []rune{}
				mode = starting
			}
		}
	}
	if len(tok) != 0 {
		switch mode {
		case inIdentifier:
			toks = ll.token(toks, n, ind+from, tok)
		case stringEnding:
			toks = ll.strtok(toks, n, ind+from, tok)
		default:
			panic("should not have leftover tok:" + string(tok))
		}
	}
	return toks
}

func (ll *LineLexicator) token(toks []pluggable.Token, line, start int, text []rune) []pluggable.Token {
	tok := NewIdentifierToken(ll.file, line, start, string(text))
	return append(toks, tok)
}

func (ll *LineLexicator) strtok(toks []pluggable.Token, line, start int, text []rune) []pluggable.Token {
	tok := NewStringToken(ll.file, line, start, string(text))
	return append(toks, tok)
}

func NewLineLexicator(reporter errors.ErrorRepI, file string) Lexicator {
	return &LineLexicator{reporter: reporter, file: file}
}
