package underscorego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 判断是否为空
func TestU_IsEmpty(t *testing.T) {
	var u U
	// 测试字符串为空
	is := u.IsEmpty("")
	assert.Equal(t, is, true, "空字符串判断")
	// 测试字符串不为空
	is = u.IsEmpty("test")
	assert.Equal(t, is, false, "非空字符串判断")
	is = u.IsEmpty(0)
	assert.Equal(t, is, true, "0判断")
	is = u.IsEmpty(1)
	assert.Equal(t, is, false, "1判断")
	var ii int32
	ii = 0
	is = u.IsEmpty(ii)
	assert.Equal(t, is, true)

	strArr1 := []string{}
	is = u.IsEmpty(strArr1)
	assert.Equal(t, is, true)
	strArr2 := []string{""}
	is = u.IsEmpty(strArr2)
	assert.Equal(t, is, false)
	is = u.IsEmpty(make(chan struct{}))
	assert.Equal(t, is, true)
}

func TestU_IsNil(t *testing.T) {
	var u U
	is := u.IsNil(nil)
	assert.Equal(t, is, true, "nil 是否是nil判断")
	// 字符串判断
	is = u.IsNil("")
	assert.Equal(t, is, false, "空字符串 是否是nil判断")
}
