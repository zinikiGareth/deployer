package exprs

import (
	"fmt"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type MapExpr struct {
	pairs []driverbottom.MapEntry
}

func (l *MapExpr) Loc() *errorsink.Location {
	panic("unimplemented")
}

func (l *MapExpr) ShortDescription() string {
	panic("unimplemented")
}

func (l *MapExpr) DumpTo(to driverbottom.IndentWriter) {
	panic("unimplemented")
}

func (l *MapExpr) Resolve(r driverbottom.Resolver) {
	for _, e := range l.pairs {
		e.Value().Resolve(r)
	}
}

func (l *MapExpr) Eval(s driverbottom.RuntimeStorage) any {
	ret := map[string]any{}
	for _, e := range l.pairs {
		key := e.Key().Id()
		val := s.Eval(e.Value())
		ret[key] = val
	}
	return ret
}

func (l *MapExpr) IsEmpty() bool {
	return len(l.pairs) == 0
}

func (l *MapExpr) Size() int {
	return len(l.pairs)
}

func (l *MapExpr) String() string {
	return fmt.Sprintf("[<%d>]", len(l.pairs))
}

func (l *MapExpr) Members() []driverbottom.MapEntry {
	return l.pairs
}

type MapPair struct {
	key   driverbottom.Identifier
	value driverbottom.Expr
}

func (m *MapPair) Key() driverbottom.Identifier {
	return m.key
}

func (m *MapPair) Value() driverbottom.Expr {
	return m.value
}

func NewMapPair(key driverbottom.Identifier, value driverbottom.Expr) driverbottom.MapEntry {
	return &MapPair{key: key, value: value}
}

func NewMapExpr(pairs []driverbottom.MapEntry) driverbottom.Map {
	return &MapExpr{pairs: pairs}
}

var _ driverbottom.Map = &MapExpr{}
var _ driverbottom.MapEntry = &MapPair{}
