package target

import (
	"ziniki.org/deployer/coremod/internal/basic"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

// TODO: it feels wrong that this is doing two things: handling tokens AND collecting the results
type commandScope struct {
	repo      pluggable.Repository
	container pluggable.ContainingContext
}

func (cs *commandScope) HaveTokens(tools *pluggable.Tools, tokens []pluggable.Token) pluggable.Interpreter {
	// I am hacking this in first, and then I need to come back and do more on it

	if len(tokens) < 1 {
		panic("need a command")
	}
	if tokens[0].(pluggable.Identifier).Id() != "ensure" {
		panic("token[0] is wrong")
	}

	var assignTo pluggable.Identifier
	if len(tokens) > 3 && tokens[len(tokens)-2].(pluggable.Operator).Is("=>") {
		assignTo = tokens[len(tokens)-1].(pluggable.Identifier)
		tokens = tokens[0 : len(tokens)-2]
	}
	cmd, ok := tokens[0].(pluggable.Identifier)
	if !ok {
		panic("command token must be an identifier")
	}
	var action pluggable.TargetCommand
	if cmd.Id() == "ensure" {
		action = &basic.EnsureCommandHandler{}
	} else {
		panic("invalid target command")
	}

	innerScope := action.Handle(tools, cs.container, tokens, assignTo)
	return innerScope
}

func (b *commandScope) Completed(tools *pluggable.Tools) {
}

func TargetCommandInterpreter(repo pluggable.Repository, container pluggable.ContainingContext) pluggable.Interpreter {
	return &commandScope{repo: repo, container: container}
}
