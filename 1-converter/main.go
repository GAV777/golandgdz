package main

import "fmt"

func main() {
	// Константы конвертации
	const usdToEur = 0.85 // 1 USD = 0.85 EUR
	const usdToRub = 75.5 // 1 USD = 75.5 RUB

	// Рассчитываем курс EUR в RUB через USD
	// Если 1 USD = 0.85 EUR и 1 USD = 75.5 RUB, то:
	// 1 EUR = 75.5 / 0.85 RUB
	const eurToRub = usdToRub / usdToEur

	// Выводим курсы
	fmt.Printf("Курсы конвертации:\n")
	fmt.Printf("1 USD = %.2f EUR\n", usdToEur)
	fmt.Printf("1 USD = %.2f RUB\n", usdToRub)
	fmt.Printf("1 EUR = %.2f RUB\n", eurToRub)

	// Пример конвертации
	usdAmount := 100.0
	eurAmount := usdAmount * usdToEur
	rubAmount := usdAmount * usdToRub

	fmt.Printf("\nПример конвертации:\n")
	fmt.Printf("%.2f USD = %.2f EUR\n", usdAmount, eurAmount)
	fmt.Printf("%.2f USD = %.2f RUB\n", usdAmount, rubAmount)
	fmt.Printf("%.2f EUR = %.2f RUB\n", eurAmount, eurAmount*eurToRub)
}
