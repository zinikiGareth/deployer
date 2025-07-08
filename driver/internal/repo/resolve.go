package repo

import (
	"log"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

func (repo *SimpleRepository) ResolveAll(tools *driverbottom.CoreTools) bool {
	hasError := false
	for _, what := range repo.tops {
		searcher := &Searcher{repo: repo, recall: tools.Recall, reporter: tools.Reporter}
		binding := what.Resolve(searcher)
		if binding == driverbottom.ERROR_OCCURRED {
			hasError = true
		}
	}
	return hasError
}

func (d *SimpleRepository) GetDefinition(name driverbottom.SymbolName) driverbottom.Describable {
	scope := d.symbols[name]
	if scope == nil {
		return nil
	}
	return scope
}

type Searcher struct {
	repo     *SimpleRepository
	recall   driverbottom.Recall
	reporter errorsink.ErrorRepI
}

func (s *Searcher) Resolve(scope driverbottom.Scope, name driverbottom.Identifier) any {
	// log.Printf("attempting to resolve %s\n", name)
	defn := scope.FindDefinition(driverbottom.SymbolName(name.Id()))
	if defn != nil {
		// log.Printf("defn for %s => %T %v\n", name, defn, defn)
		return defn
	}
	ret := s.recall.Find("blank", name.Id())
	if ret != nil {
		return ret
	}
	log.Printf("failed to resolve %s\n", name)
	s.reporter.ReportAtf(name.Loc(), "could not resolve symbol %s", name.Id())

	return nil
}
