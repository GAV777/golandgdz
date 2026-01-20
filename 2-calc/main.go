package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Проверка количества аргументов
	if len(os.Args) < 3 {
		printUsage()
		return
	}

	// Получаем операцию и строку с числами
	operation := strings.ToUpper(os.Args[1])
	numbersStr := os.Args[2]

	// Разбиваем строку чисел по запятым
	numbers, err := parseNumbers(numbersStr)
	if err != nil {
		fmt.Printf("Ошибка при разборе чисел: %v\n", err)
		return
	}

	// Проверяем, что есть хотя бы одно число
	if len(numbers) == 0 {
		fmt.Println("Ошибка: не указаны числа для расчета")
		return
	}

	// Выполняем операцию
	result, err := calculate(operation, numbers)
	if err != nil {
		fmt.Printf("Ошибка при расчете: %v\n", err)
		return
	}

	// Выводим результат
	fmt.Printf("Результат операции %s: %.2f\n", operation, result)
}

// parseNumbers разбивает строку чисел по запятым и преобразует в числа
func parseNumbers(numbersStr string) ([]float64, error) {
	// Удаляем пробелы и разбиваем по запятым
	strNumbers := strings.Split(strings.ReplaceAll(numbersStr, " ", ""), ",")

	numbers := make([]float64, 0, len(strNumbers))

	for _, strNum := range strNumbers {
		if strNum == "" {
			continue // Пропускаем пустые значения
		}

		num, err := strconv.ParseFloat(strNum, 64)
		if err != nil {
			return nil, fmt.Errorf("неверное число: '%s'", strNum)
		}

		numbers = append(numbers, num)
	}

	return numbers, nil
}

// calculate выполняет указанную операцию над числами
func calculate(operation string, numbers []float64) (float64, error) {
	switch operation {
	case "SUM":
		return sum(numbers), nil
	case "AVG":
		return avg(numbers), nil
	case "MED":
		return median(numbers), nil
	default:
		return 0, fmt.Errorf("неизвестная операция: %s. Доступные операции: SUM, AVG, MED", operation)
	}
}

// sum вычисляет сумму чисел
func sum(numbers []float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}
	return total
}

// avg вычисляет среднее значение
func avg(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}
	return sum(numbers) / float64(len(numbers))
}

// median вычисляет медиану
func median(numbers []float64) float64 {
	if len(numbers) == 0 {
		return 0
	}

	// Создаем копию для сортировки (чтобы не менять исходный слайс)
	sorted := make([]float64, len(numbers))
	copy(sorted, numbers)
	sort.Float64s(sorted)

	mid := len(sorted) / 2

	if len(sorted)%2 == 1 {
		// Нечетное количество элементов
		return sorted[mid]
	} else {
		// Четное количество элементов
		return (sorted[mid-1] + sorted[mid]) / 2
	}
}

// printUsage выводит информацию об использовании программы
func printUsage() {
	fmt.Println("Использование:")
	fmt.Println("  calculator <операция> <числа через запятую>")
	fmt.Println()
	fmt.Println("Операции:")
	fmt.Println("  SUM - вычисляет сумму чисел")
	fmt.Println("  AVG - вычисляет среднее значение")
	fmt.Println("  MED - вычисляет медиану")
	fmt.Println()
	fmt.Println("Примеры:")
	fmt.Println("  calculator SUM 2,10,9")
	fmt.Println("  calculator AVG 1,2,3,4,5")
	fmt.Println("  calculator MED 7,3,5,1,9")
}
