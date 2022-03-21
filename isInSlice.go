package convpack

import "reflect"

func IsInSlice(element interface{}, slice interface{}) bool {
	slice_itf := reflect.ValueOf(slice)
	if slice_itf.Kind() == reflect.Slice {
		for i := 0; i < slice_itf.Len(); i++ {
			if slice_itf.Index(i).Interface() == element {
				return true
			}
		}
	}
	return false
}
