package target

import (
	"ziniki.org/deployer/coremod/basic"
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type commandList struct {
	commands []pluggable.Definition
}

func (cc *commandList) Add(entry pluggable.Definition) {
	cc.commands = append(cc.commands, entry)
}

type commandScope struct {
	repo     pluggable.Repository
	commands *commandList
}

func (b *commandScope) HaveTokens(reporter errors.ErrorRepI, tokens []pluggable.Token) pluggable.Interpreter {
	// I am hacking this in first, and then I need to come back and do more on it

	if len(tokens) != 3 {
		panic("tokens are wrong")
	}
	if tokens[0].(pluggable.Identifier).Id() != "ensure" {
		panic("token[0] is wrong")
	}
	if tokens[1].(pluggable.Identifier).Id() != "test.S3.Bucket" {
		panic("token[1] is wrong")
	}
	if tokens[2].(pluggable.String).Text() != "org.ziniki.launch_bucket" {
		panic("token[2] is wrong")
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
	return action.Handle(reporter, b.repo, b.commands, tokens)
}

func TargetCommandInterpreter(repo pluggable.Repository, commands *commandList) pluggable.Interpreter {
	return &commandScope{repo: repo, commands: commands}
}
