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

	if idLength % 2 != 0 { return 0 }

	for i := 0; i < idLength/2; i++ {
		if idString[i] != idString[idLength/2+i] { return 0 }
	}
	return id
}

func main() {
	inputData, err := os.ReadFile("input/challenge.txt")
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
		if err != nil { log.Fatal(err) }

		// parse ending number
		end, err := strconv.Atoi(rangeEnds[1])
		if err != nil { log.Fatal(err) }

		// check each number in the range (closed)
		for candidateID := start; candidateID <= end; candidateID++ {
			res := validateID(candidateID)

			if res != 0 { fmt.Printf("invalid ID %d\n", candidateID) }
			sumOfInvalidIDs += res
		}
	}

	fmt.Println(sumOfInvalidIDs)
}
