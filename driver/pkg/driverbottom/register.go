package driverbottom

type Register interface {
	ExtensionPoint(name string)
	Register(point string, called string, item any)
	ProvideDriver(s string, env any)
}

type Recall interface {
	Find(point string, called string) any
	ObtainDriver(driver string) any
}
