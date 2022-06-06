package functional

// takes an initial value and functions of the same input->output
// type and chain pipe the intials value through all of the functions
// e.g: Pipe(5, h, g, f) => f(g(h(5)))
func Pipe[T any](x T, fns ...func(T) T) T {
	if IsEmptySlice(fns) {
		return x
	}
	return Pipe(fns[0](x), fns[1:]...)
}

// levarages the Pipe function with currying.
// this function takes functions of the same signature
// and returns a function that can be called later with a value.
// this value is then passed through a Pipe function
// composed with the functions passed in the Compose function call
func Compose[T any](fns ...func(T) T) func(T) T {
	return func(x T) T {
		return Pipe(x, fns...)
	}
}
