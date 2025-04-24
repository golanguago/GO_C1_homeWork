package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("_____Подсчёт ИМТ_____")
	for {
		defer func() {
			r := recover()
			if r != nil {
				fmt.Println("ТЕКСТ РЕКАВЕРА if r != nil ФУНКЦИИ defer, где r = ", r)
			}

		}()

		IMT, err := calculate(getIMT())
		if err != nil {
			panic("Что то пошло не так, вероятно вы ошиблись при вводе роста или веса.")
			// fmt.Println(err)
			// continue
		}
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

func calculate(userH float64, userW float64) (float64, error) {
	if userH <= 0 || userW <= 0 {
		return 0, errors.New("WRONG_PARAM_HEIGHT_OR_WEIGHT_ERROR_TRY_AGAIN")
	}
	const IMTPower = 2
	result := userW / math.Pow(userH/100, IMTPower)
	return result, nil
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
