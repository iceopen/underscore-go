package underscorego

import "reflect"

// 判读是否是个方法
func (u *U) IsFunction(arg interface{}) bool {
	if arg == nil {
		return false
	}
	return reflect.TypeOf(arg).Kind() == reflect.Func
}
