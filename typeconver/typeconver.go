package main

import "fmt"

var a int
type myown int
var b myown
func main() {
a = 10
b = 12

fmt.Println(a)
fmt.Printf("%T\n", a)
fmt.Println(b)
fmt.Printf("%T\n", b)

a = int(b)

fmt.Println(a)

}
