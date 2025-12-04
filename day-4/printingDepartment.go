package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func readInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		return nil
	}
	defer file.Close()

	rolls := make([]string, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		rolls = append(rolls, line)
	}

	return rolls
}

func deleteFromRange(r []int, elem int) []int {
	if len(r) == 0 {
		return r
	}

	return slices.DeleteFunc(r, func(i int) bool {
		return i == elem
	})
}

func countAccessibleRolls(rolls []string) int {

	counter := 0

	for r, rollLine := range rolls {
		for c, roll := range rollLine {
			if roll == rune('@') {

				var check func(r, c int)
				check = func(r, c int) {
					countRollsAdjacent := 0
					checkingRange := make(map[int][]int) // key == row, value == column

					// Complete the range to search
					checkingRange[0] = []int{-1, 1}
					if r != 0 {
						checkingRange[-1] = []int{-1, 0, 1}
					}
					if r != len(rolls)-1 {
						checkingRange[1] = []int{-1, 0, 1}
					}

					// Deduct the range using the columns
					if c == 0 {
						// Delete the x,-1 from all the ranges
						checkingRange[-1] = deleteFromRange(checkingRange[-1], -1)
						checkingRange[0] = deleteFromRange(checkingRange[0], -1)
						checkingRange[1] = deleteFromRange(checkingRange[1], -1)
					}
					if c == len(rolls)-1 {
						// Delete the x,1 from all the ranges
						checkingRange[-1] = deleteFromRange(checkingRange[-1], 1)
						checkingRange[0] = deleteFromRange(checkingRange[0], 1)
						checkingRange[1] = deleteFromRange(checkingRange[1], 1)
					}

					for row, listOfColumns := range checkingRange {
						for _, column := range listOfColumns {
							if rolls[r+row][c+column] == '@' {
								countRollsAdjacent++
							}
						}
					}

					if countRollsAdjacent < 4 && rolls[r][c] == byte('@') {
						rolls[r] = rolls[r][:c] + "x" + rolls[r][c+1:]
						counter++
						for row, listOfColumns := range checkingRange {
							for _, column := range listOfColumns {
								if rolls[r+row][c+column] == '@' {
									check(r+row, c+column)
								}
							}
						}
					}
				}

				check(r, c)
			}
		}

		//fmt.Printf("Founded %d accessible rolls in row %d\n", counter, r)
	}

	return counter
}

func main() {
	rolls := readInput()
	counter := countAccessibleRolls(rolls)

	fmt.Printf("Founded %d accessible rolls\n", counter)
}
