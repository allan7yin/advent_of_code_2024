package main

import (
	"fmt"

	d "github.com/allan7yin/advent_of_code_2024/days"
)

func main() {
	cs := &d.ChristmasSaver{}

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
}
