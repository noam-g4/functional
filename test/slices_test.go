package functional

import (
	"reflect"
	"testing"

	f "github.com/noam-g4/functional"
)

func TestEmptySet(t *testing.T) {
	intSet := f.EmptySet[int]()
	floatSet := f.EmptySet[float32]()
	boolSet := f.EmptySet[bool]()
	stringSet := f.EmptySet[string]()
	mapSet := f.EmptySet[map[string]interface{}]()

	if reflect.TypeOf(intSet) != reflect.TypeOf([]int{}) {
		t.Fail()
	}
	if reflect.TypeOf(floatSet) != reflect.TypeOf([]float32{}) {
		t.Fail()
	}
	if reflect.TypeOf(boolSet) != reflect.TypeOf([]bool{}) {
		t.Fail()
	}
	if reflect.TypeOf(stringSet) != reflect.TypeOf([]string{}) {
		t.Fail()
	}
	if reflect.TypeOf(mapSet) != reflect.TypeOf([]map[string]interface{}{}) {
		t.Fail()
	}
}

func TestMap(t *testing.T) {
	intSet := []int{1, 2, 3}
	stringSet := []string{"foo", "bar"}

	intRes := f.Map(intSet, func(x int) int {
		return x + 1
	}, f.EmptySet[int]())

	stringRes := f.Map(stringSet, func(s string) string {
		return s + "!"
	}, f.EmptySet[string]())

	for n, x := range intSet {
		if (x + 1) != intRes[n] {
			t.Fail()
		}
	}
	for n, s := range stringSet {
		if (s + "!") != stringRes[n] {
			t.Fail()
		}
	}
}

func TestIsIn(t *testing.T) {
	floatSet := []float32{0.24, 0.5, 0.9, -0.33}
	stringSet := []string{"hello", "world"}

	if !f.IsIn(0.9, floatSet) {
		t.Fail()
	}
	if f.IsIn(0.13, floatSet) {
		t.Fail()
	}
	if !f.IsIn("hello", stringSet) {
		t.Fail()
	}
}

func TestFilter(t *testing.T) {
	stringSet := []string{"one", "two", "three", "four"}
	filtered := f.Filter(stringSet, func(s string) bool {
		return len(s) != 3
	}, f.EmptySet[string]())

	if f.IsIn("one", filtered) || f.IsIn("two", filtered) || len(filtered) != 2 {
		t.Fail()
	}
}

func TestReduce(t *testing.T) {
	intSet := []int{1, 2, 3, 4}
	floatSet := []float32{5.5, 10, 2.5}
	sum := f.Reduce(intSet, func(y, x int) int {
		return y + x
	}, 0)

	doubleSum := f.Reduce(floatSet, func(y, x float32) float32 {
		return y + x*2
	}, 0)

	if sum != 10 {
		t.Fail()
	}

	if doubleSum != 36 {
		t.Fail()
	}
}
