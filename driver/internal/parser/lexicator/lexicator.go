package lexicator

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type Lexicator interface {
	BlockedLine(line *errorsink.LineLoc) []driverbottom.Token
}

type LineLexicator struct {
	tools *driverbottom.CoreTools
	file  string
}

type lexmode int

const (
	starting lexmode = iota
	inIdentifier
	inNumber
	inString
	inSymbol
	inAdverb
	stringEnding
)

// TODO: still need to handle PUNC chars: ( ) { } [ ] ,
// Don't handle @ # ^ & ? | ~ \ yet (prob symbol but could be punc)
// Also : ; (prob punc)
// Don't do anything with ``

func (ll *LineLexicator) BlockedLine(line *errorsink.LineLoc) []driverbottom.Token {
	txt := line.Text
	ll.tools.Reporter.At(line)
	var toks []driverbottom.Token
	from := 0
	runes := []rune(txt)
	var quoteRune rune
	mode := starting
	var tok []rune
loop:
	for k := 0; k < len(runes); k++ {
		goAgain := true
		for goAgain && k < len(runes) {
			r := runes[k]
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
				} else if unicode.IsDigit(r) {
					from = k
					mode = inNumber
					tok = append(tok, r)
				} else if isSymbol(r) {
					from = k
					mode = inSymbol
					tok = append(tok, r)
				} else if isIdentifierChar(r) {
					from = k
					mode = inIdentifier
					tok = append(tok, r)
				} else if isPuncChar(r) {
					toks = ll.punctok(toks, line, k, r)
				} else if r == '@' {
					from = k + 1
					mode = inAdverb
				} else {
					ll.tools.Reporter.Report(k, fmt.Sprintf("invalid char '%c'", r))
					return nil
				}
			case inIdentifier:
				if unicode.IsSpace(r) || isSymbol(r) || isPuncChar(r) {
					toks = ll.token(toks, line, from, tok)
					tok = []rune{}
					mode = starting
					goAgain = true
				} else if r == '"' || r == '\'' {
					ll.tools.Reporter.Report(k, "space required after identifier before string")
					return nil
				} else if isIdentifierChar(r) {
					tok = append(tok, r)
				} else { // TODO: stop on non-valid identifier char
				}
			case inNumber:
				if r == '"' || r == '\'' {
					ll.tools.Reporter.Report(k, "space required after identifier before string")
					return nil
				} else if isNumberChar(r) {
					tok = append(tok, r)
					if k+1 < len(runes) {
						continue loop
					} else {
						k++ // simulate once more round the loop if we could
					}
				}
				var err error
				toks, err = ll.numtok(toks, line, from, tok)
				if err != nil {
					needgt := 0
					if string(tok[0:2]) == "0x" {
						needgt = 2
					}
					for len(tok) > needgt {
						toks, err = ll.numtok(toks, line, from, tok)
						if err == nil {
							break
						}
						k--
						tok = tok[0 : len(tok)-1]
					}
					if err != nil {
						ll.tools.Reporter.Report(from, fmt.Sprintf("not a valid number: %s", string(tok)))
						return nil
					}
				}
				tok = []rune{}
				mode = starting
				goAgain = true
			case inSymbol:
				if !isSymbol(r) {
					if strings.HasPrefix(string(tok), "//") {
						break loop
					} else {
						toks = ll.symbol(toks, line, from, tok)
						tok = []rune{}
						mode = starting
						goAgain = true
					}
				} else {
					tok = append(tok, r)
				}
			case inAdverb:
				if unicode.IsSpace(r) {
					if k == from {
						ll.tools.Reporter.Report(k, "adverb name required")
						return nil
					}
					toks = ll.adverb(toks, line, from, tok)
					tok = []rune{}
					mode = starting
					goAgain = true
				} else if unicode.IsLetter(r) {
					tok = append(tok, r)
				} else { // TODO: stop on non-valid identifier char
					ll.tools.Reporter.Report(k, "space required after adverb")
					return nil
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
				} else if isIdentifierChar(r) {
					ll.tools.Reporter.Report(k, "space required after string before identifier")
					return nil
				} else if isNumberChar(r) {
					ll.tools.Reporter.Report(k, "space required after string before number")
					return nil
				} else {
					toks = ll.strtok(toks, line, from, tok)
					tok = []rune{}
					mode = starting
					goAgain = true
				}
			}
		}
	}
	if len(tok) != 0 {
		switch mode {
		case inIdentifier:
			toks = ll.token(toks, line, from, tok)
		case stringEnding:
			toks = ll.strtok(toks, line, from, tok)
		case inSymbol:
			if !strings.HasPrefix(string(tok), "//") {
				toks = ll.symbol(toks, line, from, tok)
			}
		case inNumber:
			// Because of handling the error condition above,
			// this can only happen if we were in the "start"
			// mode and saw a single digit
			if len(tok) == 1 {
				var e error
				toks, e = ll.numtok(toks, line, from, tok)
				if e != nil {
					panic(e)
				}
			} else {
				panic("should have been handled above")
			}
		case inAdverb:
			toks = ll.adverb(toks, line, from, tok)
		case inString:
			ll.tools.Reporter.Report(from, "unterminated string")
			return nil
		default:
			panic("should not have leftover tok:" + string(tok))
		}
	} else {
		switch mode {
		case inAdverb:
			ll.tools.Reporter.Report(from-1, "adverb name required")
			return nil
		default:
			// no worries
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

func isNumberChar(r rune) bool {
	if unicode.IsDigit(r) {
		return true
	}
	if unicode.IsLetter(r) {
		return true
	}
	switch r {
	case 'e', '+', '-', '.': // floating point things
		return true
	case 'x': // radix things
		return true
	}
	return false
}

func isSymbol(r rune) bool {
	switch r {
	case '/', '*', '+', '-':
		return true
	case '!', '$', '%', '~':
		return true
	case '<', '=', '>':
		return true
	default:
		return false
	}
}

func isPuncChar(r rune) bool {
	switch r {
	case '(', ')', '[', ']', '{', '}':
		return true
	case ',', ';', ':':
		return true
	default:
		return false
	}
}

func (ll *LineLexicator) token(toks []driverbottom.Token, line *errorsink.LineLoc, start int, text []rune) []driverbottom.Token {
	tok := NewIdentifierToken(line, start, string(text))
	return append(toks, tok)
}

func (ll *LineLexicator) symbol(toks []driverbottom.Token, line *errorsink.LineLoc, start int, text []rune) []driverbottom.Token {
	tok := NewOperatorToken(line, start, string(text))
	return append(toks, tok)
}

func (ll *LineLexicator) strtok(toks []driverbottom.Token, line *errorsink.LineLoc, start int, text []rune) []driverbottom.Token {
	tok := NewStringToken(line, start, string(text))
	return append(toks, tok)
}

func (ll *LineLexicator) punctok(toks []driverbottom.Token, line *errorsink.LineLoc, start int, p rune) []driverbottom.Token {
	tok := NewPuncToken(line, start, p)
	return append(toks, tok)
}

func (ll *LineLexicator) numtok(toks []driverbottom.Token, line *errorsink.LineLoc, start int, text []rune) ([]driverbottom.Token, error) {
	tx := string(text)
	var f64 float64
	var err error
	if len(tx) > 2 && strings.HasPrefix(tx, "0x") {
		var i64 int64
		i64, err = strconv.ParseInt(tx[2:], 16, 64)
		f64 = float64(i64)
	} else {
		f64, err = strconv.ParseFloat(tx, 64)
	}
	if err != nil {
		return toks, err
	}
	tok := NewNumberToken(line, start, f64)
	return append(toks, tok), nil
}

func (ll *LineLexicator) adverb(toks []driverbottom.Token, line *errorsink.LineLoc, start int, text []rune) []driverbottom.Token {
	tok := NewAdverbToken(line, start, string(text))
	return append(toks, tok)
}

func NewLineLexicator(tools *driverbottom.CoreTools, file string) Lexicator {
	return &LineLexicator{tools: tools, file: file}
}
