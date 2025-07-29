package exprs

import (
	"fmt"

	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type MapExpr struct {
	loc   *errorsink.Location
	pairs []driverbottom.MapEntry
}

func (m *MapExpr) Loc() *errorsink.Location {
	return m.loc
}

func (m *MapExpr) ShortDescription() string {
	panic("unimplemented")
}

func (m *MapExpr) DumpTo(to driverbottom.IndentWriter) {
	to.Intro("List")
	to.AttrsWhere(m)
	for k, e := range m.pairs {
		to.NestedAttr(fmt.Sprintf("pair %d", k), e)
	}
	to.EndAttrs()
}

func (m *MapExpr) Resolve(r driverbottom.Resolver) driverbottom.BindingRequirement {
	ret := driverbottom.MAY_BE_BOUND

	for _, e := range m.pairs {
		e.Value().Resolve(r)
	}
	return ret
}

func (m *MapExpr) Eval(s driverbottom.RuntimeStorage) any {
	ret := map[string]any{}
	for _, e := range m.pairs {
		key := e.Key().Id()
		val := s.Eval(e.Value())
		ret[key] = val
	}
	return ret
}

func (m *MapExpr) IsEmpty() bool {
	return len(m.pairs) == 0
}

func (m *MapExpr) Size() int {
	return len(m.pairs)
}

func (m *MapExpr) String() string {
	return fmt.Sprintf("[<%d>]", len(m.pairs))
}

func (m *MapExpr) Members() []driverbottom.MapEntry {
	return m.pairs
}

type MapPair struct {
	loc   *errorsink.Location
	key   driverbottom.Identifier
	value driverbottom.Expr
}

func (mp *MapPair) Loc() *errorsink.Location {
	return mp.loc
}

func (mp *MapPair) ShortDescription() string {
	return fmt.Sprintf("MapPair[%s<-%s]", mp.key, mp.value.ShortDescription())
}

func (mp *MapPair) DumpTo(to driverbottom.IndentWriter) {
	to.Intro("MapPair")
	to.AttrsWhere(mp)
	to.TextAttr("key", mp.key.Id())
	to.NestedAttr("value", mp.value)
	to.EndAttrs()
}

func (mp *MapPair) Key() driverbottom.Identifier {
	return mp.key
}

func (mp *MapPair) Value() driverbottom.Expr {
	return mp.value
}

func NewMapPair(loc *errorsink.Location, key driverbottom.Identifier, value driverbottom.Expr) driverbottom.MapEntry {
	return &MapPair{loc: loc, key: key, value: value}
}

func NewMapExpr(loc *errorsink.Location, pairs []driverbottom.MapEntry) driverbottom.Map {
	return &MapExpr{loc: loc, pairs: pairs}
}

var _ driverbottom.Map = &MapExpr{}
var _ driverbottom.MapEntry = &MapPair{}
