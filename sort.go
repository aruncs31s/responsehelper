package responsehelper

import (
	"slices"
	"strings"
)

type WithName interface {
	String() string
}

func Sort[T WithName](in []T) []T {
	slices.SortFunc(in, func(a, b T) int {
		return strings.Compare(a.String(), b.String())
	})
	return in
}

func SortDesc[T WithName](in []T) []T {
	slices.SortFunc(in, func(a, b T) int {
		return strings.Compare(b.String(), a.String())
	})
	return in
}
