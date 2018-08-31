package underscorego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 判断是否为空
func TestIsEmpty(t *testing.T) {
	var u U
	// 测试字符串为空
	is, _ := u.IsEmpty("")
	assert.Equal(t, is, true)
	// 测试字符串不为空
	is, _ = u.IsEmpty("test")
	assert.Equal(t, is, false)
	is, _ = u.IsEmpty(0)
	assert.Equal(t, is, true)
	is, _ = u.IsEmpty(1)
	assert.Equal(t, is, false)
	var ii int32
	ii = 0
	is, _ = u.IsEmpty(ii)
	assert.Equal(t, is, true)

	strArr1 := []string{}
	is, _ = u.IsEmpty(strArr1)
	assert.Equal(t, is, true)
	strArr2 := []string{""}
	is, _ = u.IsEmpty(strArr2)
	assert.Equal(t, is, false)
}

// 判断是否是数字
func TestIsNumber(t *testing.T) {
	var u U
	// 测试字符串为空
	is, _ := u.IsNumber("")
	assert.Equal(t, is, false)
	// 测试字符串不为空
	is, _ = u.IsNumber("test")
	assert.Equal(t, is, false)
	is, _ = u.IsNumber(0)
	assert.Equal(t, is, true)
	is, _ = u.IsNumber(1)
	assert.Equal(t, is, true)
	var ii int32
	ii = 0
	is, _ = u.IsNumber(ii)
	assert.Equal(t, is, true)

	strArr1 := []string{}
	is, _ = u.IsNumber(strArr1)
	assert.Equal(t, is, false)
	strArr2 := []string{""}
	is, _ = u.IsNumber(strArr2)
	assert.Equal(t, is, false)
}
