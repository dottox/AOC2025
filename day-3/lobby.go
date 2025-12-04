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
		inputs = append(inputs, line)
	}

	return inputs
}

func convertByteToInt(b byte) int {
	return int(b - '0')
}

type Battery struct {
	bankIndex int
	joltage   int
}

func getJoltage(bank string, numOfBatteries int) (int, error) {
	batts := make([]Battery, numOfBatteries)
	for i := range numOfBatteries {
		batts[i] = Battery{bankIndex: i, joltage: -1}
	}

	fmt.Println("Processing bank:", bank)
	// Iteramos sobre cada valor del bank
	for i := 0; i < len(bank); i++ {
		currBankDigit := convertByteToInt(bank[i])

		// Iteramos sobre el numero de baterias que necesitamos
		for b := 0; b < numOfBatteries; b++ {
			if b > i {
				break
			}

			// Chequeamos si quedan suficientes digitos en el bank para llenar las baterias restantes
			if numOfBatteries-b > len(bank)-i {
				continue
			}

			// Chequeamos que ninguna de las baterias anteriores esté usando este dígito
			alreadyUsed := false
			for prevB := 0; prevB < b; prevB++ {
				if batts[prevB].bankIndex >= i {
					alreadyUsed = true
				}
			}
			if alreadyUsed {
				break
			}

			// Si el valor actual es mayor al guardado en la bateria b, lo actualizamos
			if currBankDigit > batts[b].joltage {
				batts[b].joltage = currBankDigit
				batts[b].bankIndex = i

				for nextB := b + 1; nextB < numOfBatteries; nextB++ {
					batts[nextB].joltage = convertByteToInt(bank[i+(nextB-b)])
					batts[nextB].bankIndex = i + (nextB - b)
				}
				break
			}
		}

	}

	var finalJoltageList []string
	for _, b := range batts {
		finalJoltageList = append(finalJoltageList, fmt.Sprint(b.joltage))
	}

	finalJoltageInt, err := strconv.Atoi(strings.Join(finalJoltageList, ""))
	if err != nil {
		return 0, err
	}

	return finalJoltageInt, nil
}

func main() {
	NUM_OF_BATTERIES := 12

	banks := readInput()

	joltageSum := 0
	for _, bank := range banks {
		joltage, err := getJoltage(bank, NUM_OF_BATTERIES)
		if err != nil {
			fmt.Println(err)
		}

		joltageSum += joltage
		fmt.Println("Joltage for bank", bank, "is", joltage)
	}

	fmt.Println("Total joltage sum is", joltageSum)
}
