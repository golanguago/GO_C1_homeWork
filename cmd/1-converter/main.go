package main

import "fmt"

func main() {
	const USDtoEUR = 0.88
	const USDtoRUB = 84.0
	EURtoRUB := 1 / USDtoEUR * USDtoRUB
	fmt.Println("1 EUR to RUB = ", EURtoRUB)
}
