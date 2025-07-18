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

func AsF64(val any) (float64, bool) {
	switch val := val.(type) {
	case float64:
		return val, true
	case float32:
		return float64(val), true
	case int64:
		return float64(val), true
	case int32:
		return float64(val), true
	case int:
		return float64(val), true
	default:
		return 0, false
	}
}

func AsI32(val any) (int32, bool) {
	switch val := val.(type) {
	case float64:
		return int32(val), true
	case float32:
		return int32(val), true
	case int64:
		return int32(val), true
	case int32:
		return val, true
	case int:
		return int32(val), true
	default:
		return 0, false
	}
}
