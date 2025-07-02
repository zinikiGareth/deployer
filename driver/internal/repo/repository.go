package repo

import (
	"fmt"
	"log"

	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type RepoScope struct {
	repo    *SimpleRepository
	name    string
	entries map[driverbottom.SymbolName]driverbottom.Describable
	parent  *RepoScope
	inners  []*RepoScope
}

// Name implements driverbottom.Scope.
func (s *RepoScope) Name() string {
	return s.name
}

func (s *RepoScope) IntroduceSymbol(who driverbottom.SymbolName, is driverbottom.Describable) error {
	if s.entries[who] != nil {
		ret := fmt.Errorf("duplicate definition of %s", who)
		log.Printf("%v", ret)
		return ret
	}
	s.entries[who] = is
	for _, lsnr := range s.repo.symbolLsnrs {
		lsnr.Symbol(s, who, is)
	}
	return nil
}

func (s *RepoScope) FindDefinition(name driverbottom.SymbolName) any {
	e := s.entries[name]
	if e != nil {
		return e
	}
	if s.parent == nil {
		return nil
	}
	return s.parent.FindDefinition(name)
}

func (s *RepoScope) Traverse(lsnr driverbottom.RepositoryTraverser) {
	// log.Printf("scope %p has entries %d %v\n", s, len(s.entries), s.entries)
	for k, v := range s.entries {
		lsnr.Visit(k, v)
		inner, ok := v.(driverbottom.HasScope)
		if ok {
			if inner.Scope() == s {
				panic("not actually nested")
			}
			if inner.Scope() != nil {
				inner.Scope().Traverse(lsnr)
			}
		} else {
			log.Fatalf("%p %T is not HasScope\n", v, v)
		}
	}
}

type SimpleRepository struct {
	symbolLsnrs []driverbottom.SymbolListener
	symbols     map[driverbottom.SymbolName]driverbottom.Describable // deprecated - use scope.entries
	tops        []driverbottom.TopLevelForm
	stack       []*RepoScope
}

func (d *SimpleRepository) ReadingFile(file string) {
	// d.stack = []*RepoScope{{entries: make(map[driverbottom.SymbolName]driverbottom.Describable)}}
	for _, lsnr := range d.symbolLsnrs {
		lsnr.ReadingFile(file)
	}
}

func (d *SimpleRepository) AtLevel(level int) driverbottom.Scope {
	level = level + 1
	d.stack = d.stack[0:level]
	for level >= len(d.stack) {
		var p *RepoScope
		if len(d.stack) > 0 {
			p = d.stack[len(d.stack)-1]
		}
		inner := &RepoScope{repo: d, name: "undefined", parent: p, entries: make(map[driverbottom.SymbolName]driverbottom.Describable)}
		if p != nil {
			p.inners = append(p.inners, inner)
		}
		d.stack = append(d.stack, inner)
	}
	return d.stack[len(d.stack)-1]
}

func (d *SimpleRepository) TopLevel(defn driverbottom.TopLevelForm) {
	d.tops = append(d.tops, defn)
}

func (d *SimpleRepository) AddSymbolListener(lsnr driverbottom.SymbolListener) {
	d.symbolLsnrs = append(d.symbolLsnrs, lsnr)
}

func (d *SimpleRepository) Traverse(lsnr driverbottom.RepositoryTraverser) {
	d.TopScope().Traverse(lsnr)
}

func (d *SimpleRepository) FindTop(name driverbottom.SymbolName) driverbottom.TopLevelForm {
	for _, top := range d.tops {
		if top.Name() == name {
			return top
		}
	}
	return nil
}

func (d *SimpleRepository) TopScope() driverbottom.Scope {
	return d.stack[0]
}

func (d *SimpleRepository) CurrentScope() driverbottom.Scope {
	return d.stack[len(d.stack)-2]
}

func NewRepository() driverbottom.Repository {
	ret := &SimpleRepository{symbols: make(map[driverbottom.SymbolName]driverbottom.Describable), stack: []*RepoScope{}}
	topScope := &RepoScope{repo: ret, name: "{}", entries: make(map[driverbottom.SymbolName]driverbottom.Describable)}
	// log.Printf("top = %p\n", topScope)
	ret.stack = append(ret.stack, topScope)
	return ret
}

var _ driverbottom.Scope = &RepoScope{}
