package testify

import "testing"

func TestIsEqualTo(t *testing.T) {
	tests := []struct {
		exp interface{}
		act interface{}
		ok  bool
	}{
		{exp: 2, act: 2, ok: true},
		{exp: int64(2), act: int64(2), ok: true},
		{exp: int64(2), act: 2, ok: false},
		{exp: 2, act: int64(2), ok: false},
		{exp: 2.34, act: 2.34, ok: true},
		{exp: 2.34, act: "2.34", ok: false},
		{exp: float32(10), act: 10, ok: false},
		{exp: 2.34000000001, act: 2.34000000001, ok: true},
		{exp: 2.34000000001, act: 2.34000000002, ok: false},
		{exp: 2, act: "2", ok: false},
		{exp: "yeu7280;ö", act: "yeu7280;ö", ok: true},
		{exp: "yeu7280;ö", act: "we", ok: false},
	}

	for _, test := range tests {
		ok := true
		printErr = func(t *testing.T, skip int, exp, act interface{}) {
			ok = false
		}
		Equal(t)(test.exp, test.act)
		if ok != test.ok {
			t.Errorf("Expected %v, actual: %v", test.exp, test.act)
		}
	}
}

func TestIsNotEqualTo(t *testing.T) {
	tests := []struct {
		exp interface{}
		act interface{}
		ok  bool
	}{
		{exp: 2, act: 2, ok: false},
		{exp: int64(2), act: int64(2), ok: false},
		{exp: int64(2), act: 2, ok: true},
		{exp: 2, act: int64(2), ok: true},
		{exp: 2.34, act: 2.34, ok: false},
		{exp: 2.34, act: "2.34", ok: true},
		{exp: float32(10), act: 10, ok: true},
		{exp: 2.34000000001, act: 2.34000000001, ok: false},
		{exp: 2.34000000001, act: 2.34000000002, ok: true},
		{exp: 2, act: "2", ok: true},
		{exp: "yeu7280;ö", act: "yeu7280;ö", ok: false},
		{exp: "yeu7280;ö", act: "we", ok: true},
	}

	for _, test := range tests {
		ok := true
		printErr = func(t *testing.T, skip int, exp, act interface{}) {
			ok = false
		}
		NotEqual(t)(test.exp, test.act)
		if ok != test.ok {
			t.Errorf("Expected %v, actual: %v", test.exp, test.act)
		}
	}
}

func TestIsNil(t *testing.T) {
	tests := []struct {
		act   interface{}
		isNil bool
	}{
		{act: 2, isNil: false},
		{act: int64(2), isNil: false},
		{act: nil, isNil: true},
		{act: "", isNil: false},
		{act: 2.34, isNil: false},
		{act: 0, isNil: false},
		{act: map[string]string{}, isNil: false},
	}

	for _, test := range tests {
		ok := true
		printErr = func(t *testing.T, skip int, exp, act interface{}) {
			ok = false
		}
		Nil(t)(test.act)
		if ok != test.isNil {
			t.Errorf("Expected %v, actual: %v", test.isNil, ok)
		}
	}
}

func TestIsNotNil(t *testing.T) {
	tests := []struct {
		act      interface{}
		isNotNil bool
	}{
		{act: 2, isNotNil: true},
		{act: int64(2), isNotNil: true},
		{act: nil, isNotNil: false},
		{act: "", isNotNil: true},
		{act: 2.34, isNotNil: true},
		{act: 0, isNotNil: true},
		{act: map[string]string{}, isNotNil: true},
	}

	for _, test := range tests {
		ok := true
		printErr = func(t *testing.T, skip int, exp, act interface{}) {
			ok = false
		}
		NotNil(t)(test.act)
		if ok != test.isNotNil {
			t.Errorf("Expected %v, actual: %v", test.isNotNil, ok)
		}
	}
}

func TestIsTrue(t *testing.T) {
	tests := []struct {
		act    interface{}
		isTrue bool
	}{
		{act: 2, isTrue: false},
		{act: int64(2), isTrue: false},
		{act: nil, isTrue: false},
		{act: "", isTrue: false},
		{act: 2.34, isTrue: false},
		{act: true, isTrue: true},
		{act: map[string]string{}, isTrue: false},
		{act: false, isTrue: false},
		{act: (10 > 5), isTrue: true},
		{act: (10 == 5), isTrue: false},
		{act: (10 < 5), isTrue: false},
	}

	for _, test := range tests {
		ok := true
		printErr = func(t *testing.T, skip int, exp, act interface{}) {
			ok = false
		}
		True(t)(test.act)
		if ok != test.isTrue {
			t.Errorf("Expected %v, actual: %v", test.isTrue, ok)
		}
	}
}

func TestIsFalse(t *testing.T) {
	tests := []struct {
		act     interface{}
		isFalse bool
	}{
		{act: 2, isFalse: false},
		{act: int64(2), isFalse: false},
		{act: nil, isFalse: false},
		{act: "", isFalse: false},
		{act: 2.34, isFalse: false},
		{act: true, isFalse: false},
		{act: map[string]string{}, isFalse: false},
		{act: false, isFalse: true},
		{act: (10 > 5), isFalse: false},
		{act: (10 == 5), isFalse: true},
		{act: (10 < 5), isFalse: true},
	}

	for _, test := range tests {
		ok := true
		printErr = func(t *testing.T, skip int, exp, act interface{}) {
			ok = false
		}
		False(t)(test.act)
		if ok != test.isFalse {
			t.Errorf("Expected %v, actual: %v", test.isFalse, ok)
		}
	}
}
