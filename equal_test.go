package underscorego

import "testing"

func TestObjectsAreEqual(t *testing.T) {
	var u U
	isEq, _ := u.IsEqual("Hello World", "Hello World")
	if ! isEq {
		t.Error("Equal should return true")
	}
	isEq, _ = u.IsEqual(123, 123)
	if ! isEq {
		t.Error("Equal should return true")
	}
	isEq, _ = u.IsEqual(123.5, 123.5)
	if ! isEq {
		t.Error("Equal should return true")
	}
	isEq, _ = u.IsEqual([]byte("Hello World"), []byte("Hello World"))
	if ! isEq {
		t.Error("Equal should return true")
	}
	isEq, _ = u.IsEqual(nil, nil)
	if ! isEq {
		t.Error("Equal should return true")
	}
	isEq, _ = u.IsEqual(int32(123), int32(123))
	if ! isEq {
		t.Error("Equal should return true")
	}
	isEq, _ = u.IsEqual(uint64(123), uint64(123))
	if ! isEq {
		t.Error("Equal should return true")
	}
	var m map[string]interface{}
	isEq, _ = u.IsEqual(m["bar"], "something")
	if isEq {
		t.Error("Equal should return false")
	}
}
