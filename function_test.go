package underscorego

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestU_IsFunction(t *testing.T) {
	var u U
	a := func() {}
	aIsFun := u.IsFunction(a)
	assert.Equal(t, true, aIsFun, "是否是方法判断错误")

	b := 1
	bIsFun := u.IsFunction(b)
	assert.Equal(t, false, bIsFun, "不是方法判断错误")

	nilIsFun := u.IsFunction(nil)
	assert.Equal(t, false, nilIsFun, "nil判断错误")
}
