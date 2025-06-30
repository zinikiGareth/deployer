package repo

import (
	"fmt"

	"ziniki.org/deployer/driver/pkg/driverbottom"
)

type SimpleRepository struct {
	symbolLsnrs []driverbottom.SymbolListener
	symbols     map[driverbottom.SymbolName]driverbottom.Describable
	tops        []driverbottom.TopLevelForm
}

func (d *SimpleRepository) ReadingFile(file string) {
	for _, lsnr := range d.symbolLsnrs {
		lsnr.ReadingFile(file)
	}
}

func (d *SimpleRepository) IntroduceSymbol(who driverbottom.SymbolName, is driverbottom.Describable) error {
	if d.symbols[who] != nil {
		return fmt.Errorf("duplicate definition of %s", who)
	}
	d.symbols[who] = is
	for _, lsnr := range d.symbolLsnrs {
		lsnr.Symbol(who, is)
	}
	return nil
}

func (d *SimpleRepository) TopLevel(defn driverbottom.TopLevelForm) {
	d.tops = append(d.tops, defn)
}

func (d *SimpleRepository) AddSymbolListener(lsnr driverbottom.SymbolListener) {
	d.symbolLsnrs = append(d.symbolLsnrs, lsnr)
}

func (d *SimpleRepository) Traverse(lsnr driverbottom.RepositoryTraverser) {
	for who, what := range d.symbols {
		lsnr.Visit(who, what)
	}
}

func (d *SimpleRepository) FindTop(name driverbottom.SymbolName) driverbottom.TopLevelForm {
	for _, top := range d.tops {
		if top.Name() == name {
			return top
		}
	}
	return nil
}

func NewRepository() driverbottom.Repository {
	return &SimpleRepository{symbols: make(map[driverbottom.SymbolName]driverbottom.Describable)}
}
