package blocker_test

import "ziniki.org/deployer/driver/pkg/driverbottom"

type RepoDouble struct {
}

func (m *RepoDouble) AddSymbolListener(lsnr driverbottom.SymbolListener) {
}

func (m *RepoDouble) AtLevel(level int) driverbottom.Scope {
	return nil
}

func (m *RepoDouble) CurrentScope() driverbottom.Scope {
	return nil
}

func (m *RepoDouble) FindTop(name driverbottom.SymbolName) driverbottom.TopLevelForm {
	return nil
}

func (m *RepoDouble) GetDefinition(id driverbottom.SymbolName) driverbottom.Describable {
	return nil
}

func (m *RepoDouble) ReadingFile(file string) {
}

func (m *RepoDouble) ResolveAll(tools *driverbottom.CoreTools) {
}

func (m *RepoDouble) TopLevel(tlf driverbottom.TopLevelForm) {
}

func (m *RepoDouble) TopScope() driverbottom.Scope {
	return nil
}

func (m *RepoDouble) Traverse(lsnr driverbottom.RepositoryTraverser) {
}

var _ driverbottom.Repository = &RepoDouble{}
