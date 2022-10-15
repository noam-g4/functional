package functional

import (
	"errors"
	"log"
	"os"
	"strconv"
	"testing"

	f "github.com/noam-g4/functional"
)

func getEnv(name string) (error, string) {
	if e := os.Getenv(name); e == "" {
		return errors.New("empty var"), ""
	} else {
		return nil, e
	}
}

func parseFloat(str string) (error, float64) {
	if x, err := strconv.ParseFloat(str, 32); err != nil {
		return errors.New("cannont parse to float"), 0
	} else {
		return nil, x
	}
}

func fiveDivideBy(x float64) (error, float64) {
	if x == 0 {
		return errors.New("undefiend"), 0
	}
	return nil, 5 / x
}

func TestTry(t *testing.T) {

	os.Setenv("NUM", "asd")

	e := f.Try(getEnv("NUM"))
	x := f.Then(e, parseFloat)
	y := f.Then(x, fiveDivideBy)
	f.HandleError(y, log.Println)

	if y.Error == nil {
		t.Error(y)
	}

	os.Setenv("NUM", "5")

	e1 := f.Try(getEnv("NUM"))
	x1 := f.Then(e1, parseFloat)
	y1 := f.Then(x1, fiveDivideBy)

	if y1.Value != 1 {
		t.Error(y1)
	}

}
