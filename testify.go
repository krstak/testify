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
var False = mono(true, isNotEqual)
var Equal = pair(isEqual)
var NotEqual = pair(isNotEqual)

func mono(exp interface{}, cond func(exp, act interface{}) bool) func(exp interface{}) func(*testing.T) {
	return func(act interface{}) func(t *testing.T) {
		return pair(cond)(exp, act)
	}
}

func pair(cond func(exp, act interface{}) bool) func(exp, act interface{}) func(*testing.T) {
	return func(exp, act interface{}) func(t *testing.T) {
		return func(t *testing.T) {
			if !cond(exp, act) {
				printErr(t, exp, act)
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

var printErr = func(t *testing.T, exp, act interface{}) {
	_, b, l, _ := runtime.Caller(2)
	path := fmt.Sprintf("%s:%d", b, l)
	t.Errorf("\n%s: \nexp value: %v (%s)\nact value: %v (%s)\n", path, exp, reflect.TypeOf(exp), act, reflect.TypeOf(act))
}
