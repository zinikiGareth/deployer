package repo

import "ziniki.org/deployer/deployer/pkg/deployer"

type Repository interface {
	AddSymbolListener(lsnr deployer.SymbolListener)
}

type SimpleRepository struct {
	symbolLsnrs []deployer.SymbolListener
}

func (d *SimpleRepository) AddSymbolListener(lsnr deployer.SymbolListener) {
	d.symbolLsnrs = append(d.symbolLsnrs, lsnr)
}

func NewRepository() Repository {
	return &SimpleRepository{}
}
