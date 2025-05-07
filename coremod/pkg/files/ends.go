package files

type FileSource interface {
}

type FileDest interface {
	Pour(name string) // TODO: return some means of actually copying it ...
}

type ThingyHolder interface {
	ObtainDest() FileDest
}
