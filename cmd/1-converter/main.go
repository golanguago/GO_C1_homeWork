package main

import "fmt"

func main() {

	// //hw #1
	// const USDtoEUR = 0.88
	// const USDtoRUB = 84.0
	// EURtoRUB := 1 / USDtoEUR * USDtoRUB
	// fmt.Println("1 EUR to RUB = ", EURtoRUB)

	// my variant hw #2
	const convUSDtoEUR = 0.88
	const convUSDtoRUB = 84.0
	getUserInput := getUserInput()
	fmt.Printf("Готово! %.2f EUR = %.2f RUB.", getUserInput, calculate(getUserInput, convUSDtoEUR, convUSDtoRUB))
}

func getUserInput() float64 {
	fmt.Println("Введите сумму EUR для конвертации в RUB: ")
	var EURtoRUB float64
	fmt.Scan(&EURtoRUB)
	fmt.Println("Принято, считаем...")
	return EURtoRUB
}

func calculate(EURtoRUB float64, USDtoEUR float64, USDtoRUB float64) float64 {
	convEURtoRUB := EURtoRUB / USDtoEUR * USDtoRUB
	return convEURtoRUB
}
