package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("_____Подсчёт ИМТ_____")
	for {
		IMT := calculate(getIMT())
		fmt.Println(printResult(IMT))
		userChoice := repeatOrNotCalc()
		if !userChoice {
			fmt.Print("Goodbye!")
			break
		}

	}
}

// func outputResult(imt float64) {
// 	fmt.Println("Ваш индекс массы тела: ")
// 	fmt.Printf("%.1f", imt)
// }

func getIMT() (float64, float64) {

	var userHeight float64
	var userWeight float64
	fmt.Println("Введите свой рост в сантиметрах: ")
	fmt.Scan(&userHeight)
	fmt.Print("Введите свой вес в килограммах: ")
	fmt.Scan(&userWeight)
	return userHeight, userWeight
}

func calculate(userH float64, userW float64) float64 {
	const IMTPower = 2
	result := userW / math.Pow(userH/100, IMTPower)
	return result
}

func printResult(IMT float64) string {
	if IMT < 16 {
		return "Очень маленький вес"
	} else {
		return "Не самый маленький вес"
	}
}

func repeatOrNotCalc() bool {
	var choice string
	fmt.Print("Do you want try again? (y/n): ")
	fmt.Scan(&choice)
	if choice == "y" || choice == "Y" {
		return true
	}
	return false
}
