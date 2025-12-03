package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput() []string {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var inputs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		inputs = append(inputs, strings.Split(line, ",")...)
	}

	return inputs
}

// Pre-requisite: s length is divisible by chunkSize
func divideString(s string, chunkSize int) []string {
	var chunks []string

	var start int
	for start < len(s) {
		chunks = append(chunks, s[start:start+chunkSize])
		start += chunkSize
	}

	return chunks
}

func checkValidity(id int) bool {
	idStr := strconv.Itoa(id)

	length := len(idStr)

	// Iteramos sobre cada posible divisor del largo string
	for i := 1; i < length; i++ {
		if length%i != 0 { // Si no es divisible, vamos con el siguiente
			continue
		}

		// Dividimos el string en sub-strings de largo i
		subStrs := divideString(idStr, i)

		// Comparamos los sub-strings entre sí
		// Asumimos que son iguales, cuando encontremos uno diferente, lo desmentimos
		allEqual := true
		for j := 0; j < len(subStrs)-1; j++ {
			if subStrs[j] != subStrs[j+1] {
				allEqual = false
				break
			}
		}

		// Si todos los sub-strings son iguales -> id inválida
		if allEqual {
			return false
		}
	}

	return true

}

// rang is in format "start-end", e.g. "10-50"
// returns start and end as integers
func getRange(rang string) (int, int) {
	var start, end int
	fmt.Sscanf(rang, "%d-%d", &start, &end)
	return start, end
}

func main() {
	inputs := readInput()

	sum := 0
	for _, input := range inputs {
		start, end := getRange(input)
		for id := start; id <= end; id++ {
			if !checkValidity(id) {
				sum += id
			}
		}
	}

	fmt.Printf("Sum of all valid gift IDs: %d\n", sum)
}
