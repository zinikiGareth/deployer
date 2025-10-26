package corestrats

import (
	"reflect"

	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/driverbottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type CompositeStrategy interface {
}

type CoreActionStrategy interface {
	ShortDescription() string
	DumpArgs(iw driverbottom.IndentWriter)
	Resolve(resolver driverbottom.Resolver) driverbottom.BindingRequirement
}

type StatefulActionStrategy interface {
	DetermineInitialState(tools *corebottom.Tools, loc *errorsink.Location, pres corebottom.ValuePresenter)
	DetermineDesiredState(tools *corebottom.Tools, loc *errorsink.Location, pres corebottom.ValuePresenter)
}

type RealityUpdaterStrategy interface {
	CoreActionStrategy
	UpdateReality(tools *corebottom.Tools)
}

type CreationStrategy interface {
	DetermineDesiredState(creator CommonCreator, pres corebottom.ValuePresenter)
	UpdateReality(creator CommonCreator, initial any, desired any)
	TearDown(creator CommonCreator, initial any, teardown corebottom.TearDown)
}

type CommonCreator interface {
	GetEnv(driver string, ofType reflect.Type, meth string, field string)
	Name() string
	Loc() *errorsink.Location
	Adopt(item any)
	Created(item any)
	DeferredMethod(name string) driverbottom.Method
}
