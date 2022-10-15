package functional

type Maybe[a any] struct {
	Value a
	Error error
}

func Try[a any](err error, val a) Maybe[a] {
	if err != nil {
		return Maybe[a]{
			Error: err,
		}
	}
	return Maybe[a]{
		Value: val,
	}
}

func Then[a, b any](f func(a) (error, b), m Maybe[a]) Maybe[b] {
	if m.Error != nil {
		return Maybe[b]{
			Error: m.Error,
		}
	}
	return Try(f(m.Value))
}

func HandleError[a any](f func(...interface{}), m Maybe[a]) Maybe[a] {
	if m.Error != nil {
		f(m.Error)
	}
	return m
}
