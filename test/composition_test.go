package functional

import (
	"strings"
	"testing"

	f "github.com/noam-g4/functional"
)

func add5(x int) int {
	return x + 5
}
func sub2(x int) int {
	return x - 2
}
func addExMark(str string) string {
	return str + "!"
}
func capitalize(str string) string {
	return strings.Title(str)
}

func TestPipe(t *testing.T) {
	y1 := f.Pipe(5, add5, sub2)
	y2 := f.Pipe("foo", capitalize, addExMark)

	if y1 != 8 && y2 != "Foo!" {
		t.Fail()
	}
}
