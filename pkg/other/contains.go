package other

func Exist[T any, PtrT *T](ss []PtrT, filterFunc func(data PtrT) bool) (int, PtrT) {
	for i, s := range ss {
		if filterFunc(s) {
			return i, s
		}
	}
	return 0, nil
}
