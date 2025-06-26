package main

import "fmt"

func main() {
	x := 7897
	switch {
	case x < 5:
		fmt.Println("x<5")
		fmt.Println("multistr test")
		x++
		fmt.Print("x=", x)
	case x < 10:
		fmt.Print("x<10")
	case x < 15:
		fmt.Print("x<15")
	default:
		fmt.Print("default")
	}

}
