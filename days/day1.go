package days

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func readRowsDay1() ([]int, []int, error) {
	file, err := os.Open("./dataFiles/day1.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, nil, err
	}

	defer file.Close()
	var col1, col2 []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := []int{}
		for _, part := range strings.Fields(line) {
			num, err := strconv.Atoi(part)
			if err == nil {
				numbers = append(numbers, num)
			}
		}

		if len(numbers) >= 2 {
			col1 = append(col1, numbers[0])
			col2 = append(col2, numbers[1])
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	sort.Ints(col1)
	sort.Ints(col2)

	return col1, col2, nil
}

func (c *ChristmasSaver) GetListDistance() (int, error) {
	col1, col2, err := readRowsDay1()
	if err != nil {
		return 0, err
	}

	distance := 0
	for i := range col1 {
		diff := col1[i] - col2[i]
		if diff < 0 {
			diff = -diff
		}
		distance += diff
	}

	return distance, nil
}

func (c *ChristmasSaver) GetSimilarityScore() (int, error) {
	col1, col2, err := readRowsDay1()
	if err != nil {
		return 0, err
	}

	rightCount := make(map[int]int)
	for _, num := range col2 {
		rightCount[num]++
	}

	score := 0
	for _, num := range col1 {
		if value, exists := rightCount[num]; exists {
			score += num * value
		}
	}

	return score, nil
}
