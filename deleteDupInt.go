package convpack

func DeleteDupInt(slice []int) []int {
	ret := []int{}
	m := make(map[int]struct{})

	for _, val := range slice {
		if _, ok := m[val]; !ok {
			m[val] = struct{}{}
			ret = append(ret, val)
		}
	}

	return ret
}
