package utils

import (
	"fmt"
	"log"
)

type StringerString struct {
	val string
}

func AsStringer(val string) StringerString {
	return StringerString{val: val}
}

func AsString(obj any) string {
	k, isString := obj.(string)
	if isString {
		return k
	}
	l, isStringer := obj.(fmt.Stringer)
	if isStringer {
		return l.String()
	}
	log.Fatalf("Cannot convert to string: %T", obj)
	return ""
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

func DeferString(reader func() string) fmt.Stringer {
	return deferredReading{f: reader}
}
