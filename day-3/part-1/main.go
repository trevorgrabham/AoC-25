package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputData, err := os.ReadFile("../input/challenge.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputData = bytes.TrimSpace(inputData)

	var joltageSum int
	// break input into banks (separate on newlines)
	banks := bytes.SplitSeq(inputData, []byte{'\n'})

	// for each bank
	for bank := range banks {
		maxFirstIndex := 0
		maxSecondIndex := len(bank) - 1

		//		find the largest first digit (excluding the last digit)
		for firstDigitIndex := 0; firstDigitIndex < len(bank)-1; firstDigitIndex++ {
			if bank[firstDigitIndex] == 9 {
				maxFirstIndex = firstDigitIndex
				break
			}

			if bank[firstDigitIndex] > bank[maxFirstIndex] {
				maxFirstIndex = firstDigitIndex
			}
		}

		//		find the largest digit that comes after the largest first digit
		for secondDigitIndex := maxFirstIndex + 1; secondDigitIndex < len(bank); secondDigitIndex++ {
			if bank[secondDigitIndex] == 9 {
				maxSecondIndex = secondDigitIndex
				break
			}

			if bank[secondDigitIndex] > bank[maxSecondIndex] {
				maxSecondIndex = secondDigitIndex
			}
		}

		//		combine them and store their value
		joltage, err := strconv.Atoi(string([]byte{bank[maxFirstIndex], bank[maxSecondIndex]}))
		if err != nil {
			log.Fatal(err)
		}

		// sum up the joltages from all of the banks
		joltageSum += joltage
	}

	fmt.Println(joltageSum)
}
