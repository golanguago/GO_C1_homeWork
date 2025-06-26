package main

import (
	"fmt"
)

func main() {
	fmt.Println("Добро пожаловать в конвертор валюты!")
	sourceCurrency, sum, exitCurrency := getСurrency()
	result := calculateResult(sourceCurrency, sum, exitCurrency)
	fmt.Printf("Итого: %.2f %s = %.2f %s", sum, sourceCurrency, result, exitCurrency)
}

func getСurrency() (string, float64, string) {
	var sourceCurrency, exitCurrency string
	var sum float64
	for {
		fmt.Print("Пожалуйста, введите конвертируемую валюту (EUR, RUB или USD): ")
		fmt.Scan(&sourceCurrency)
		if sourceCurrency == "EUR" || sourceCurrency == "RUB" || sourceCurrency == "USD" {
			break
		}
		fmt.Println("\nТакой валюты не предусмотрено или введено некорректно, пожалуйста повторите ввод согласно образцу соблюдая регистр!")
	}

	for {
		fmt.Printf("Сколько %s нужно конвертировать: ", sourceCurrency)
		fmt.Scan(&sum)
		if sum <= 0 {
			fmt.Println("\nВы ошиблись при вводе числа, попробуйте ещё раз!") //других способов проверки корректного ввода числа с плавающей точкой в уроках не было
			continue
		}
		break
	}

	var choiceCurrency string
	switch sourceCurrency {
	case "EUR":
		choiceCurrency = "(RUB или USD)"
	case "RUB":
		choiceCurrency = "(EUR или USD)"
	case "USD":
		choiceCurrency = "(EUR или RUB)"
	default:
		choiceCurrency = "(default)"

	}

	for {
		fmt.Printf("В какую валюту будем пересчитывать? %s: ", choiceCurrency)
		fmt.Scan(&exitCurrency)
		if exitCurrency == "EUR" || exitCurrency == "RUB" || exitCurrency == "USD" && choiceCurrency != exitCurrency {
			break
		} else if choiceCurrency != exitCurrency {
			fmt.Println("Из этой валюты конвертируем! Выберите из предложенных!")
			continue
		}
		fmt.Println("\nТакой валюты не предусмотрено или введено некорректно, пожалуйста повторите ввод согласно образцу соблюдая регистр!")
	}

	return sourceCurrency, sum, exitCurrency
}

func calculateResult(sourceCurrency string, sum float64, exitCurrency string) float64 {
	switch {
	//якобы варианты верных расчетов после ввода данных
	case sourceCurrency == "EUR" && exitCurrency == "RUB":
		return sum * 2
	case sourceCurrency == "EUR" && exitCurrency == "USD":
		return sum * 3
	case sourceCurrency == "RUB" && exitCurrency == "EUR":
		return sum * 4
	case sourceCurrency == "RUB" && exitCurrency == "USD":
		return sum * 5
	case sourceCurrency == "USD" && exitCurrency == "EUR":
		return sum * 6
	case sourceCurrency == "USD" && exitCurrency == "RUB":
		return sum * 7
	default:
		return 0
	}
}

// 1 EUR to RUB = 1 / USDtoEUR * USDtoRUB
//
// 	const convUSDtoEUR = 0.88
// 	const convUSDtoRUB = 84.0
//	func calculate(EURtoRUB float64, USDtoEUR float64, USDtoRUB float64) float64 {
//		convEURtoRUB := EURtoRUB / USDtoEUR * USDtoRUB
//		return convEURtoRUB
