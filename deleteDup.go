package convpack

func DeleteDup[T comparable](slice []T) []T {
	ret := []T{}
	m := make(map[T]struct{})

	for _, val := range slice {
		if _, ok := m[val]; !ok {
			m[val] = struct{}{}
			ret = append(ret, val)
		}
	}

	return ret
}
