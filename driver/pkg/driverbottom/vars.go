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

type AttacherCreator interface {
	Create(scope Scope) AttachResult
}
