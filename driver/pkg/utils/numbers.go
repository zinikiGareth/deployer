package utils

import "ziniki.org/deployer/driver/pkg/driverbottom"

type f64AsNumber struct {
	val float64
}

func (h f64AsNumber) F64() float64 {
	return h.val
}

func F64AsNumber(val float64) driverbottom.AsNumber {
	return f64AsNumber{val: val}
}
