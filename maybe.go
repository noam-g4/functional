package functional

type Maybe[a any] struct {
	Value a
	Error error
}

func Try[a any](val a, err error) Maybe[a] {
	return Maybe[a]{
		Value: val,
		Error: err,
	}
}

func Then[a, b any](f func(a) (b, error), m Maybe[a]) Maybe[b] {
	if m.Error != nil {
		return Maybe[b]{
			Error: m.Error,
		}
	}
	return Try(f(m.Value))
}

func (m Maybe[a]) Catch(f func(error)) Maybe[a] {
	if m.Error != nil {
		f(m.Error)
	}
	return m
}
