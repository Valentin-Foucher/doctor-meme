package utils

func SliceContains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func MapSlice[T interface{}, K interface{}](s []T, f func(T) K) []K {
	mapped := make([]K, len(s))

	for i, e := range s {
		mapped[i] = f(e)
	}

	return mapped
}
