package main

import (
	"fmt"
	"strings"
)

// Map с курсами валют относительно USD
var exchangeRates = map[string]float64{
	"USD": 1.0,
	"EUR": 0.85,
	"RUB": 75.5,
}

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
	_, exists := exchangeRates[currency]
	return exists
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

// Функция для расчета конвертации с использованием map
func convertCurrency(amount float64, fromCurrency, toCurrency string) (float64, error) {
	// Если валюты одинаковые, возвращаем ту же сумму
	if fromCurrency == toCurrency {
		return amount, nil
	}

	// Получаем курсы валют
	fromRate, fromExists := exchangeRates[fromCurrency]
	toRate, toExists := exchangeRates[toCurrency]

	if !fromExists {
		return 0, fmt.Errorf("неподдерживаемая валюта: %s", fromCurrency)
	}

	if !toExists {
		return 0, fmt.Errorf("неподдерживаемая валюта: %s", toCurrency)
	}

	// Конвертируем через USD как базовую валюту
	amountInUSD := amount / fromRate
	result := amountInUSD * toRate

	return result, nil
}

// Функция для получения курса конвертации между двумя валютами
func getExchangeRate(fromCurrency, toCurrency string) (float64, error) {
	if fromCurrency == toCurrency {
		return 1.0, nil
	}

	fromRate, fromExists := exchangeRates[fromCurrency]
	toRate, toExists := exchangeRates[toCurrency]

	if !fromExists || !toExists {
		return 0, fmt.Errorf("неподдерживаемая валюта")
	}

	return toRate / fromRate, nil
}

func main() {
	fmt.Println("=== КОНВЕРТЕР ВАЛЮТ ===")
	fmt.Println()

	// Выводим курсы из map
	fmt.Println("Текущие курсы конвертации (относительно USD):")
	for currency, rate := range exchangeRates {
		if currency != "USD" {
			fmt.Printf("1 USD = %.2f %s\n", rate, currency)
		}
	}

	// Рассчитываем и выводим кросс-курсы
	fmt.Println("\nКросс-курсы:")
	currencies := []string{"USD", "EUR", "RUB"}
	for i := 0; i < len(currencies); i++ {
		for j := 0; j < len(currencies); j++ {
			if i != j {
				rate, _ := getExchangeRate(currencies[i], currencies[j])
				fmt.Printf("1 %s = %.4f %s\n", currencies[i], rate, currencies[j])
			}
		}
	}
	fmt.Println()

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

	// Вычисляем результат с использованием map
	result, err := convertCurrency(amount, fromCurrency, toCurrency)

	// Выводим результат
	if err == nil {
		fmt.Printf("\n✅ Конвертация успешно завершена!\n")
		fmt.Printf("%.2f %s = %.2f %s\n", amount, fromCurrency, result, toCurrency)

		// Получаем и выводим использованный курс
		rate, _ := getExchangeRate(fromCurrency, toCurrency)
		fmt.Printf("\nИспользованный курс: 1 %s = %.4f %s\n", fromCurrency, rate, toCurrency)

		// Выводим обратный курс
		reverseRate, _ := getExchangeRate(toCurrency, fromCurrency)
		fmt.Printf("Обратный курс: 1 %s = %.4f %s\n", toCurrency, reverseRate, fromCurrency)
	} else {
		fmt.Printf("Ошибка при конвертации: %v\n", err)
	}

	// Примеры конвертации для демонстрации
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("Примеры других конвертаций:")

	// Примеры из map
	examples := []struct {
		amount       float64
		fromCurrency string
		toCurrency   string
	}{
		{100.0, "USD", "EUR"},
		{100.0, "EUR", "RUB"},
		{1000.0, "RUB", "USD"},
		{50.0, "EUR", "USD"},
		{5000.0, "RUB", "EUR"},
	}

	for _, example := range examples {
		exampleResult, err := convertCurrency(example.amount, example.fromCurrency, example.toCurrency)
		if err == nil {
			rate, _ := getExchangeRate(example.fromCurrency, example.toCurrency)
			fmt.Printf("%.2f %s = %.2f %s (курс: 1 %s = %.4f %s)\n",
				example.amount, example.fromCurrency,
				exampleResult, example.toCurrency,
				example.fromCurrency, rate, example.toCurrency)
		}
	}

	// Показываем все доступные курсы
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("Все доступные курсы конвертации:")

	for _, fromCurr := range availableCurrencies {
		for _, toCurr := range availableCurrencies {
			if fromCurr != toCurr {
				rate, _ := getExchangeRate(fromCurr, toCurr)
				fmt.Printf("1 %s = %.4f %s\n", fromCurr, rate, toCurr)
			}
		}
	}
}
