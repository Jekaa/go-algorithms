package main

import (
	"container/list"
	"fmt"
)

// Graph представляет граф как список смежности
type Graph struct {
	adjacencyList map[string][]string
}

// BFS выполняет обход графа в ширину, начиная с указанного узла
func (g *Graph) BFS(start string) []string {
	var visited []string
	visitedMap := make(map[string]bool)
	queue := list.New()

	// Добавляем начальный узел в очередь
	queue.PushBack(start)
	visitedMap[start] = true

	// Основной цикл BFS
	for queue.Len() > 0 {
		element := queue.Front()
		current := queue.Remove(element).(string)
		visited = append(visited, current)

		// Обрабатываем всех соседей текущего узла
		for _, neighbor := range g.adjacencyList[current] {
			if !visitedMap[neighbor] {
				visitedMap[neighbor] = true
				queue.PushBack(neighbor)
			}
		}
	}

	return visited
}

func main() {
	graph := Graph{
		adjacencyList: map[string][]string{
			"A": {"B", "C"},
			"B": {"A", "D"},
			"C": {"A", "E"},
			"D": {"B"},
			"E": {"C"},
		},
	}

	result := graph.BFS("A")
	fmt.Println("Порядок посещения узлов:", result)
}
