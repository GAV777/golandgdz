package main

import (
	"fmt"
	"strings"
)

// Константы конвертации
const (
	usdToEur = 0.85 // 1 USD = 0.85 EUR
	usdToRub = 75.5 // 1 USD = 75.5 RUB
)

// Курс EUR в RUB через USD
const eurToRub = usdToRub / usdToEur

// Функция для считывания ввода пользователя
func getUserInput() (float64, string, string) {
	var amount float64
	var fromCurrency, toCurrency string

	fmt.Print("Введите сумму для конвертации: ")
	fmt.Scan(&amount)

	fmt.Print("Введите исходную валюту (USD, EUR, RUB): ")
	fmt.Scan(&fromCurrency)

	fmt.Print("Введите целевую валюту (USD, EUR, RUB): ")
	fmt.Scan(&toCurrency)

	return amount, strings.ToUpper(fromCurrency), strings.ToUpper(toCurrency)
}

// Функция для расчета конвертации
func convertCurrency(amount float64, fromCurrency, toCurrency string) (float64, error) {
	// Если валюты одинаковые, возвращаем ту же сумму
	if fromCurrency == toCurrency {
		return amount, nil
	}

	// Конвертируем сначала в USD как промежуточную валюту
	var amountInUSD float64

	// Конвертируем исходную валюту в USD
	switch fromCurrency {
	case "USD":
		amountInUSD = amount
	case "EUR":
		amountInUSD = amount / usdToEur
	case "RUB":
		amountInUSD = amount / usdToRub
	default:
		return 0, fmt.Errorf("неподдерживаемая валюта: %s", fromCurrency)
	}

	// Конвертируем из USD в целевую валюту
	var result float64
	switch toCurrency {
	case "USD":
		result = amountInUSD
	case "EUR":
		result = amountInUSD * usdToEur
	case "RUB":
		result = amountInUSD * usdToRub
	default:
		return 0, fmt.Errorf("неподдерживаемая валюта: %s", toCurrency)
	}

	return result, nil
}

func main() {
	// Выводим курсы
	fmt.Println("Курсы конвертации:")
	fmt.Printf("1 USD = %.2f EUR\n", usdToEur)
	fmt.Printf("1 USD = %.2f RUB\n", usdToRub)
	fmt.Printf("1 EUR = %.2f RUB\n\n", eurToRub)

	// Получаем ввод от пользователя
	amount, fromCurrency, toCurrency := getUserInput()

	// Выполняем конвертацию
	convertedAmount, err := convertCurrency(amount, fromCurrency, toCurrency)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	// Выводим результат
	fmt.Printf("\nРезультат конвертации:\n")
	fmt.Printf("%.2f %s = %.2f %s\n", amount, fromCurrency, convertedAmount, toCurrency)

	// Примеры конвертации для демонстрации
	fmt.Println("\nПримеры конвертации:")

	// USD -> EUR
	usdAmount := 100.0
	eurFromUsd, _ := convertCurrency(usdAmount, "USD", "EUR")
	fmt.Printf("%.2f USD = %.2f EUR\n", usdAmount, eurFromUsd)

	// EUR -> RUB
	eurAmount := 100.0
	rubFromEur, _ := convertCurrency(eurAmount, "EUR", "RUB")
	fmt.Printf("%.2f EUR = %.2f RUB\n", eurAmount, rubFromEur)

	// RUB -> USD
	rubAmount := 1000.0
	usdFromRub, _ := convertCurrency(rubAmount, "RUB", "USD")
	fmt.Printf("%.2f RUB = %.2f USD\n", rubAmount, usdFromRub)
}
