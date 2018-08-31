package underscorego

import "fmt"

// 判断是否为空
func (u *U) IsEmpty(i interface{}) bool{
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
	return isEmpty
}
