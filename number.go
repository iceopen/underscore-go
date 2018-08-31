package underscorego

// 判断是否是数字
func (u *U) IsNumber(i interface{}) bool {
	var isNumber bool
	isNumber = false
	switch i.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		isNumber = true
	}
	return isNumber
}
