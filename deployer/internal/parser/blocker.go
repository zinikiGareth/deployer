package parser

import (
	"fmt"
	"unicode"
)

type Blocker struct {
}

// deffo need an error handler as well
func (b *Blocker) HaveLine(n int, txt string) {
	ind, line := Split(txt)
	fmt.Printf("%d: %s %s\n", n, ind, line)
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

func NewBlocker() *Blocker {
	return &Blocker{}
}
