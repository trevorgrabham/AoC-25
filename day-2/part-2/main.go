package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func validateID(id int) int {
	idString := strconv.Itoa(id)
	idLength := len(idString)

	// check patterns of length 1 through length 5 (largest id has 10 digits)
patternLengthLoop:
	for patternLength := 1; patternLength <= 5; patternLength++ {
		// if the id is shorter than the pattern length or isn't evenly divisible by the pattern length, try the next one
		if idLength <= patternLength || idLength % patternLength != 0 { continue }

		// the pattern should be repeated idLength/patternLength times
		patternRepetitions := idLength/patternLength

		// check each position in the pattern 
		for patternPosition := range patternLength {

			// check that each repetition matches in the appropriate position offset for the pattern
			for patternRepetition := 1; patternRepetition < patternRepetitions; patternRepetition++ {
				// if we find a mismatch, continue onto the next pattern length
				if idString[patternPosition] != idString[patternRepetition * patternLength + patternPosition] { continue patternLengthLoop }
			}
		}

		// if there were no mismatches, we found an invalid ID
		return id
	}
	return 0
}

func main() {
	inputData, err := os.ReadFile("../input/challenge.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputData = bytes.TrimSpace(inputData)

	var sumOfInvalidIDs int

	// split into ranges
	ranges := strings.SplitSeq(string(inputData), ",")

	// for each range
	for idRange := range ranges {
		rangeEnds := strings.Split(idRange, "-")
		// parse starting number
		start, err := strconv.Atoi(rangeEnds[0])
		if err != nil {
			log.Fatal(err)
		}

		// parse ending number
		end, err := strconv.Atoi(rangeEnds[1])
		if err != nil {
			log.Fatal(err)
		}

		// check each number in the range (closed)
		for candidateID := start; candidateID <= end; candidateID++ {
			res := validateID(candidateID)

			if res != 0 {
				fmt.Printf("invalid ID %d\n", candidateID)
			}
			sumOfInvalidIDs += res
		}
	}

	fmt.Println(sumOfInvalidIDs)
}
