package utils

import "fmt"

type StringerString struct {
	val string
}

func AsStringer(val string) StringerString {
	return StringerString{val: val}
}

func (stringer StringerString) String() string {
	return stringer.val
}

type deferredReading struct {
	f func() string
}

func (dr deferredReading) String() string {
	return dr.f()
}

func DeferReading(reader func() string) fmt.Stringer {
	return deferredReading{f: reader}
}
