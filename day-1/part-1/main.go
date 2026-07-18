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

	var numberOfZeros int
	dialPosition := 50

	inputLines := strings.SplitSeq(string(inputData), "\n")

	// for each input line
	for inputLine := range inputLines {
		//		grab the direction
		direction := inputLine[0]

		//		grab the number of rotations
		var numberRotations int
		numberRotations, err = strconv.Atoi(inputLine[1:])
		if err != nil {
			log.Fatal(err)
		}

		//		update dialPosition and check if dialPosition == 0
		switch direction {
		case 'L':
			dialPosition = (dialPosition - numberRotations) % 100
		case 'R':
			dialPosition = (dialPosition + numberRotations) % 100
		default:
			log.Fatalf("unknown direction %s", string(direction))
		}
		if dialPosition == 0 {
			numberOfZeros++
		}
	}

	fmt.Println(numberOfZeros)
}
