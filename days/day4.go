package days

import (
	"bufio"
	"fmt"
	"os"
)

func readRowsDay4() ([]string, error) {
	file, err := os.Open("./dataFiles/day4.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	var graph []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		graph = append(graph, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return graph, nil
}

func (c *ChristmasSaver) FindAllInstances(word string) (int, error) {
	input, err := readRowsDay4()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0, err
	}

	var graph [][]rune
	for _, line := range input {
		graph = append(graph, []rune(line))
	}

	result := 0
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[0]); j++ {
			for _, direction := range getDirections() {
				if searchInDirection(graph, []rune(word), i, j, direction) {
					result++
				}
			}
		}
	}

	return result, nil
}

func searchInDirection(graph [][]rune, target []rune, x, y int, direction []int) bool {
	for i := 0; i < len(target); i++ {
		newX, newY := x+direction[0]*i, y+direction[1]*i

		if newX < 0 || newX >= len(graph) || newY < 0 || newY >= len(graph[0]) || graph[newX][newY] != target[i] {
			return false
		}
	}

	return true
}

func getDirections() [][]int {
	return [][]int{
		{-1, 0}, {1, 0},
		{0, -1}, {0, 1},
		{-1, -1}, {1, 1},
		{-1, 1}, {1, -1},
	}
}
