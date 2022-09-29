package functional

import (
	"errors"
	"fmt"
	"reflect"
)

type HashMap map[string]any

// getting the value of a key in an unknown struct
func GetValue(key string, o any) (error, any) {
	err, m := ToHashMap(o)
	if err != nil {
		return err, nil
	}
	val := m[key]
	if val == nil {
		return errors.New(fmt.Sprintf("value of key: '%s' is not found", key)), nil
	}
	return nil, val
}

// converts any struct type to a hash-map map[string]any
func ToHashMap(o any) (error, HashMap) {
	if !isStruct(o) {
		return errors.New("cannot convert a non-struct type to a hash-map"), nil
	}
	fields := reflect.VisibleFields(reflect.TypeOf(o))
	return nil, Reduce(fields, func(y HashMap, f reflect.StructField) HashMap {
		if f.IsExported() {
			y[f.Name] = reflect.ValueOf(o).FieldByName(f.Name).Interface()
		}
		return y
	}, make(HashMap))
}

func isStruct(o any) bool {
	return reflect.TypeOf(o).Kind() == reflect.Struct
}
