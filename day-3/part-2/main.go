package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func findBestDigits(bank []byte, digitsRemaining int) []byte {
	digits := make([]byte, 0, digitsRemaining)

	maxDigitIndex := 0
	for digitIndex := 0; digitIndex <= len(bank)-digitsRemaining; digitIndex++ {
		if bank[digitIndex] == '9' {
			maxDigitIndex = digitIndex
			break
		}

		if bank[digitIndex] > bank[maxDigitIndex] {
			maxDigitIndex = digitIndex
		}
	}

	if digitsRemaining <= 1 {
		return []byte{bank[maxDigitIndex]}
	}

	return append(append(digits, bank[maxDigitIndex]), findBestDigits(bank[maxDigitIndex+1:], digitsRemaining-1)...)
}

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
		bestDigits := findBestDigits(bank, 12)

		//		combine them and store their value
		joltage, err := strconv.Atoi(string(bestDigits))
		if err != nil {
			log.Fatal(err)
		}

		// sum up the joltages from all of the banks
		joltageSum += joltage
	}

	fmt.Println(joltageSum)
}
