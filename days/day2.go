package days

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readRowsDay2() ([][]int, error) {
	file, err := os.Open("./dataFiles/day2.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	defer file.Close()
	var rows [][]int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		numbers := []int{}
		for _, part := range strings.Fields(line) {
			num, err := strconv.Atoi(part)
			if err == nil {
				numbers = append(numbers, num)
			} else {
				return nil, fmt.Errorf("error converting string to int: %v", err)
			}
		}

		rows = append(rows, numbers)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return rows, nil
}

func isRowSafe(row []int) bool {
	if len(row) == 1 {
		return true
	}

	increasing := row[1] > row[0]

	for i := 1; i < len(row); i++ {
		diff := row[i] - row[i-1]
		if diff < 0 {
			diff *= -1
		}

		if diff < 1 || diff > 3 {
			return false
		}

		if increasing != (row[i-1] < row[i]) {
			return false
		}
	}

	return true
}

func isRowSafeWithTolerance(row []int) bool {
	if isRowSafe(row) {
		return true
	}

	for i := 0; i < len(row); i++ {
		modifiedRow := append([]int{}, row[:i]...)
		modifiedRow = append(modifiedRow, row[i+1:]...)
		if isRowSafe(modifiedRow) {
			return true
		}
	}

	return false
}

func (c *ChristmasSaver) CountSafeReports() (int, error) {
	rows, err := readRowsDay2()
	if err != nil {
		return 0, err
	}

	safeCount := 0

	for _, row := range rows {
		if isRowSafe(row) {
			safeCount++
		}
	}

	return safeCount, nil
}

func (c *ChristmasSaver) CountSafeReportsTolerance() (int, error) {
	rows, err := readRowsDay2()
	if err != nil {
		return 0, err
	}

	safeCount := 0

	for _, row := range rows {
		if isRowSafeWithTolerance(row) {
			safeCount++
		}
	}

	return safeCount, nil
}
