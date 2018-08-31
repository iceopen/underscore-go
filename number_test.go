package underscorego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 判断是否是数字
func TestIsNumber(t *testing.T) {
	var u U
	// 测试字符串为空
	is := u.IsNumber("")
	assert.Equal(t, is, false)
	// 测试字符串不为空
	is = u.IsNumber("test")
	assert.Equal(t, is, false)
	is = u.IsNumber(0)
	assert.Equal(t, is, true)
	is = u.IsNumber(1)
	assert.Equal(t, is, true)
	var ii int32
	ii = 0
	is = u.IsNumber(ii)
	assert.Equal(t, is, true)

	strArr1 := []string{}
	is = u.IsNumber(strArr1)
	assert.Equal(t, is, false)
	strArr2 := []string{""}
	is = u.IsNumber(strArr2)
	assert.Equal(t, is, false)
}
