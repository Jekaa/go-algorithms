package main

import (
	"fmt"
	"math"
)

// jumpSearch выполняет поиск элемента в отсортированном массиве методом прыжков.
// Возвращает индекс первого найденного элемента, равного target.
// Если элемент не найден, возвращает -1.
func jumpSearch(arr []int, target int) int {
	n := len(arr)
	if n == 0 {
		return -1
	}

	// Размер шага прыжка (округляется вниз до ближайшего целого)
	m := int(math.Sqrt(float64(n)))

	// Поиск подходящего блока
	prev := 0
	for arr[Min(m, n)-1] < target {
		prev = m
		m += int(math.Sqrt(float64(n)))
		if prev >= n {
			return -1
		}
	}

	// Линейный поиск в найденном блоке
	for i := prev; i < Min(m, n); i++ {
		if arr[i] == target {
			return i
		}
	}

	return -1
}

// Min возвращает минимальное из двух чисел
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Пример использования
	sortedArray := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 11

	resultIndex := jumpSearch(sortedArray, target)

	if resultIndex != -1 {
		fmt.Printf("Элемент %d найден на позиции %d\n", target, resultIndex)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}

	// Демонстрация граничных случаев
	fmt.Println("\nТестирование граничных случаев:")
	tests := []struct {
		name     string
		arr      []int
		target   int
		expected int
	}{
		{
			name:     "Пустой массив",
			arr:      []int{},
			target:   5,
			expected: -1,
		},
		{
			name:     "Один элемент (найден)",
			arr:      []int{42},
			target:   42,
			expected: 0,
		},
		{
			name:     "Один элемент (не найден)",
			arr:      []int{42},
			target:   7,
			expected: -1,
		},
		{
			name:     "Дубликаты элементов",
			arr:      []int{2, 4, 4, 4, 6, 8},
			target:   4,
			expected: 1,
		},
		{
			name:     "Элемент в начале",
			arr:      sortedArray,
			target:   sortedArray[0],
			expected: 0,
		},
		{
			name:     "Элемент в конце",
			arr:      sortedArray,
			target:   sortedArray[len(sortedArray)-1],
			expected: len(sortedArray) - 1,
		},
		{
			name:     "Элемент отсутствует",
			arr:      sortedArray,
			target:   12,
			expected: -1,
		},
	}

	for _, test := range tests {
		index := jumpSearch(test.arr, test.target)
		fmt.Printf("%s: %v\n", test.name, test.expected == index)
	}

	// Демонстрация эффективности
	fmt.Println("\nСравнение эффективности:")
	comparisonTests := []struct {
		name        string
		arr         []int
		target      int
		jumpSteps   int
		linearSteps int
		binarySteps int
	}{
		{
			name:        "Большой массив (1000 элементов)",
			arr:         generateSortedArray(1000),
			target:      887,
			jumpSteps:   43,  // Приблизительно sqrt(1000) = 32 шагов + 11 в линейном поиске
			linearSteps: 888, // Средний случай для линейного поиска
			binarySteps: 10,  // log2(1000) ≈ 10 шагов
		},
		{
			name:        "Маленький массив",
			arr:         []int{1, 3, 5, 7, 9, 11, 13, 15},
			target:      11,
			jumpSteps:   4, // sqrt(8) = 2.8 → 2 шага + 2 в линейном поиске
			linearSteps: 6, // 6 сравнений
			binarySteps: 3, // log2(8) = 3
		},
	}

	for _, test := range comparisonTests {
		fmt.Printf("\nТест: %s\n", test.name)

		// Подсчет шагов для Jump Search
		jumpSteps := 0
		jumpFunc := func(arr []int, target int) int {
			n := len(arr)
			if n == 0 {
				return -1
			}

			m := int(math.Sqrt(float64(n)))
			prev := 0

			for arr[Min(m, n)-1] < target {
				jumpSteps++
				prev = m
				m += int(math.Sqrt(float64(n)))
				if prev >= n {
					return -1
				}
			}

			for i := prev; i < Min(m, n); i++ {
				jumpSteps++
				if arr[i] == target {
					return i
				}
			}

			return -1
		}
		jumpFunc(test.arr, test.target)
		fmt.Printf("Jump Search: %d шаг(ов) | Ожидание: %d\n", jumpSteps, test.jumpSteps)

		// Подсчет шагов для Linear Search
		linearSteps := 0
		linearFunc := func(arr []int, target int) int {
			for i, val := range arr {
				linearSteps++
				if val == target {
					return i
				}
			}
			return -1
		}
		linearFunc(test.arr, test.target)
		fmt.Printf("Linear Search: %d шаг(ов) | Ожидание: %d\n", linearSteps, test.linearSteps)

		// Подсчет шагов для Binary Search
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
		fmt.Printf("Binary Search: %d шаг(ов) | Ожидание: %d\n", binarySteps, test.binarySteps)
	}
}

// generateSortedArray создает отсортированный массив заданной длины
func generateSortedArray(size int) []int {
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = i * 2 // Создаем массив с четными числами для наглядности
	}
	return arr
}
