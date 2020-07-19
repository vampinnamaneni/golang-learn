package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Don't print")
	case (2 ==4):
		fmt.Println("Not this too")
	case (3 == 4):
		fmt.Println("Print this now")
		fallthrough
	case (4 ==5):
		fmt.Println("This will print too")
		fallthrough
	case (5 == 6):
		fmt.Println("Again don't print this")
		fallthrough
	case (6 == 7):
		fmt.Println("Print this time again")
	default:
		fmt.Println("this is default")
	}
}