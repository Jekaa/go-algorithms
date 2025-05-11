package main

import (
	"fmt"
)

func binarySearch(arr []int, target int) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
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

func main() {
	sortedArray := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	target := 11

	resultIndex := binarySearch(sortedArray, target)

	if resultIndex != -1 {
		fmt.Printf("Элемент %d найден на позиции %d\n", target, resultIndex)
	} else {
		fmt.Printf("Элемент %d не найден в массиве\n", target)
	}

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
			expected: 2,
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
	}

	for _, test := range tests {
		index := binarySearch(test.arr, test.target)
		fmt.Printf("%s: %v\n", test.name, test.expected == index)
	}
}
