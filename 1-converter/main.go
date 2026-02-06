package main

import (
	"fmt"
	"strings"
)

// Map с курсами валют относительно USD (теперь передаём по указателю)
var exchangeRates = map[string]float64{
	"USD": 1.0,
	"EUR": 0.85,
	"GBP": 0.73,
	"JPY": 110.5,
	"RUB": 75.5,
	"CNY": 6.45,
}

// Доступные валюты
var availableCurrencies = []string{"USD", "EUR", "GBP", "JPY", "RUB", "CNY"}

// Функция для вывода доступных валют
func printAvailableCurrencies() {
	fmt.Println("Доступные валюты:")
	for i, currency := range availableCurrencies {
		fmt.Printf("%d. %s\n", i+1, currency)
	}
	fmt.Println()
}

// Функция для проверки корректности валюты (теперь принимает указатель на map)
func isValidCurrency(currency string, rates *map[string]float64) bool {
	currency = strings.ToUpper(currency)
	_, exists := (*rates)[currency]
	return exists
}

// Функция для ввода валюты с проверкой (теперь с указателем)
func inputCurrency(prompt string, rates *map[string]float64) string {
	var currency string

	for {
		fmt.Print(prompt)
		fmt.Scan(&currency)
		currency = strings.ToUpper(currency)

		if isValidCurrency(currency, rates) {
			return currency
		}

		fmt.Println("Ошибка: неподдерживаемая валюта")
		printAvailableCurrencies()
		fmt.Print("Пожалуйста, введите валюту правильно: ")
	}
}

// Функция для расчета конвертации с использованием указателя на map
func convertCurrency(amount float64, fromCurrency, toCurrency string, rates *map[string]float64) (float64, error) {
	// Если валюты одинаковые, возвращаем ту же сумму
	if fromCurrency == toCurrency {
		return amount, nil
	}

	// Получаем курсы валют через указатель
	fromRate, fromExists := (*rates)[fromCurrency]
	toRate, toExists := (*rates)[toCurrency]

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

// Функция для получения курса конвертации между двумя валютами (с указателем)
func getExchangeRate(fromCurrency, toCurrency string, rates *map[string]float64) (float64, error) {
	if fromCurrency == toCurrency {
		return 1.0, nil
	}

	fromRate, fromExists := (*rates)[fromCurrency]
	toRate, toExists := (*rates)[toCurrency]

	if !fromExists || !toExists {
		return 0, fmt.Errorf("неподдерживаемая валюта")
	}

	return toRate / fromRate, nil
}

// Функция для вывода всех курсов (с указателем)
func printAllRates(rates *map[string]float64) {
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("Все доступные курсы конвертации:")

	// Преобразуем map в slice для стабильного порядка
	currencies := make([]string, 0, len(*rates))
	for currency := range *rates {
		currencies = append(currencies, currency)
	}

	// Сортируем для красивого вывода (упрощённая сортировка)
	for i := 0; i < len(currencies); i++ {
		for j := i + 1; j < len(currencies); j++ {
			if currencies[i] > currencies[j] {
				currencies[i], currencies[j] = currencies[j], currencies[i]
			}
		}
	}

	for _, fromCurr := range currencies {
		for _, toCurr := range currencies {
			if fromCurr != toCurr {
				rate, _ := getExchangeRate(fromCurr, toCurr, rates)
				fmt.Printf("1 %s = %.4f %s\n", fromCurr, rate, toCurr)
			}
		}
	}
}

func main() {
	fmt.Println("=== КОНВЕРТЕР ВАЛЮТ (оптимизированная версия) ===")
	fmt.Println()

	// Выводим курсы из map (передаём указатель)
	fmt.Println("Текущие курсы конвертации (относительно USD):")
	for currency, rate := range exchangeRates {
		if currency != "USD" {
			fmt.Printf("1 USD = %.2f %s\n", rate, currency)
		}
	}
	fmt.Printf("Всего валют в базе: %d\n", len(exchangeRates))

	// Рассчитываем и выводим кросс-курсы
	fmt.Println("\nКросс-курсы (через USD):")
	for fromCurr, fromRate := range exchangeRates {
		for toCurr, toRate := range exchangeRates {
			if fromCurr != toCurr {
				rate := toRate / fromRate
				fmt.Printf("1 %s = %.4f %s\n", fromCurr, rate, toCurr)
			}
		}
	}
	fmt.Println()

	// Шаг 1: Ввод исходной валюты (передаём указатель на map)
	fmt.Println("ШАГ 1: Ввод исходной валюты")
	printAvailableCurrencies()
	fromCurrency := inputCurrency("Введите исходную валюту: ", &exchangeRates)
	fmt.Println()

	// Шаг 2: Ввод суммы
	fmt.Println("ШАГ 2: Ввод суммы")
	var amount float64
	fmt.Print("Введите сумму для конвертации: ")
	fmt.Scan(&amount)
	fmt.Println()

	// Шаг 3: Ввод целевой валюты (передаём указатель на map)
	fmt.Println("ШАГ 3: Ввод целевой валюты")
	printAvailableCurrencies()
	toCurrency := inputCurrency("Введите целевую валюту: ", &exchangeRates)
	fmt.Println()

	// Шаг 4: Конвертация и вывод результата (передаём указатель на map)
	fmt.Println("ШАГ 4: Результат конвертации")

	// Вычисляем результат с использованием указателя на map
	result, err := convertCurrency(amount, fromCurrency, toCurrency, &exchangeRates)

	// Выводим результат
	if err == nil {
		fmt.Printf("\n✅ Конвертация успешно завершена!\n")
		fmt.Printf("%.2f %s = %.2f %s\n", amount, fromCurrency, result, toCurrency)

		// Получаем и выводим использованный курс
		rate, _ := getExchangeRate(fromCurrency, toCurrency, &exchangeRates)
		fmt.Printf("\nИспользованный курс: 1 %s = %.4f %s\n", fromCurrency, rate, toCurrency)

		// Выводим обратный курс
		reverseRate, _ := getExchangeRate(toCurrency, fromCurrency, &exchangeRates)
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
		{100.0, "EUR", "GBP"},
		{10000.0, "JPY", "USD"},
		{50.0, "GBP", "RUB"},
		{1000.0, "CNY", "EUR"},
	}

	for _, example := range examples {
		exampleResult, err := convertCurrency(example.amount, example.fromCurrency, example.toCurrency, &exchangeRates)
		if err == nil {
			rate, _ := getExchangeRate(example.fromCurrency, example.toCurrency, &exchangeRates)
			fmt.Printf("%.2f %s = %.2f %s (курс: 1 %s = %.4f %s)\n",
				example.amount, example.fromCurrency,
				exampleResult, example.toCurrency,
				example.fromCurrency, rate, example.toCurrency)
		}
	}

	// Показываем все доступные курсы (передаём указатель)
	printAllRates(&exchangeRates)

	// Демонстрация производительности с большим количеством валют
	fmt.Println("\n" + strings.Repeat("=", 40))
	fmt.Println("Демонстрация работы с большим количеством валют:")

	// Создаём большую map для демонстрации
	bigExchangeRates := make(map[string]float64)
	// Копируем существующие курсы
	for k, v := range exchangeRates {
		bigExchangeRates[k] = v
	}
	// Добавляем много новых валют (для демонстрации)
	for i := 0; i < 100; i++ {
		currencyCode := fmt.Sprintf("CR%d", i)
		bigExchangeRates[currencyCode] = 0.5 + float64(i)/100.0
	}

	fmt.Printf("Создана большая база курсов: %d валют\n", len(bigExchangeRates))

	// Конвертация с большой map (передаём указатель)
	bigResult, _ := convertCurrency(100.0, "USD", "EUR", &bigExchangeRates)
	fmt.Printf("Конвертация 100 USD в EUR в большой базе: %.2f EUR\n", bigResult)
}
