package repo

import (
	"log"
	"reflect"

	"ziniki.org/deployer/deployer/internal/parser/exprs"
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

func (repo *SimpleRepository) ResolveAll(tools *pluggable.Tools) {
	for _, what := range repo.tops {
		searcher := &Searcher{repo: repo, recall: tools.Recall, reporter: tools.Reporter}
		what.Resolve(searcher)
	}
}

func (d *SimpleRepository) GetDefinition(name pluggable.SymbolName) pluggable.Describable {
	return d.symbols[name]
}

type Searcher struct {
	repo     *SimpleRepository
	recall   pluggable.Recall
	reporter errors.ErrorRepI
}

func (s *Searcher) MakeNew(name pluggable.Identifier) pluggable.Var {
	return exprs.SolidVar(name)
}

func (s *Searcher) Resolve(name pluggable.Identifier) pluggable.Describable {
	defn := s.repo.GetDefinition(pluggable.SymbolName(name.Id()))
	ret, ok := defn.(pluggable.Describable)
	if ret != nil && ok {
		return ret
	}
	ret, ok = s.recall.Find(reflect.TypeFor[pluggable.Blank](), name.Id()).(pluggable.Describable)
	if ret != nil {
		return ret
	}
	log.Printf("failed to resolve %s\n", name)
	s.reporter.At(name.Loc().Line)
	s.reporter.Reportf(name.Loc().Offset, "could not resolve symbol %s", name.Id())
	return nil
}
