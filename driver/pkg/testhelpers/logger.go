package testhelpers

type TestStepLogger interface {
	Log(fmt string, args ...any)
}
