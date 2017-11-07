package testify

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"
)

var Nil = mono(nil, isEqual)
var NotNil = mono(nil, isNotEqual)
var True = mono(true, isEqual)
var False = mono(false, isEqual)
var Equal = pair(isEqual, 2)
var NotEqual = pair(isNotEqual, 2)

func mono(exp interface{}, cond func(exp, act interface{}) bool) func(*testing.T) func(act interface{}) {
	return func(t *testing.T) func(act interface{}) {
		return func(act interface{}) {
			pair(cond, 3)(t)(exp, act)
		}
	}
}

func pair(cond func(exp, act interface{}) bool, skip int) func(*testing.T) func(exp, act interface{}) {
	return func(t *testing.T) func(exp, act interface{}) {
		return func(exp, act interface{}) {
			if !cond(exp, act) {
				printErr(t, skip, exp, act)
			}
		}
	}
}

func isEqual(exp, act interface{}) bool {
	return reflect.DeepEqual(exp, act)
}

func isNotEqual(exp, act interface{}) bool {
	return !isEqual(exp, act)
}

var printErr = func(t *testing.T, skip int, exp, act interface{}) {
	_, b, l, _ := runtime.Caller(skip)
	path := fmt.Sprintf("%s:%d", b, l)
	t.Errorf("\n%s: \nexp value: %v (%s)\nact value: %v (%s)\n", path, exp, reflect.TypeOf(exp), act, reflect.TypeOf(act))
}
