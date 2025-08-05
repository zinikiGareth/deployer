package basic

import (
	"ziniki.org/deployer/coremod/pkg/corebottom"
	"ziniki.org/deployer/driver/pkg/errorsink"
)

type dummy struct {
}

// NotFound implements corebottom.ValuePresenter.
func (d *dummy) NotFound() {
}

// Present implements corebottom.ValuePresenter.
func (d *dummy) Present(value any) {
}

// WantDestruction implements corebottom.ValuePresenter.
func (d *dummy) WantDestruction(loc *errorsink.Location) {
}

func NewDummyPresenter() corebottom.ValuePresenter {
	return &dummy{}
}
