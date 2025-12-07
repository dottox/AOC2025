package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readInput() ([]string, []string) {
	var freshRanges, ingredients []string

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error while opening the input file")
		return nil, nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	areRanges := true
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			areRanges = false
			continue
		}

		if areRanges {
			freshRanges = append(freshRanges, line)
		} else {
			ingredients = append(ingredients, line)
		}
	}

	return freshRanges, ingredients
}

func getRanges(freshRanges []string) ([]int, []int) {
	var startRanges, endRanges []int
	var s, e int

	for _, r := range freshRanges {
		fmt.Sscanf(r, "%d-%d\n", &s, &e)
		startRanges = append(startRanges, s)
		endRanges = append(endRanges, e)
	}

	return startRanges, endRanges
}

func countFreshIngredients(startRanges, endRanges []int, ingredients []string) int {
	counter := 0

	for _, i := range ingredients {

		ingInt, err := strconv.Atoi(i)
		if err != nil {
			fmt.Printf("Error converting %s to int\n", i)
			continue
		}

		for r := 0; r < len(startRanges); r++ {
			if ingInt >= startRanges[r] && ingInt <= endRanges[r] {
				fmt.Printf("New fresh ingredient %d in range: %d, %d\n", ingInt, startRanges[r], endRanges[r])
				counter++
				break
			}
		}
	}

	return counter
}

func main() {
	freshRanges, ingredients := readInput()

	startRanges, endRanges := getRanges(freshRanges)
	counter := countFreshIngredients(startRanges, endRanges, ingredients)

	fmt.Printf("Fresh ingredients: %d\n", counter)
}
