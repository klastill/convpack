package convpack

func GetUnion[T comparable](arr ...[]T) []T {
	var res []T
	m := make(map[T]bool)

	for _, a := range arr {
		for _, item := range a {
			if _, exist := m[item]; !exist {
				m[item] = true
			}
		}
	}
	for key := range m {
		res = append(res, key)
	}

	return res
}
