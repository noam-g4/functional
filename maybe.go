package functional

/*
the Either type is a functor that holds an error type and
a generic value. it's also in itself a valid golang error type
since it implements the error interface.
*/
type Either struct {
	Err   error
	Value any
}

// implementation of the Error() method to return the error message
func (e Either) Error() string {
	return e.Err.Error()
}

/*
an Either method that takes a handler function and produces
a desired effect by calling this handler function and passing
the Either.Err through it
the handler function produces effects by its nature, thus the type signature
is as defined below
*/
func (e Either) HandleErr(h func(...interface{})) Either {
	if e.Err != nil {
		h(e.Err)
	}
	return e
}

/*
a monad type that lets you pass an error and a value of any type
**note the you can pass function call that returns (error, any)
and based on the result of the inner function, it returns an Either type
with either an error or the value. the combinations of the Maybe and Either
lets you handle error as in the following example:

y := Maybe(Divide(5, 0)).HandleErr(log.Println).Value()
*/
func Maybe[a any](err error, val a) Either {
	if err != nil {
		return Either{Err: err}
	}
	return Either{Value: val}
}
