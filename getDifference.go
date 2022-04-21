package convpack

func GetDefference[T comparable](main []T, target []T) []T {
	var res []T
	m := make(map[T]bool)

	for _, item := range main {
		m[item] = true
	}
	for _, item := range target {
		m[item] = false
	}
	for key, value := range m {
		if value {
			res = append(res, key)
		}
	}

	return res
}
