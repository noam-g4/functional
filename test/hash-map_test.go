package functional

import (
	"log"
	"testing"

	f "github.com/noam-g4/functional"
)

type E struct {
	Value int
	Error error
	Nstd  Nested
}

type Nested struct {
	NVal string
}

type Empty struct {
	thisIsPrivate string
}

func TestToHashMap(t *testing.T) {
	v := E{
		Value: 5,
		Error: nil,
	}
	m := f.ToHashMap(v)
	if m.Value.(f.HashMap)["Value"] != 5 {
		t.Error(m)
	}

	d := 5
	m2 := f.ToHashMap(d).HandleErr(log.Println)
	if m2.Err == nil {
		t.Fail()
	}

	e := Empty{}
	m3 := f.ToHashMap(e)
	if m3.Err != nil {
		t.Error(m3)
	}
}

func TestGetValue(t *testing.T) {
	e := E{
		Value: 5,
		Nstd: Nested{
			NVal: "Nested",
		},
	}
	val1 := f.GetValue[int]("Value", f.ToHashMap(e))
	if val1.Value != 5 {
		t.Error(val1)
	}
	val2 := f.GetValue[any]("Undefined", f.ToHashMap(e)).HandleErr(log.Println)
	if val2.Err == nil {
		t.Fail()
	}
	nstd := f.GetValue[Nested]("Nstd", f.ToHashMap(e)).Value
	val3 := f.GetValue[string]("NVal", f.ToHashMap(nstd))
	if val3.Value != "Nested" {
		t.Error(val3)
	}
}
