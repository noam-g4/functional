package functional

import (
	"errors"
	"testing"

	f "github.com/noam-g4/functional"
)

func TestMaybeDeclaration(t *testing.T) {
	intOrString := f.Maybe[int, string]{
		Left:  5,
		Right: "foo",
	}

	if intOrString.Left != 5 && intOrString.Right != "foo" {
		t.Fail()
	}
}

func safeDiv(a, b float32) f.Maybe[float32, error] {
	var y f.Maybe[float32, error]
	if b == 0 {
		y.Right = errors.New("cannot divide by 0")
		return y
	}
	y.Left = a / b
	return y
}

func TestErrorUseCase(t *testing.T) {
	noErr := safeDiv(23.75, 7)
	withErr := safeDiv(2.5, 0)

	if noErr.Right != nil &&
		withErr.Right == nil &&
		noErr.Left > 0 &&
		withErr.Left == 0 {
		t.Fail()
	}
}

func TestDestructMaybe(t *testing.T) {
	m1 := safeDiv(5, 0)
	m2 := f.Maybe[string, bool]{
		Left:  "bar",
		Right: true,
	}

	flt, err := f.DestructMaybe(m1)
	str, bool := f.DestructMaybe(m2)

	if flt != 0 &&
		err.Error() != "cannot divide by 0" &&
		str != "bar" &&
		!bool {
		t.Fail()
	}
}
