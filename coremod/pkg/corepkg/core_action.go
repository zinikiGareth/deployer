package corepkg

import (
	"fmt"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type CoreActionStrategy interface {
	ShortDescription() string
	DumpArgs(iw driverbottom.IndentWriter)
	Resolve(resolver driverbottom.Resolver) driverbottom.BindingRequirement
	DetermineInitialState(tools *corebottom.Tools, loc *errorsink.Location, pres corebottom.ValuePresenter)
	DetermineDesiredState(tools *corebottom.Tools, loc *errorsink.Location, pres corebottom.ValuePresenter)
}

type coreAction struct {
	tools *corebottom.Tools
	loc   *errorsink.Location
	label string
	strat CoreActionStrategy
}

func (da *coreAction) Loc() *errorsink.Location {
	return da.loc
}

func (da *coreAction) ShortDescription() string {
	return fmt.Sprintf("%s[%s]", da.label, da.strat.ShortDescription())
}

func (da *coreAction) DumpTo(w driverbottom.IndentWriter) {
	w.Intro(da.label)
	w.AttrsWhere(da)
	da.strat.DumpArgs(w)
	w.EndAttrs()
}

func (da *coreAction) Resolve(resolver driverbottom.Resolver) driverbottom.BindingRequirement {
	return da.strat.Resolve(resolver)
}

func (da *coreAction) DetermineInitialState(pres corebottom.ValuePresenter) {
	da.strat.DetermineInitialState(da.tools, da.loc, pres)
}

func (da *coreAction) DetermineDesiredState(pres corebottom.ValuePresenter) {
	da.strat.DetermineDesiredState(da.tools, da.loc, pres)
}

func NewCoreAction(tools *corebottom.Tools, loc *errorsink.Location, label string, strat CoreActionStrategy) corebottom.ModelBuilder {
	return &coreAction{tools: tools, loc: loc, label: label, strat: strat}
}

var _ corebottom.ModelBuilder = &coreAction{}
