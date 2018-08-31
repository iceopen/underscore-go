package underscorego

import "fmt"

// 初始化结构体
type U struct {
}

const (
	Version = "0.0.1.dev"
)

// 判断是否为空
func (u U) IsEmpty(i interface{}) (bool, error) {
	var isEmpty bool
	isEmpty = false
	switch i.(type) {
	case string:
		if i == "" {
			isEmpty = true
		}
		break
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		str := fmt.Sprintf("%d", i)
		if str == "0" || str == "-0" {
			isEmpty = true
		}
		break
	case []string:
		if len(i.([]string)) == 0 {
			isEmpty = true
		}
	}
	return isEmpty, nil
}

// 判断是否是数字
func (u U) IsNumber(i interface{}) (bool, error) {
	var isNumber bool
	isNumber = false
	switch i.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		isNumber = true
	}
	return isNumber, nil
}

// 判断是否是字符串
func (u U) IsString(i interface{}) (bool, error) {
	var isString bool
	isString = false
	switch i.(type) {
	case string:
		isString = true
	}
	return isString, nil
}
