package repo

import (
	"log"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

func (repo *SimpleRepository) ResolveAll(tools *driverbottom.CoreTools) {
	for _, what := range repo.tops {
		searcher := &Searcher{repo: repo, recall: tools.Recall, reporter: tools.Reporter}
		what.Resolve(searcher)
	}
}

func (d *SimpleRepository) GetDefinition(name driverbottom.SymbolName) driverbottom.Describable {
	return d.symbols[name]
}

type Searcher struct {
	repo     *SimpleRepository
	recall   driverbottom.Recall
	reporter errorsink.ErrorRepI
}

func (s *Searcher) Resolve(name driverbottom.Identifier) any {
	// log.Printf("attempting to resolve %s\n", name)
	defn := s.repo.GetDefinition(driverbottom.SymbolName(name.Id()))
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
