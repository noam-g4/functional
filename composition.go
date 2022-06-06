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
