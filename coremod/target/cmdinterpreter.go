package target

import (
	"ziniki.org/deployer/coremod/basic"
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type commandScope struct {
	repo     pluggable.Repository
	commands *[]pluggable.Definition
	storeAs  pluggable.Identifier
}

func (cc *commandScope) Add(entry pluggable.Definition) {
	*cc.commands = append(*cc.commands, entry)
	if cc.storeAs != nil {
		cc.repo.IntroduceSymbol(pluggable.SymbolName(cc.storeAs.Id()), entry)
	}
}

func (b *commandScope) HaveTokens(reporter errors.ErrorRepI, tokens []pluggable.Token) pluggable.Interpreter {
	// I am hacking this in first, and then I need to come back and do more on it

	if len(tokens) < 1 {
		panic("need a command")
	}
	if tokens[0].(pluggable.Identifier).Id() != "ensure" {
		panic("token[0] is wrong")
	}

	var assignTo pluggable.Identifier
	if len(tokens) >= 3 && tokens[len(tokens)-2].(pluggable.Operator).Is("=>") {
		assignTo = tokens[len(tokens)-1].(pluggable.Identifier)
	}
	cmd, ok := tokens[0].(pluggable.Identifier)
	if !ok {
		panic("command token must be an identifier")
	}
	var action pluggable.Action
	if cmd.Id() == "ensure" {
		action = &basic.EnsureCommandHandler{}
	} else {
		panic("invalid target command")
	}
	b.storeAs = assignTo
	// TODO: refactor context handler so that it can also store in the repo
	return action.Handle(reporter, b.repo, b, tokens)
}

func TargetCommandInterpreter(repo pluggable.Repository, commands *[]pluggable.Definition) pluggable.Interpreter {
	return &commandScope{repo: repo, commands: commands}
}
