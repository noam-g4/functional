package functional

import (
	"errors"
	"log"
	"os"
	"strconv"
	"testing"

	f "github.com/noam-g4/functional"
)

func getEnv(name string) (string, error) {
	if e := os.Getenv(name); e == "" {
		return "", errors.New("empty var")
	} else {
		return e, nil
	}
}

func parseFloat(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

func fiveDivideBy(x float64) (float64, error) {
	if x == 0 {
		return 0, errors.New("undefiend")
	}
	return 5 / x, nil
}

func TestTry(t *testing.T) {

	os.Setenv("NUM", "asd")

	e := f.Try(getEnv("NUM"))
	x := f.Then(parseFloat, e)
	y := f.Then(fiveDivideBy, x)
	f.HandleError(func(e error) {
		log.Println(e)
	}, y)

	if y.Error == nil {
		t.Error(y)
	}

	os.Setenv("NUM", "5")

	e1 := f.Try(getEnv("NUM"))
	x1 := f.Then(parseFloat, e1)
	y1 := f.Then(fiveDivideBy, x1)

	if y1.Value != 1 {
		t.Error(y1)
	}

}
