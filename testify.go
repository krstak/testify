package testify

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

type assert struct {
	Equal    func(exp, act interface{})
	NotEqual func(exp, act interface{})
	Nil      func(item interface{})
	NotNil   func(item interface{})
}

func New(t *testing.T) assert {
	return assert{
		Equal:    isEqual(t),
		NotEqual: isNotEqual(t),
		Nil:      isNil(t),
		NotNil:   isNotNil(t),
	}
}

func isEqual(t *testing.T) func(exp, act interface{}) {
	return func(exp, act interface{}) {
		if !reflect.DeepEqual(exp, act) {
			printErr(t, exp, act)
		}
	}
}

func isNotEqual(t *testing.T) func(exp, act interface{}) {
	return func(exp, act interface{}) {
		if reflect.DeepEqual(exp, act) {
			printErr(t, exp, act)
		}
	}
}

func isNil(t *testing.T) func(item interface{}) {
	return func(item interface{}) {
		if item != nil {
			printErr(t, nil, item)
		}
	}
}

func isNotNil(t *testing.T) func(item interface{}) {
	return func(item interface{}) {
		if item == nil {
			printErr(t, nil, item)
		}
	}
}

var printErr = func(t *testing.T, exp, act interface{}) {
	_, b, l, _ := runtime.Caller(2)
	path := fmt.Sprintf("%s:%d", b, l)
	t.Errorf("\n%s: \nexp value: %v (%s)\nact value: %v (%s)\n", path, exp, reflect.TypeOf(exp), act, reflect.TypeOf(act))
}
