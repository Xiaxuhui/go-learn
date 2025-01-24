package main

import (
	"fmt"
	"go_learn/src/series"
	"os"
)

func main() {
	fmt.Println("hello world");
	if len(os.Args) > 1 {
		fmt.Println(os.Args[1]);
	}
	fmt.Println(series.GetSum(1,2))
	os.Exit(0)
}