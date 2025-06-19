package driverbottom

type Method interface {
	Invoke(storage RuntimeStorage, obj Expr, args []Expr) any
}

type HasMethods interface {
	ObtainMethod(name string) Method
}
