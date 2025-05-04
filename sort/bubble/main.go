package main

import (
	"fmt"
)

// bubbleSort сортирует срез целых чисел по возрастанию
func bubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
		// Если ни одного обмена не произошло, массив уже отсортирован
		if !swapped {
			break
		}
	}
}

func main() {
	arr := []int{5, 3, 8, 6, 7, 2}
	fmt.Println("Исходный массив:", arr)
	bubbleSort(arr)
	fmt.Println("Отсортированный массив:", arr)
}
