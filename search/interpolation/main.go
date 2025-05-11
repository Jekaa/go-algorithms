package main

import (
	"fmt"
)

func interpolationSearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right && target >= arr[left] && target <= arr[right] {
		// Избегаем деления на ноль при одинаковых значениях
		if arr[left] == arr[right] {
			if arr[left] == target {
				return left
			}
			return -1
		}

		// Интерполяционная формула для определения позиции
		mid := left + ((target-arr[left])*(right-left))/(arr[right]-arr[left])

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	uniformArray := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	target := 70

	resultIndex := interpolationSearch(uniformArray, target)

	if resultIndex != -1 {
		fmt.Printf("Элемент %d найден на позиции %d\n", target, resultIndex)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}

	// Демонстрация сравнения с бинарным поиском
	fmt.Println("\nСравнение эффективности:")
	tests := []struct {
		name          string
		arr           []int
		target        int
		expectedIndex int
		binarySteps   int
		interSteps    int
	}{
		{
			name:          "Равномерное распределение",
			arr:           uniformArray,
			target:        70,
			expectedIndex: 6,
			binarySteps:   3, // Бинарный: 4 → 7 → 6
			interSteps:    1, // Интерполяционный: сразу попадает в 6
		},
		{
			name:          "Неравномерное распределение",
			arr:           []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 100},
			target:        100,
			expectedIndex: 9,
			binarySteps:   4, // Бинарный: 4 → 7 → 8 → 9
			interSteps:    3, // Интерполяционный: 0 → 4 → 9
		},
		{
			name:          "Одинаковые элементы",
			arr:           []int{5, 5, 5, 5, 5},
			target:        5,
			expectedIndex: 0,
			binarySteps:   3, // Бинарный: 2 → 0
			interSteps:    1, // Интерполяционный: сразу 0
		},
		{
			name:          "Элемент отсутствует",
			arr:           uniformArray,
			target:        65,
			expectedIndex: -1,
			binarySteps:   3, // Бинарный: 4 → 7 → 5
			interSteps:    2, // Интерполяционный: 6 → 5
		},
	}

	for _, test := range tests {
		fmt.Printf("\nТест: %s\n", test.name)
		fmt.Printf("Ожидаемый результат: %d\n", test.expectedIndex)

		// Подсчет шагов для бинарного поиска
		binarySteps := 0
		binaryFunc := func(arr []int, target int) int {
			left := 0
			right := len(arr) - 1
			for left <= right {
				binarySteps++
				mid := left + (right-left)/2
				if arr[mid] == target {
					return mid
				} else if arr[mid] < target {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
			return -1
		}
		binaryFunc(test.arr, test.target)
		fmt.Printf("Бинарный поиск: %d шаг(ов) | Ожидание: %d\n", binarySteps, test.binarySteps)

		// Подсчет шагов для интерполяционного поиска
		interSteps := 0
		interFunc := func(arr []int, target int) int {
			left := 0
			right := len(arr) - 1
			for left <= right && target >= arr[left] && target <= arr[right] {
				interSteps++
				if arr[left] == arr[right] {
					if arr[left] == target {
						return left
					}
					return -1
				}
				mid := left + ((target-arr[left])*(right-left))/(arr[right]-arr[left])
				if arr[mid] == target {
					return mid
				} else if arr[mid] < target {
					left = mid + 1
				} else {
					right = mid - 1
				}
			}
			return -1
		}
		interFunc(test.arr, test.target)
		fmt.Printf("Интерполяционный поиск: %d шаг(ов) | Ожидание: %d\n", interSteps, test.interSteps)
	}
}
