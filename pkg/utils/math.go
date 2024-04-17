package utils

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Float | constraints.Integer
}

func Sum[N Number](numbers []N) N {
	var result N = 0

	for _, n := range numbers {
		result += n
	}

	return N(result)
}
