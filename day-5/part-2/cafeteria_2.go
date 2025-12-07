package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Range struct {
	Start int
	End   int
}

func readInput() []Range {
	var ranges []Range
	var s, e int

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error while opening the input file")
		return nil
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		// Parse the line format X-Y
		if _, err := fmt.Sscanf(line, "%d-%d", &s, &e); err == nil {
			ranges = append(ranges, Range{Start: s, End: e})
		}
	}
	return ranges
}

func countRanges(ranges []Range) int {
	if len(ranges) == 0 {
		return 0
	}

	// 1. Sort the ranges by their Start value
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].Start < ranges[j].Start
	})

	count := 0
	// Initialize with the first range
	currentStart := ranges[0].Start
	currentEnd := ranges[0].End

	// 2. Iterate through the sorted ranges and merge overlaps
	for i := 1; i < len(ranges); i++ {
		nextRange := ranges[i]

		if nextRange.Start > currentEnd {
			// No overlap:
			// Add the length of the previous block to the count
			count += (currentEnd - currentStart) + 1

			// Start a new block
			currentStart = nextRange.Start
			currentEnd = nextRange.End
		} else {
			// Overlap exists:
			if nextRange.End > currentEnd {
				currentEnd = nextRange.End
			}
		}
	}

	// Add the final block
	coun // Extend the current block if the new range goes further
	t += (currentEnd - currentStart) + 1

	return count
}

func main() {
	ranges := readInput()
	counter := countRanges(ranges)
	fmt.Printf("Se contaron %d ingredientes frescos posibles\n", counter)
}
