package other

func Filter[T1 any](ss []T1, filterFunc func(data T1) bool) []T1 {
	ret := make([]T1, 0)
	for _, s := range ss {
		if filterFunc(s) {
			ret = append(ret, s)
		}
	}
	return ret
}
