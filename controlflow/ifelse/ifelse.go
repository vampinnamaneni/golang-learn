package main

import "fmt"

func main() {
	x := 57
	if x > 3 && x < 57 {
		fmt.Println("Right range")
	} else if x == 57{
		fmt.Println("Vamsi")
		} else {
		fmt.Println("Invalid roll number")
	}
}