package main

import "fmt"

func linearSearch(arr []int, target int) int {
	fmt.Printf("\n▶ Начинаем поиск %d в массиве: %v\n", target, arr)
	for i, value := range arr {
		fmt.Printf("→ Проверка индекса [%d] - значение %d", i, value)

		if value == target {
			fmt.Printf(" → Совпадение найдено!\n")
			return i
		}
		fmt.Printf(" → Не совпадает\n")
	}
	fmt.Println("× Элемент не найден")
	return -1
}

func main() {
	data := []int{24, 7, 15, 99, 43, 61, 3, 88}
	targets := []int{43, 5, 24, 88, 100}

	for _, target := range targets {
		fmt.Printf("\n======= Поиск значения %d =======\n", target)
		index := linearSearch(data, target)
		if index != -1 {
			fmt.Printf("\nРезультат: %d найден на позиции %d\n", target, index)
		} else {
			fmt.Printf("\nРезультат: %d отсутствует в массиве\n", target)
		}
	}
}
