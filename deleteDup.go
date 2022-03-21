package convpack

import (
	"reflect"
)

func DeleteDup(slice interface{}) []interface{} {
	var ret []interface{}
	slice_itf := reflect.ValueOf(slice)
	var aMapType = reflect.MapOf(reflect.TypeOf(slice_itf.Index(0).Interface()), reflect.TypeOf(struct{}{}))
	aMap := reflect.MakeMapWithSize(aMapType, 0)

	for i := 0; i < slice_itf.Len(); i++ {
		val := slice_itf.Index(i).Interface()
		if !aMap.MapIndex(reflect.ValueOf(val)).IsValid() {
			aMap.SetMapIndex(reflect.ValueOf(val), reflect.ValueOf(struct{}{}))
			ret = append(ret, val)
		}
	}

	return ret
}
