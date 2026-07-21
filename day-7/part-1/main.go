package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

const (
	beam byte = '|'
	splitter byte = '^'
)

func main() {
	inputData, err := os.ReadFile("../input/challenge.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputData = bytes.TrimSpace(inputData)

	// vars
	//// numberSplits
	var numberSplits int

	// transform the input into a row x col matrix
	matrix := bytes.Split(inputData, []byte("\n"))

	// find the 'S' in the first row 
	startIndex := bytes.IndexByte(matrix[0], byte('S'))
	if startIndex < 0 { log.Fatalf("unable to locate start in:\n%s\n", matrix[0]) }

	// look below the start 
	//// if it's a splitter, transform either side of it into a beam
	if matrix[1][startIndex] == splitter {
		numberSplits++
		matrix[1][startIndex-1], matrix[1][startIndex+1] = beam, beam

	//// else transform it into a beam
	} else { 
		matrix[1][startIndex] = beam
	}

	// for every row [second row, end)
	for row := 1; row < len(matrix) - 1; row++ {
	//// scan each col in the row
		for col := 0; col < len(matrix[row]); col++ {
	////// if it isn't a beam, continue on
			if matrix[row][col] != beam { continue }

	////// if it's a splitter, transform either side of it into a beam
			if matrix[row+1][col] == splitter {
				numberSplits++
				matrix[row+1][col-1], matrix[row+1][col+1] = beam, beam

	////// else transform it into a beam
			} else {
				matrix[row+1][col] = beam
			}
		}
	}

	fmt.Println(numberSplits)
}
