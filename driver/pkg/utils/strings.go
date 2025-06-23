package utils

import (
	"fmt"
	"log"
)

type StringerString struct {
	val string
}

func AsStringer(val any) (fmt.Stringer, bool) {
	s, ok := val.(string)
	if ok {
		return StringerString{val: s}, true
	}
	sr, ok := val.(fmt.Stringer)
	if ok {
		return sr, true
	}
	log.Fatalf("Cannot convert to string: %T", val)
	return nil, false
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

func AsStringList(whatevs any) ([]string, bool) {
	isAlready, ok := whatevs.([]string)
	if ok {
		return isAlready, true
	}
	inList, ok := whatevs.([]any)
	if !ok {
		return nil, false
	}
	ret := []string{}
	for _, d := range inList {
		a, ok := d.(string)
		if !ok {
			return nil, false
		}
		ret = append(ret, a)
	}
	return ret, true
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
