package utils

import "cmp"

func GenericMax[T cmp.Ordered](a, b T) T {
	if a > b {
		return a
	}
	return b
}
