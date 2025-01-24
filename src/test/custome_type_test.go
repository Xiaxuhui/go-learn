package try_test

import (
	"fmt"
	"time"
)

type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
	return func(op int) int {
		start := time.Now()
		res := inner(op)
		fmt.Println("time spent:", time.Since(start))
		return res
	}
}
