package parser

import (
	"strings"
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
	inSymbol
	stringEnding
)

// TODO: still need to handle PUNC chars: ( ) { } [ ] ,
// Don't handle @ # & ? | yet (prob symbol but could be punc)

func (ll *LineLexicator) BlockedLine(lineNo, ind int, txt string) []pluggable.Token {
	ll.reporter.At(lineNo, txt)
	var toks []pluggable.Token
	from := 0
	runes := []rune(txt)
	var quoteRune rune
	mode := starting
	var tok []rune
loop:
	for k, r := range runes {
		goAgain := true
		for goAgain {
			goAgain = false
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
				} else if isSymbol(r) {
					from = k
					mode = inSymbol
					tok = append(tok, r)
				} else { // TODO: numbers, symbols, punc
					from = k
					mode = inIdentifier
					tok = append(tok, r)
				}
			case inIdentifier:
				if unicode.IsSpace(r) || isSymbol(r) {
					toks = ll.token(toks, lineNo, ind+from, tok)
					tok = []rune{}
					mode = starting
					goAgain = true
				} else if r == '"' || r == '\'' {
					ll.reporter.Report(k, "space required after identifier before string")
					return nil
				} else if isIdentifierChar(r) {
					tok = append(tok, r)
				} else { // TODO: stop on non-valid identifier char
				}
			case inSymbol:
				if !isSymbol(r) {
					if strings.HasPrefix(string(tok), "//") {
						break loop
					} else {
						toks = ll.symbol(toks, lineNo, ind+from, tok)
						tok = []rune{}
						mode = starting
						goAgain = true
					}
				} else {
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
				} else if !unicode.IsSpace(r) {
					ll.reporter.Report(k, "space required after string before identifier")
					return nil
				} else {
					toks = ll.strtok(toks, lineNo, ind+from, tok)
					tok = []rune{}
					mode = starting
				}
			}
		}
	}
	if len(tok) != 0 {
		switch mode {
		case inIdentifier:
			toks = ll.token(toks, lineNo, ind+from, tok)
		case stringEnding:
			toks = ll.strtok(toks, lineNo, ind+from, tok)
		case inSymbol:
			if !strings.HasPrefix(string(tok), "//") {
				toks = ll.strtok(toks, lineNo, ind+from, tok)
			}
		case inString:
			ll.reporter.Report(from, "unterminated string")
			return nil
		default:
			panic("should not have leftover tok:" + string(tok))
		}
	}
	return toks
}

func isIdentifierChar(r rune) bool {
	if unicode.IsLetter(r) {
		return true
	}
	if unicode.IsDigit(r) {
		return true
	}
	if r == '_' || r == '.' {
		return true
	}
	return false
}

func isSymbol(r rune) bool {
	if r == '/' || r == '*' || r == '+' || r == '-' {
		return true
	} else if r == '!' || r == '$' || r == '%' {
		return true
	} else if r == '<' || r == '=' || r == '>' {
		return true
	} else {
		return false
	}
}

func (ll *LineLexicator) token(toks []pluggable.Token, line, start int, text []rune) []pluggable.Token {
	tok := NewIdentifierToken(ll.file, line, start, string(text))
	return append(toks, tok)
}

func (ll *LineLexicator) symbol(toks []pluggable.Token, line, start int, text []rune) []pluggable.Token {
	tok := NewOperatorToken(ll.file, line, start, string(text))
	return append(toks, tok)
}

func (ll *LineLexicator) strtok(toks []pluggable.Token, line, start int, text []rune) []pluggable.Token {
	tok := NewStringToken(ll.file, line, start, string(text))
	return append(toks, tok)
}

func NewLineLexicator(reporter errors.ErrorRepI, file string) Lexicator {
	return &LineLexicator{reporter: reporter, file: file}
}
