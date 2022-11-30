package functional

import (
	"errors"
	"fmt"
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
	y := f.Then(fiveDivideBy, x).
		Catch(func(e error) {
			log.Println(e)
		})

	if y.Error == nil {
		t.Error(y.Value)
	}

	os.Setenv("NUM", "5")

	e1 := f.Try(getEnv("NUM"))
	x1 := f.Then(parseFloat, e1)
	y1 := f.Then(fiveDivideBy, x1)

	if y1.Value != 1 {
		t.Error(y1)
	}

}

func TestTryCatch(t *testing.T) {
	y := f.
		Try(func(name string) (float64, error) {
			e := os.Getenv(name)
			if e == "" {
				return 0, errors.New(fmt.Sprintf("env %s is not set", e))
			}
			x, err := strconv.ParseFloat(e, 64)
			if err != nil {
				return 0, err
			}
			if x == 0 {
				return 0, errors.New("cannot divide by 0")
			}
			return 5 / x, nil
		}("NUM")).
		Catch(func(e error) {
			log.Println(e)
		})

	t.Error(y)
}
