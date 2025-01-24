package try_test

import (
	"fmt"
	"testing"
)

func TestFibList(t *testing.T) {
	var (
		a = 1
		b = 1
	)
	fmt.Println(a)
	for i :=0; i<10; i++ {  
		fmt.Println(" ", b);
		temp := a;
		a = b;
		b = temp + b;
	}
}