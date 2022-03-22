package convpack

func DeleteDupStr(slice []string) []string {
	ret := []string{}
	m := make(map[string]struct{})

	for _, val := range slice {
		if _, ok := m[val]; !ok {
			m[val] = struct{}{}
			ret = append(ret, val)
		}
	}

	return ret
}
