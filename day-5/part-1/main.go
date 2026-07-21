package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

	type idRange struct {
		start int 
		end int
	}

	func (r idRange) contains(id int) bool {
	return id >= r.start && id <= r.end
}

func main() {
	inputData, err := os.ReadFile("../input/example.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputData = bytes.TrimSpace(inputData)

	var numberOfFreshIngredients int

	// split the input into a list of ranges and a list of ingredient ids 
	splitInput := bytes.Split(inputData, []byte("\n\n"))
	rangeList, ingredientsList := splitInput[0], splitInput[1]

	// split the range list into a slice of ranges 
	ranges := bytes.Split(rangeList, []byte("\n"))

	// convert each range into an idRange struct 
	idRanges := make([]idRange, 0, len(ranges))
	for _, rng := range ranges {
		splitRange := bytes.Split(rng, []byte("-"))
		startID, endID := splitRange[0], splitRange[1]
		start, err := strconv.Atoi(string(startID))
		if err != nil { log.Fatal(err) }
		end, err := strconv.Atoi(string(endID))
		if err != nil { log.Fatal(err) }
		idRanges = append(idRanges, idRange{start: start, end: end})
	}

	// split the ingredients list into a slice of ingredient ids 
	ingredients := bytes.SplitSeq(ingredientsList, []byte("\n"))

	// for each ingredient 
	for ingredient := range ingredients {
		// convert ingredient into an ingredient id
		ingredientID, err := strconv.Atoi(string(ingredient))
		if err != nil { log.Fatal(err) }

		// for each idRange 
		for _, rng := range idRanges {
			// if the ingredient is contained in the range 
			if rng.contains(ingredientID) {
				// increment the numberOfFreshIngredients and move onto the next ingredient
				numberOfFreshIngredients++
				break
			}
		}
	}

	fmt.Println(numberOfFreshIngredients)
}
