package main

import (
	"ziniki.org/deployer/deployer/pkg/deployer"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

var testRunner deployer.TestRunner

func ProvideTestRunner(runner deployer.TestRunner) error {
	testRunner = runner
	return nil
}

func RegisterWithDeployer(deployer deployer.Deployer) error {
	eh := testRunner.ErrorHandlerFor("log")
	eh.WriteMsg("Need to install things from coremod\n")
	register := deployer.ObtainRegister()
	register.RegisterVerb("target", &CoreTarget{})
	return nil
}

type CoreTarget struct {
}

func (t *CoreTarget) Handle(repo pluggable.Repository, tokens []pluggable.Token) {
	t1 := tokens[1].(pluggable.Identifier)
	repo.IntroduceSymbol(t1.Loc(), pluggable.SymbolType("core.Target"), pluggable.SymbolName(t1.Id()))
}
