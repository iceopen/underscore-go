package underscorego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 判断是否为空
func TestIsEmpty(t *testing.T) {
	var u U
	// 测试字符串为空
	is := u.IsEmpty("")
	assert.Equal(t, is, true)
	// 测试字符串不为空
	is = u.IsEmpty("test")
	assert.Equal(t, is, false)
	is = u.IsEmpty(0)
	assert.Equal(t, is, true)
	is = u.IsEmpty(1)
	assert.Equal(t, is, false)
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
}
