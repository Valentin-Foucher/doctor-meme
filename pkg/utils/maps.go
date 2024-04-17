package utils

func MergeMaps[K comparable, V interface{}](m1 map[K]V, m2 map[K]V) map[K]V {
	if m1 == nil {
		return m2
	}
	if m2 == nil {
		return m1
	}

	merged := make(map[K]V)
	for k, v := range m1 {
		merged[k] = v
	}
	for k, v := range m2 {
		merged[k] = v
	}
	return merged
}

func MapFromSlices[K comparable, V interface{}](keys []K, values []V) map[K]V {
	result := make(map[K]V)
	for i := range keys {
		result[keys[i]] = values[i]
	}
	return result
}

func MapKeys[K comparable, V interface{}](m map[K]V) []K {
	keys := make([]K, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

func MapValues[K comparable, V interface{}](m map[K]V) []V {
	values := make([]V, len(m))

	i := 0
	for _, v := range m {
		values[i] = v
		i++
	}
	return values
}
