package functional

// checks if a slice is empty
func IsEmptySlice[T any](s []T) bool {
	return len(s) == 0
}

// takes a slice and an element and returns a copy
// of the slice with the element appended to it
func AppendToSlice[T any](y []T, x T) []T {
	return append(y, x)
}

// takes a slice, function and empty EmptySet
// and returns a new slice of which each element
// has being passed through the function parameter
func Map[A, B any](xs []A, f func(A) B) []B {
	b := make([]B, len(xs))
	for i, x := range xs {
		b[i] = f(x)
	}
	return b
}

// takes a slice, predicate function and empty EmptySet
// and returns a new slice of only the elements that have
// resulted true when passed to the predicate function
func Filter[T any](xs []T, p func(T) bool) []T {
	out := make([]T, 0)
	for _, x := range xs {
		if p(x) {
			out = append(out, x)
		}
	}
	return out
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
	return append(s1, s2...)
}

// takes a silce
// and returns a slice with only distinct value
func NewSet[T comparable](slc []T) []T {
	m := make(map[T]bool)
	y := make([]T, 0)
	for _, x := range slc {
		m[x] = true
	}
	for k := range m {
		y = append(y, k)
	}
	return y
}
