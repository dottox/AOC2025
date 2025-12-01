package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func moveDialPartTwo(direction rune, steps int, dial int, counter int) (int, int) {
	startDial := dial
	isDialInitZero := dial == 0
	if direction == rune('L') {
		dial = dial - steps
	} else {
		dial = dial + steps
	}

	hitsToZero := counter

	for dial > 99 {
		hitsToZero++
		dial = dial - 100
	}

	for dial < 0 {
		if isDialInitZero {
			isDialInitZero = false
		} else {
			hitsToZero++
		}
		dial = dial + 100
	}

	if direction == rune('L') && dial == 0 {
		hitsToZero++
	}

	fmt.Printf("Start dial: %d -> New dial: %d  |  Move: %c, steps: %d  |  Counter: %d\n", startDial, dial, direction, steps, hitsToZero)
	return dial, hitsToZero
}

// part 1
// func moveDial(direction rune, steps int, dial int) int {

// 	if direction == rune('L') {
// 		dial = dial - steps
// 	} else {
// 		dial = dial + steps
// 	}

// 	if dial > 99 {
// 		dial = dial % 100
// 	} else if dial < 0 {
// 		dial = -dial
// 		dial = dial % 100
// 		dial = 100 - dial
// 		if dial == 100 {
// 			dial = 0
// 		}
// 	}

// 	return dial
// }

func processMovement(movement string, dial int, counter int) (int, int) {
	direction := rune(movement[0])
	stepsStr := movement[1:]

	if direction != rune('L') && direction != rune('R') {
		log.Fatalf("invalid direction %c", direction)
	}

	steps, err := strconv.Atoi(stepsStr)
	if err != nil {
		log.Fatalf("invalid steps %q: %v\n", stepsStr, err)
	}

	// part 1
	// return moveDial(direction, steps, dial)
	return moveDialPartTwo(direction, steps, dial, counter)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	dial := 50
	counter := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dial, counter = processMovement(scanner.Text(), dial, counter)
		// if dial == 0 {
		//	counter++
		// }
	}

	fmt.Println(counter)
}
