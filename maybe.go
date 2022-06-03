package functional

// a struct that packs together 2 values of any type
// this can be used, when a function can return two
// different types, depending on the condition
type Maybe[A, B any] struct {
	Left  A
	Right B
}

// returns a destructed version of the Maybe types
// separating Maybe.Left and Maybe.Right
func DestructMaybe[A, B any](m Maybe[A, B]) (A, B) {
	return m.Left, m.Right
}
