package underscorego

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestU_IsString(t *testing.T) {
	var u U
	isStr := u.IsString("")
	assert.Equal(t, true, isStr, "字符串判断错误")
	isStr = u.IsString(1)
	assert.Equal(t, false, isStr, "字符串判断错误")
}
