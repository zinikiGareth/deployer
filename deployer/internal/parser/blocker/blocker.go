package blocker

import (
	"unicode"

	"ziniki.org/deployer/deployer/internal/parser/lexicator"
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type Blocker struct {
	tools    *pluggable.Tools
	indents  []string
	lex      lexicator.Lexicator
	handlers []pluggable.Interpreter
	file     *errors.FileLoc
}

func (b *Blocker) BeginFile(file string) {
	b.file = errors.InFile(file)
}

func (b *Blocker) HaveLine(lineNo int, txt string) {
	ind, line := Split(txt)
	if ind == "" {
		return
	}
	level := b.matchIndent(ind)
	if level == -1 {
		// that's an error, already reported
		return
	} else if level >= len(b.indents) {
		b.indents = append(b.indents, ind)
	} else {
		// Close and remove and surplus handlers
		for level < len(b.handlers)-1 {
			b.handlers[len(b.handlers)-1].Completed()
			b.handlers = b.handlers[0 : len(b.handlers)-1]
		}
	}
	ll := b.file.AtLine(lineNo, level, line)
	b.tools.Reporter.At(ll)
	toks := b.lex.BlockedLine(ll)
	if toks == nil {
		return
	}
	hdlr := b.handlers[level].HaveTokens(toks)
	if hdlr == nil {
		panic("handler cannot return nil; if no nested scope, return NoInnerScope")
	}
	b.handlers = append(b.handlers, hdlr)
}

func (b *Blocker) EndFile() {
	for len(b.handlers) > 0 {
		b.handlers[len(b.handlers)-1].Completed()
		b.handlers = b.handlers[0 : len(b.handlers)-1]
	}
}

func (b *Blocker) matchIndent(ind string) int {
	for idx, curr := range b.indents {
		if ind == curr {
			return idx
		} else if len(curr) >= len(ind) {
			b.tools.Reporter.Report(0, "invalid indent")
			return -1
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

func NewBlocker(tools *pluggable.Tools, lex lexicator.Lexicator, topLevel pluggable.Interpreter) *Blocker {
	return &Blocker{tools: tools, lex: lex, handlers: []pluggable.Interpreter{topLevel}}
}
