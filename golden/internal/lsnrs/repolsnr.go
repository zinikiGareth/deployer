package lsnrs

import (
	"log"

	"ziniki.org/deployer/deployer/pkg/deployer"
)

type RepoListener struct {
}

func (r *RepoListener) Symbol(where deployer.Location, what deployer.SymbolType, who deployer.SymbolName) {
	log.Printf("Have a symbol at %v %s %s\n", where, what, who)
}

func NewRepoListener(outdir string) *RepoListener {
	return &RepoListener{}
}
