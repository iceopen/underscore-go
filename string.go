package underscorego

// 判断是否是字符串
func (u *U) IsString(i interface{}) bool {
	var isString bool
	isString = false
	switch i.(type) {
	case string:
		isString = true
	}
	return isString
}
