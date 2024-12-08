package days

import (
	"bufio"
	"fmt"
	"os"
)

func (c *ChristmasSaver) FindAllInstances(word string) (int, error) {
	input, err := readRowsDay4("dataFiles/day4.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0, err
	}

	return countWord(input, word), nil
}

func (c *ChristmasSaver) FindAllCrossInstances() (int, error) {
	input, err := readRowsDay4("dataFiles/day4.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0, err
	}

	return countCrossXMAS(input), nil
}

func readRowsDay4(filename string) ([][]rune, error) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return grid, nil
}

func countWord(grid [][]rune, word string) int {
	target := []rune(word)
	count := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			for _, dir := range directions {
				if searchWord(grid, target, i, j, dir) {
					count++
				}
			}
		}
	}
	return count
}

func searchWord(grid [][]rune, target []rune, x, y int, dir []int) bool {
	for k := 0; k < len(target); k++ {
		nx, ny := x+k*dir[0], y+k*dir[1]
		if nx < 0 || nx >= len(grid) || ny < 0 || ny >= len(grid[0]) || grid[nx][ny] != target[k] {
			return false
		}
	}
	return true
}

func countCrossXMAS(grid [][]rune) int {
	rows := len(grid)
	cols := len(grid[0])
	count := 0

	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ {
			if grid[r][c] == 'A' {
				if isValidXMAS(grid, r, c) {
					count++
				}
			}
		}
	}

	return count
}

func isValidXMAS(grid [][]rune, r, c int) bool {
	topLeft := grid[r-1][c-1]
	bottomRight := grid[r+1][c+1]
	topRight := grid[r-1][c+1]
	bottomLeft := grid[r+1][c-1]

	cond1 := topLeft == 'M' && bottomRight == 'S' || topLeft == 'S' && bottomRight == 'M'
	cond2 := topRight == 'M' && bottomLeft == 'S' || topRight == 'S' && bottomLeft == 'M'

	return cond1 && cond2
}

var directions = [][]int{
	{-1, 0}, {1, 0},
	{0, -1}, {0, 1},
	{-1, -1}, {1, 1},
	{-1, 1}, {1, -1},
}
