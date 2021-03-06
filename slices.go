package functional

// returns an empty slice of any type
func EmptySet[T any]() []T {
	set := []T{}
	return set
}

// checks if a slice is empty
func IsEmptySlice[T any](s []T) bool {
	if len(s) == 0 {
		return true
	}
	return false
}

// takes a slice and an element and returns a copy
// of the slice with the element appended to it
func AppendToSlice[T any](y []T, x T) []T {
	return append(y, x)
}

// takes a slice, function and empty EmptySet
// and returns a new slice of which each element
// has being passed through the function parameter
func Map[A, B any](xs []A, f func(A) B, y []B) []B {
	if IsEmptySlice(xs) {
		return y
	}
	return Map(xs[1:], f, AppendToSlice(y, f(xs[0])))
}

// takes a slice, predicate function and empty EmptySet
// and returns a new slice of only the elements that have
// resulted true when passed to the predicate function
func Filter[T any](xs []T, p func(T) bool, y []T) []T {
	if IsEmptySlice(xs) {
		return y
	}
	if p(xs[0]) {
		return Filter(xs[1:], p, AppendToSlice(y, xs[0]))
	}
	return Filter(xs[1:], p, y)
}

// takes a slice, a function that takes an accumalator and
// and an element of the slice, an initial value for the accumalator
// and returns the accumalated result,
// govern by the parameter function
func Reduce[X, Y any](xs []X, f func(Y, X) Y, y Y) Y {
	if IsEmptySlice(xs) {
		return y
	}
	return Reduce(xs[1:], f, f(y, xs[0]))
}

// takes a single element and a slice
// and checks if this element is in the slice
func IsIn[T comparable](x T, xs []T) bool {
	if IsEmptySlice(xs) {
		return false
	}
	if x == xs[0] {
		return true
	}
	return IsIn(x, xs[1:])
}

// takes two slices of the same type
// and returns a concatanated slice [...s1, ...s2]
func ConcatSlices[T any](s1, s2 []T) []T {
	return Map(s2, func(x T) T {
		return x
	}, s1)
}
