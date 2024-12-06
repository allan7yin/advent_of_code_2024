package main

import (
	"fmt"

	d "github.com/allan7yin/advent_of_code_2024/days"
)

func main() {
	cs := &d.ChristmasSaver{}

	// Day 1
	fmt.Println("+++++=======+++++")
	fmt.Println("Day 1")
	distance, err := cs.GetListDistance()
	if err != nil {
		fmt.Println("Error calculating list distance:", err)
		return
	}
	fmt.Println("List Distance:", distance)

	score, err := cs.GetSimilarityScore()
	if err != nil {
		fmt.Println("Error calculating similarity score:", err)
		return
	}
	fmt.Println("Similarity Score:", score)

	// Day 2
	fmt.Println("+++++=======+++++")
	fmt.Println("Day 2")
	safeCount, err := cs.CountSafeReports()
	if err != nil {
		fmt.Println("Error counting number of safe reports:", err)
		return
	}
	fmt.Println("Number of safe reports:", safeCount)

	safeCountTol, err := cs.CountSafeReportsTolerance()
	if err != nil {
		fmt.Println("Error counting number of safe reports:", err)
		return
	}
	fmt.Println("Number of safe reports:", safeCountTol)

	// Day 3
	fmt.Println("+++++=======+++++")
	fmt.Println("Day 3")
	uncorruptedCount, err := cs.FindUncorruptedMults()
	fmt.Println("Number of uncorrupted multiplications:", uncorruptedCount)
	if err != nil {
		fmt.Println("Error computing sum of multiplications:", err)
		return
	}

	uncorruptedCountDisable, err := cs.FindUncorruptedMultsDisable()
	fmt.Println("Number of uncorrupted multiplications with disabled portions:", uncorruptedCountDisable)
	if err != nil {
		fmt.Println("Error computing sum of multiplications:", err)
		return
	}

	// Day 4
	fmt.Println("+++++=======+++++")
	fmt.Println("Day 4")
	numWords, err := cs.FindAllInstances("XMAS")
	fmt.Println("Number of occurrences of target word:", numWords)

	numCrossWords, err := cs.FindAllCrossInstances()
	fmt.Println("Number of occurrences of target cross word:", numCrossWords)
}
