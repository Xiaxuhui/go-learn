package try_test

import "testing"

const (
	ReadAble = 1 << iota
	WriteAble
	ExecuteAble
)

func TestCompareArray(t * testing.T) {
	// a := [...]int{1, 2, 3, 4 }
	// b := [...]int{1, 2, 3, 4 }
	// c := [...]int{1, 2, 3, 5 }
	d := 7;
	d = d &^ WriteAble
	t.Log(d)
	t.Log(ReadAble)
	t.Log(WriteAble)
	t.Log(ExecuteAble)
}