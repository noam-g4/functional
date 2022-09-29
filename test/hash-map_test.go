package functional

import (
	"log"
	"testing"

	f "github.com/noam-g4/functional"
)

type E struct {
	Value int
	Error error
}

type Empty struct {
	thisIsPrivate string
}

func TestToHashMap(t *testing.T) {
	v := E{
		Value: 5,
		Error: nil,
	}
	_, m := f.ToHashMap(v)
	if m["Value"] != 5 {
		t.Error(m, m["Value"])
	}

	d := 5
	err, _ := f.ToHashMap(d)
	if err == nil {
		t.Fail()
	}

	e := Empty{}
	err, _ = f.ToHashMap(e)
	if err != nil {
		t.Error(err)
	}
}

func TestGetValue(t *testing.T) {
	e := E{
		Value: 5,
	}
	_, val := f.GetValue("Value", e)
	if val != 5 {
		t.Error(val)
	}
	err, _ := f.GetValue("Undefined", e)
	if err == nil {
		t.Fail()
	}
}

func TestGetValueWithEitherMonad(t *testing.T) {
	e := E{
		Value: 10,
	}
	res1 := f.Try(f.GetValue("Value", e)).HandleErr(log.Println)
	if res1.Value != 10 {
		t.Error(res1)
	}
	res2 := f.Try(f.GetValue("Undefined", e)).HandleErr(log.Println)
	if res2.Err == nil {
		t.Error(res2)
	}
	res3 := f.Try(f.GetValue("SomeKey", "not a struct")).HandleErr(log.Println)
	if res3.Err == nil {
		t.Error(res3)
	}
}
