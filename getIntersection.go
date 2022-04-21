package convpack

func GetIntersection[T comparable](arr ...[]T) []T {
	var res []T
	m := make(map[T]int)

	for _, a := range arr {
		for _, item := range a {
			if _, exist := m[item]; exist {
				m[item]++
			} else {
				m[item] = 1
			}
		}
	}
	for key, value := range m {
		if value == len(arr) {
			res = append(res, key)
		}
	}

	return res
}
