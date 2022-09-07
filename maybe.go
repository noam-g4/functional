package functional

type Either struct {
	Err   error
	Value any
}

func (e Either) Error() string {
	return e.Err.Error()
}

func (e Either) HandleErr(h func(...interface{})) Either {
	h(e)
	return e
}

func Maybe[a any](err error, val a) Either {
	if err != nil {
		return Either{Err: err}
	}
	return Either{Value: val}
}
