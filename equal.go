package underscorego

import (
	"errors"
	"reflect"
	"bytes"
)

// 判断值是否相同
// 传参类型不能是方法类型
func (u *U) IsEqual(expected, actual interface{}) (bool, error) {
	if u.IsFunction(expected) || u.IsFunction(actual) {
		return false, errors.New("cannot take func type as argument")
	}
	if !ObjectsAreEqual(expected, actual) {
		return false, nil
	}
	return true, nil
}

func ObjectsAreEqual(expected, actual interface{}) bool {
	if expected == nil || actual == nil {
		return expected == actual
	}

	exp, ok := expected.([]byte)
	if !ok {
		return reflect.DeepEqual(expected, actual)
	}

	act, ok := actual.([]byte)
	if !ok {
		return false
	}
	if exp == nil || act == nil {
		return exp == nil && act == nil
	}
	return bytes.Equal(exp, act)
}
