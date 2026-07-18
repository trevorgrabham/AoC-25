package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputData, err := os.ReadFile("input/challenge.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputData = bytes.TrimSpace(inputData)

	var timesZeroCrossed int
	dialPosition := 50

	inputLines := strings.SplitSeq(string(inputData), "\n")

	// for each input line
	for inputLine := range inputLines {
		// grab the direction
		direction := inputLine[0]

		// grab the number of rotations
		var numberOfRotations int
		numberOfRotations, err = strconv.Atoi(inputLine[1:])
		if err != nil {
			log.Fatal(err)
		}

		if numberOfRotations >= 100 {
			timesZeroCrossed += numberOfRotations / 100
			numberOfRotations = numberOfRotations % 100
		}

		// check if we are going to cross the zero point
		switch direction {
		case 'L':
			// case direction 'L': check if number of rotations >= dialPosition && dialPosition > 0
			if numberOfRotations >= dialPosition && dialPosition > 0 {
				timesZeroCrossed++
			}
			dialPosition = ((dialPosition-numberOfRotations)%100 + 100) % 100
		case 'R':
			// case direction 'R': check if number of rotations >= 100 - dialPosition
			if numberOfRotations >= 100-dialPosition {
				timesZeroCrossed++
			}
			dialPosition = (dialPosition + numberOfRotations) % 100
		default:
			log.Fatalf("unknown direction %s", string(direction))
		}
	}

	fmt.Println(timesZeroCrossed)
}
