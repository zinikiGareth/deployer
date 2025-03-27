package parser

import (
	"unicode"
)

type Blocker struct {
	indents  []string
	handlers []ProvideBlockedLine
}

// deffo need an error handler as well
func (b *Blocker) HaveLine(lineNo int, txt string) {
	ind, line := Split(txt)
	if ind == "" {
		return
	}
	level := b.matchIndent(ind)
	if level >= len(b.indents) {
		b.indents = append(b.indents, ind)
	} else {
		// TODO: clean up old handlers if any (but not indents)
	}
	hdlr := b.handlers[level].BlockedLine(lineNo, len(ind), line)
	if hdlr == nil {
		panic("handler cannot return nil; if no nested scope, return NoInnerScope")
	}
	b.handlers = append(b.handlers, hdlr)
}

func (b *Blocker) matchIndent(ind string) int {
	for idx, curr := range b.indents {
		if ind == curr {
			return idx
		} else if len(curr) >= len(ind) {
			panic("invalid indent should be an error")
		}
	}
	return len(b.indents)
}

func Split(txt string) (string, string) {
	runes := []rune(txt)
	ind := ""
	for len(runes) > 0 && unicode.IsSpace(runes[0]) {
		ind = ind + string(mapSpace(runes[0]))
		runes = runes[1:]
	}
	return ind, string(runes)
}

func mapSpace(ch rune) rune {
	if ch == '\t' {
		return 'T'
	} else if ch == ' ' {
		return 'S'
	} else {
		return 'U' // unicode space character of some form (including invisible)
	}
}

func NewBlocker(topLevel ProvideBlockedLine) *Blocker {
	return &Blocker{handlers: []ProvideBlockedLine{topLevel}}
}
