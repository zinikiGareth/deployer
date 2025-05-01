package target

import (
	"fmt"

	"ziniki.org/deployer/coremod/basic"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type commandScope struct {
	repo     pluggable.Repository
	commands *[]action
	storeAs  pluggable.Identifier
}

func (cc *commandScope) Add(entry pluggable.Definition) {
	a, ok := entry.(action)
	if !ok {
		panic(fmt.Sprintf("entry %v is not an action", entry))
	}
	*cc.commands = append(*cc.commands, a)
	if cc.storeAs != nil {
		cc.repo.IntroduceSymbol(pluggable.SymbolName(cc.storeAs.Id()), entry)
	}
}

func (b *commandScope) HaveTokens(tools *pluggable.Tools, tokens []pluggable.Token) pluggable.Interpreter {
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
	var action pluggable.Action
	if cmd.Id() == "ensure" {
		action = &basic.EnsureCommandHandler{}
	} else {
		panic("invalid target command")
	}
	b.storeAs = assignTo
	// TODO: refactor context handler so that it can also store in the repo
	return action.Handle(tools, b, tokens)
}

func (b *commandScope) Completed(tools *pluggable.Tools) {
}

func TargetCommandInterpreter(repo pluggable.Repository, commands *[]action) pluggable.Interpreter {
	return &commandScope{repo: repo, commands: commands}
}
