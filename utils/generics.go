package utils

import (
	"cmp"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

func GenericMax[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func GenericPow[T Number](a, b T) T {
	ret := T(1)
	for i := T(0); i < b; i++ {
		ret *= a
	}
	return ret
}

func GenericMin[T cmp.Ordered](a []T) T {
	ret := a[0]
	for _, v := range a {
		if v < ret {
			ret = v
		}
	}
	return ret
}
