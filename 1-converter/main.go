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

// Доступные валюты
var availableCurrencies = []string{"USD", "EUR", "RUB"}

// Функция для вывода доступных валют
func printAvailableCurrencies() {
	fmt.Println("Доступные валюты:")
	for i, currency := range availableCurrencies {
		fmt.Printf("%d. %s\n", i+1, currency)
	}
	fmt.Println()
}

// Функция для проверки корректности валюты
func isValidCurrency(currency string) bool {
	currency = strings.ToUpper(currency)
	for _, validCurrency := range availableCurrencies {
		if currency == validCurrency {
			return true
		}
	}
	return false
}

// Функция для ввода валюты с проверкой
func inputCurrency(prompt string) string {
	var currency string

	for {
		fmt.Print(prompt)
		fmt.Scan(&currency)
		currency = strings.ToUpper(currency)

		if isValidCurrency(currency) {
			return currency
		}

		fmt.Println("Ошибка: неподдерживаемая валюта")
		printAvailableCurrencies()
		fmt.Print("Пожалуйста, введите валюту правильно: ")
	}
}

// Функция для ввода числа с проверкой
func inputAmount(prompt string) float64 {
	var amount float64

	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&amount)

		if err != nil || amount <= 0 {
			fmt.Println("Ошибка: введите положительное число")
			// Очищаем буфер ввода
			var discard string
			fmt.Scanln(&discard)
		} else {
			return amount
		}
	}
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
	fmt.Println("=== КОНВЕРТЕР ВАЛЮТ ===")
	fmt.Println()

	// Выводим курсы
	fmt.Println("Текущие курсы конвертации:")
	fmt.Printf("1 USD = %.2f EUR\n", usdToEur)
	fmt.Printf("1 USD = %.2f RUB\n", usdToRub)
	fmt.Printf("1 EUR = %.2f RUB\n\n", eurToRub)

	// Шаг 1: Ввод исходной валюты
	fmt.Println("ШАГ 1: Ввод исходной валюты")
	printAvailableCurrencies()
	fromCurrency := inputCurrency("Введите исходную валюту: ")
	fmt.Println()

	// Шаг 2: Ввод суммы
	fmt.Println("ШАГ 2: Ввод суммы")
	amount := inputAmount("Введите сумму для конвертации: ")
	fmt.Println()

	// Шаг 3: Ввод целевой валюты
	fmt.Println("ШАГ 3: Ввод целевой валюты")
	printAvailableCurrencies()
	toCurrency := inputCurrency("Введите целевую валюту: ")
	fmt.Println()

	// Шаг 4: Конвертация и вывод результата
	fmt.Println("ШАГ 4: Результат конвертации")

	// Вычисляем результат с помощью if/switch
	var result float64
	var err error

	if fromCurrency == toCurrency {
		result = amount
	} else {
		// Конвертируем сначала в USD
		var amountInUSD float64

		switch fromCurrency {
		case "USD":
			amountInUSD = amount
		case "EUR":
			amountInUSD = amount / usdToEur
		case "RUB":
			amountInUSD = amount / usdToRub
		}

		// Конвертируем из USD в целевую валюту
		switch toCurrency {
		case "USD":
			result = amountInUSD
		case "EUR":
			result = amountInUSD * usdToEur
		case "RUB":
			result = amountInUSD * usdToRub
		}
	}

	// Выводим результат
	if err == nil {
		fmt.Printf("\n✅ Конвертация успешно завершена!\n")
		fmt.Printf("%.2f %s = %.2f %s\n", amount, fromCurrency, result, toCurrency)

		// Дополнительная информация о курсе
		fmt.Println("\nИспользованные курсы:")
		if fromCurrency == "USD" && toCurrency == "EUR" {
			fmt.Printf("1 USD = %.2f EUR\n", usdToEur)
		} else if fromCurrency == "USD" && toCurrency == "RUB" {
			fmt.Printf("1 USD = %.2f RUB\n", usdToRub)
		} else if fromCurrency == "EUR" && toCurrency == "RUB" {
			fmt.Printf("1 EUR = %.2f RUB\n", eurToRub)
		} else if fromCurrency == "EUR" && toCurrency == "USD" {
			fmt.Printf("1 EUR = %.2f USD\n", 1/usdToEur)
		} else if fromCurrency == "RUB" && toCurrency == "USD" {
			fmt.Printf("1 RUB = %.4f USD\n", 1/usdToRub)
		} else if fromCurrency == "RUB" && toCurrency == "EUR" {
			fmt.Printf("1 RUB = %.4f EUR\n", 1/eurToRub)
		}
	} else {
		fmt.Printf("Ошибка при конвертации: %v\n", err)
	}

	// Примеры конвертации для демонстрации
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("Примеры других конвертаций:")

	// Пример 1: USD -> EUR
	exampleAmount := 100.0
	exampleResult, _ := convertCurrency(exampleAmount, "USD", "EUR")
	fmt.Printf("%.2f USD = %.2f EUR\n", exampleAmount, exampleResult)

	// Пример 2: EUR -> RUB
	exampleResult, _ = convertCurrency(exampleAmount, "EUR", "RUB")
	fmt.Printf("%.2f EUR = %.2f RUB\n", exampleAmount, exampleResult)

	// Пример 3: RUB -> USD
	exampleAmount = 1000.0
	exampleResult, _ = convertCurrency(exampleAmount, "RUB", "USD")
	fmt.Printf("%.2f RUB = %.2f USD\n", exampleAmount, exampleResult)
}
