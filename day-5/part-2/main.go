package main

import (
	"bytes"
	"cmp"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

type IDRange struct {
	start int
	end   int
}

func sortIDRanges(a, b IDRange) int {
	return cmp.Compare(a.start, b.start)
}

func main() {
	inputData, err := os.ReadFile("../input/challenge.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputData = bytes.TrimSpace(inputData)

	var numberOfFreshIngredients int

	// split the input into a list of ranges
	splitInput := bytes.Split(inputData, []byte("\n\n"))
	rangeList := splitInput[0]

	// split the range list into a slice of ranges
	ranges := bytes.Split(rangeList, []byte("\n"))

	// for each range
	idRanges := make([]IDRange, 0, len(ranges))
	for _, rng := range ranges {
		// split the range into a start and end
		splitRange := bytes.Split(rng, []byte("-"))
		startID, endID := splitRange[0], splitRange[1]
		start, err := strconv.Atoi(string(startID))
		if err != nil {
			log.Fatal(err)
		}
		end, err := strconv.Atoi(string(endID))
		if err != nil {
			log.Fatal(err)
		}

		// add to idRanges
		idRanges = append(idRanges, IDRange{start: start, end: end})
	}

	// sort idRanges
	slices.SortFunc(idRanges, sortIDRanges)

	// for each IDRange (excluding the first)
	updatedRanges := make([]IDRange, 0, len(idRanges))
	updatedRanges = append(updatedRanges, idRanges[0])
	largestEnd := idRanges[0].end
	for i := 1; i < len(idRanges); i++ {
		// if the current start > largestEnd
		if idRanges[i].start > largestEnd {
			// add current range to the updated ranges and continue
			updatedRanges = append(updatedRanges, idRanges[i])

			// update the largestEnd
			largestEnd = idRanges[i].end
			continue
		}

		// if the current end <= largestEnd
		if idRanges[i].end <= largestEnd {
			// do nothing, continue
			continue
		}

		// else update the current start to previous' end + 1 and add to the updated ranges
		updatedRanges = append(updatedRanges, IDRange{start: largestEnd + 1, end: idRanges[i].end})

		// update the largestEnd
		largestEnd = idRanges[i].end
	}

	// for each IDRange in the updated ranges
	for _, idRange := range updatedRanges {
		// add end - start + 1 to the number of fresh ingredients
		numberOfFreshIngredients += idRange.end - idRange.start + 1
	}

	fmt.Println(numberOfFreshIngredients)
}
