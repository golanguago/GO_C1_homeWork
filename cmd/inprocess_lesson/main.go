package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	fmt.Println("_____Подсчёт_ИМТ_____")
	for {
		defer func() {
			r := recover()
			if r != nil {
				fmt.Println("defer -> r != nil, где r (recover, он же переданный текст в panic) = ", r)
			}

		}()

		IMT, err := calculate(getIMT())
		if err != nil {
			panic("Что то пошло не так, скорее всего вы ошиблись при вводе роста и/или веса.")
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
		return 0, errors.New("ERROR_WRONG_USER_PARAM_HEIGHT_OR_WEIGHT")
	}
	const IMTPower = 2
	result := userW / math.Pow(userH/100, IMTPower)
	return result, nil
}

func printResult(imt float64) string {
	if imt < 16 {
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
