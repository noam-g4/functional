package functional

import (
	"errors"
	"fmt"
	"reflect"
)

type HashMap map[string]any

// getting the value of a key in an unknown struct
func GetValue[T any](key string, m Either) Either {
	if m.Err != nil {
		return m
	}
	if h, ok := m.Value.(HashMap); !ok {
		return Either{
			Value: m.Value,
			Err:   errors.New("Either.Value is not a HashMap"),
		}
	} else {
		return Try(getValuePrime[T](key, h))
	}
}

func getValuePrime[T any](key string, m HashMap) (error, any) {
	if m[key] == nil {
		return errors.New(fmt.Sprintf("value of key: '%s' is not found", key)), nil
	}
	if v, can := m[key].(T); !can {
		return errors.New(fmt.Sprintf("cannot assert Either.Value type")), nil
	} else {
		return nil, v
	}
}

// converts any struct type to a hash-map map[string]any
func ToHashMap(o any) Either {
	if s := Try(isStruct(o)); s.Err != nil {
		return s
	}
	return Either{
		Value: Reduce(reflect.VisibleFields(reflect.TypeOf(o)),
			func(y HashMap, f reflect.StructField) HashMap {
				if f.IsExported() {
					y[f.Name] = reflect.ValueOf(o).FieldByName(f.Name).Interface()
				}
				return y
			}, make(HashMap))}
}

func isStruct(o any) (error, bool) {
	if reflect.TypeOf(o).Kind() != reflect.Struct {
		return errors.New("cannot convert a non-struct type to a hash-map"), false
	}
	return nil, true
}
