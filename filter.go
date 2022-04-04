package convpack

func Filter[T any](elems []T, f func(T) bool) []T {
	var res []T

	for _, e := range elems {
		if f(e) {
			res = append(res, e)
		}
	}

	return res
}
