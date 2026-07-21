package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

const (
	beam     byte = '|'
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
	numberTimelines := 1
	var timesTraveresed [][]int

	// transform the input into a row x col matrix
	matrix := bytes.Split(inputData, []byte("\n"))
	timesTraveresed = make([][]int, len(matrix))
	for i := range matrix {
		timesTraveresed[i] = make([]int, len(matrix[i]))
	}

	// find the 'S' in the first row
	startIndex := bytes.IndexByte(matrix[0], byte('S'))
	if startIndex < 0 {
		log.Fatalf("unable to locate start in:\n%s\n", matrix[0])
	}

	// look below the start
	//// if it's a splitter, transform either side of it into a beam
	if matrix[1][startIndex] == splitter {
		numberTimelines++
		matrix[1][startIndex-1], matrix[1][startIndex+1] = beam, beam
		timesTraveresed[1][startIndex-1]++
		timesTraveresed[1][startIndex+1]++

		//// else transform it into a beam
	} else {
		matrix[1][startIndex] = beam
		timesTraveresed[1][startIndex]++
	}

	// for every row [second row, end)
	for row := 1; row < len(matrix)-1; row++ {
		//// scan each col in the row
		for col := 0; col < len(matrix[row]); col++ {
			////// if it isn't a beam, continue on
			if matrix[row][col] != beam {
				continue
			}

			////// if it's a splitter, transform either side of it into a beam
			if matrix[row+1][col] == splitter {
				numberTimelines += timesTraveresed[row][col]
				matrix[row+1][col-1], matrix[row+1][col+1] = beam, beam
				timesTraveresed[row+1][col-1] += timesTraveresed[row][col]
				timesTraveresed[row+1][col+1] += timesTraveresed[row][col]

				////// else transform it into a beam
			} else {
				matrix[row+1][col] = beam
				timesTraveresed[row+1][col] += timesTraveresed[row][col]
			}
		}
	}

	fmt.Println(numberTimelines)
}
