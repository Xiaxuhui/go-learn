package try_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func returnMultiValue()(int, int) {
	return rand.Intn(10), rand.Intn(20)
}

func timeSpent2(inner func(op int) int) func(op int) int {
	return func(op int) int {
		start := time.Now()
		res := inner(op)
		fmt.Println("time spent:", time.Since(start))
		return res
	}
}

func slowFn(op int)(int){
	time.Sleep(1 * time.Second)
	return op
}

func sum(ops ...int) int {
	var res int
	for _, op := range ops {
		res += op
	}
	return res
}

func TestFn(t *testing.T) {
	a, _ := returnMultiValue()
	t.Log(a)
	fn :=  timeSpent(slowFn)
	res := fn(10)
	t.Log(res)
	t.Log(sum(1,2,3,4, 5))
}

func Clear() {
	fmt.Println("Clear")
	if err := recover(); err != nil {
		fmt.Println("recover", err)
	}
}

func TestDeffer(t *testing.T) {
	defer Clear()
	fmt.Println("Start")
	panic("err")
}