package main

import "fmt"

func main() {

	a := (42 == 42)
	b := (42 <= 42)
	c := (42 >= 43)
	d := (42 != 45)
	e := (42 < 43)
	f := (42 >= 41)

	fmt.Println(a, b, c, d, e, f)
}