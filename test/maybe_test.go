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
	y := f.Then(fiveDivideBy, x).
		Catch(func(e error) {
			log.Println(e)
		})

	if y.Error == nil {
		t.Error(y.Value)
	}

	os.Setenv("NUM2", "5")

	e1 := f.Try(getEnv("NUM2"))
	x1 := f.Then(parseFloat, e1)
	y1 := f.Then(fiveDivideBy, x1)

	if y1.Value != 1 {
		t.Error(y1)
	}

}

func BenchmarkNoMaybe(b *testing.B) {
	os.Setenv("NUM3", "8")
	doAll := func(name string) (float64, error) {
		env, err := getEnv(name)
		if err != nil {
			return 0, err
		}
		x, err := parseFloat(env)
		if err != nil {
			return 0, err
		}
		y, err := fiveDivideBy(x)
		if err != nil {
			return 0, err
		}
		return y, err
	}

	for n := 0; n < b.N; n++ {
		_, err := doAll("NUM3")
		if err != nil {
			log.Println(err)
		}
	}

}

func BenchmarkMaybe(b *testing.B) {
	os.Setenv("NUM3", "8")
	for n := 0; n < b.N; n++ {
		env := f.Try(getEnv("NUM3"))
		x := f.Then(parseFloat, env)
		f.Then(fiveDivideBy, x).
			Catch(func(e error) { log.Println(e) })
	}
}
