package parser

import (
	"unicode"
)

type Blocker struct {
	indents  []string
	handlers []ProvideLine
}

// deffo need an error handler as well
func (b *Blocker) HaveLine(n int, txt string) {
	ind, line := Split(txt)
	if ind == "" {
		return
	}
	if len(b.indents) == 0 {
		b.indents = append(b.indents, ind)
	} else {
		last := b.indents[len(b.indents)-1]
		if last != ind {
			panic("need to actually do blocking")
		}
	}
	b.handlers[len(b.handlers)-1].HaveLine(n, line)
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

func NewBlocker(topLevel ProvideLine) *Blocker {
	ret := &Blocker{}
	ret.handlers = []ProvideLine{topLevel}
	return ret
}
