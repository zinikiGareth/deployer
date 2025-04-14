package runtime

import (
	"ziniki.org/deployer/deployer/pkg/errors"
	"ziniki.org/deployer/deployer/pkg/pluggable"
)

type Storage struct {
	sink errors.ErrorSink
}

func (s *Storage) Bind(name pluggable.SymbolName, value any) {

}

func (s *Storage) Errorf(loc pluggable.Location, msg string, args ...any) {
	s.sink.Reportf(loc.Line, loc.Offset, "", msg, args...)
}

func NewRuntimeStorage(sink errors.ErrorSink) pluggable.RuntimeStorage {
	return &Storage{sink: sink}
}
