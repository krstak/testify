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
	True     func(item interface{})
	False    func(item interface{})
}

func New(t *testing.T) assert {
	return assert{
		Equal:    pairCheck(t, isEqual),
		NotEqual: pairCheck(t, isNotEqual),
		Nil:      monoCheck(t, nil, isEqual),
		NotNil:   monoCheck(t, nil, isNotEqual),
		True:     monoCheck(t, true, isEqual),
		False:    monoCheck(t, true, isNotEqual),
	}
}

func isEqual(exp, act interface{}) bool {
	return reflect.DeepEqual(exp, act)
}

func isNotEqual(exp, act interface{}) bool {
	return !reflect.DeepEqual(exp, act)
}

func monoCheck(t *testing.T, exp interface{}, eq func(exp, act interface{}) bool) func(item interface{}) {
	return func(item interface{}) {
		pairCheck(t, eq)(exp, item)
	}
}

func pairCheck(t *testing.T, cond func(exp, act interface{}) bool) func(exp, act interface{}) {
	return func(exp, act interface{}) {
		if !cond(exp, act) {
			printErr(t, exp, act)
		}
	}
}

var printErr = func(t *testing.T, exp, act interface{}) {
	_, b, l, _ := runtime.Caller(2)
	path := fmt.Sprintf("%s:%d", b, l)
	t.Errorf("\n%s: \nexp value: %v (%s)\nact value: %v (%s)\n", path, exp, reflect.TypeOf(exp), act, reflect.TypeOf(act))
}
