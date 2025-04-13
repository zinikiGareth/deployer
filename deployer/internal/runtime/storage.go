package runtime

import "ziniki.org/deployer/deployer/pkg/pluggable"

type Storage struct {
}

func (s *Storage) Bind(name pluggable.SymbolName, value any) {

}

func NewRuntimeStorage() pluggable.RuntimeStorage {
	return &Storage{}
}
