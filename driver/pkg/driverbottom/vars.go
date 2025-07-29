package driverbottom

type Var interface {
	Expr
	Named() Identifier
	Binding() Holder
}

type Holder interface {
	Describable
	VarName() Identifier
}

type ResolvableHolder interface {
	Holder
	Resolve(s RuntimeStorage)
}
type AttacherCreator interface {
	Create(scope Scope) AttachResult
}
