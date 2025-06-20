package utils

type StringerString struct {
	val string
}

func AsStringer(val string) StringerString {
	return StringerString{val: val}
}

func (stringer StringerString) String() string {
	return stringer.val
}
