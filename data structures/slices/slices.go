package main

import (
	"fmt"
)

func main() {
	x := [] int{1, 2, 3, 4, 5}
	fmt.Println(len(x))
	for in, val := range x {
		fmt.Println(in, val)
	}
}