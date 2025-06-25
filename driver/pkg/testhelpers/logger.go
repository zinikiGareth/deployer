package testhelpers

import "ziniki.org/deployer/driver/pkg/driverbottom"

type TestStorer interface {
	GetWriter(mode int) driverbottom.IndentWriter
}

type TestStepLogger interface {
	Log(fmt string, args ...any)
}
