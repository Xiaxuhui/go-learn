package try_test

import (
	"errors"
	"testing"
)


func doSomething(n int)(int, error) {
	if n < 2 {
		return n, errors.New("n should be note less than 2")
	}
	if n > 100 {
		return n, errors.New("n should be larger than 100")
	}
	return n, nil
}

func TestError(t *testing.T) {
	if _, err := doSomething(1) ; err != nil {
		t.Error(err)
	}
}