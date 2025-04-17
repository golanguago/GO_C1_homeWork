package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("_____Подсчёт ИМТ_____")
	IMT := calculate(getIMT())
	outputResult(IMT)
}

func outputResult(imt float64) {
	fmt.Println("Ваш индекс массы тела: ")
	fmt.Printf("%.1f", imt)
}

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
