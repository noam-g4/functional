package functional

func isEmptySlice[T any](s []T) bool {
	if len(s) == 0 {
		return true
	}
	return false
}

func push[T any](y []T, x T) []T {
	return append(y, x)
}
