package common

import "reflect"

/*
判断量字符串切片是否相等
如其他类型可用：reflect.DeepEqual(a, b)
*/
func CompareStringSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	if (a == nil) != (b == nil) {
		return false
	}
	for key, value := range a {
		if value != b[key] {
			return false
		}
	}
	return true
}

// 判断切片是否存在元素
func IsExsitItem(v interface{}, s interface{}) bool {
	switch reflect.TypeOf(s).Kind() {
	case reflect.Slice:
		sl := reflect.ValueOf(s)
		for i := 0; i < sl.Len(); i++ {
			if reflect.DeepEqual(v, sl.Index(i).Interface()) {
				return true
			}
		}
	}
	return false
}
