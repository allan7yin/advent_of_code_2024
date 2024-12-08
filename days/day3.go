package days

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func readRowsDay3() (string, error) {
	input := ""
	file, err := os.Open("./dataFiles/day3.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = line
	}

	return input, nil
}

func removeDisabledIntervals(input string) (string, error) {
	return regexp.MustCompile(`don't\(\).*?(do\(\)|$)`).ReplaceAllString(input, ""), nil
}

func computeSumOfProds(input string) (uint64, error) {
	var result uint64 = 0
	pattern := `mul\([0-9]{1,3},[0-9]{1,3}\)`
	r, _ := regexp.Compile(pattern)
	operations := r.FindAllString(input, -1)

	for _, op := range operations {
		n := len(op)
		nums := op[4 : n-1]
		parts := strings.Split(nums, ",")
		if num1, err1 := strconv.Atoi(parts[0]); err1 == nil {
			if num2, err2 := strconv.Atoi(parts[1]); err2 == nil {
				result += uint64(num1) * uint64(num2)
			}
		}
	}

	return result, nil
}

func (c *ChristmasSaver) FindUncorruptedMults() (uint64, error) {
	input, err := readRowsDay3()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0, err
	}
	return computeSumOfProds(input)
}

func (c *ChristmasSaver) FindUncorruptedMultsDisable() (uint64, error) {
	input, err := readRowsDay3()
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 0, err
	}

	scrapedInput, err := removeDisabledIntervals(input)
	if err != nil {
		fmt.Println("Error removing disabled segment:", err)
		return 0, err
	}

	return computeSumOfProds(scrapedInput)
}
