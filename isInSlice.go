package convpack

import "reflect"

func IsInSlice[T comparable](element T, slice []T) bool {
	for _, target := range slice {
		if reflect.ValueOf(target).Interface() == reflect.ValueOf(element).Interface() {
			return true
		}
	}
	return false
}
