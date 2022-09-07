package functional

import (
	"errors"
	"log"
	"testing"

	f "github.com/noam-g4/functional"
)

func divide(a, b float32) (error, float32) {
	if b == 0 {
		return errors.New("cannot divide by zero!"), 0
	}
	return nil, a / b
}

func greet(name string) (error, string) {
	if name == "" {
		return errors.New("no name"), ""
	}
	return nil, "hi " + name
}

func TestMaybeMonad(t *testing.T) {

	_, y := divide(7, 5)
	resOk := f.Maybe(divide(7, 5))
	if resOk.Value != y {
		t.Error()
	}

	res := f.Maybe(divide(5, 0)).HandleErr(log.Println)
	if res.Err == nil {
		t.Error(res.Error())
	}

	str := f.Maybe(greet("john")).HandleErr(log.Println).Value
	if str != "hi john" {
		t.Error(str)
	}
}
