package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	inputData, err := os.ReadFile("../input/challenge.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputData = bytes.TrimSpace(inputData)

	var reachableRolls int
	var matrix [][]rune

	// split the input into lines
	inputLines := bytes.SplitSeq(inputData, []byte{'\n'})

	// convert each line into a rune slice
	// collect all the rows back into a 2D rune matrix
	for line := range inputLines {
		matrix = append(matrix, []rune(string(line)))
	}

	// for each entry in the matrix
	for row := 0; row < len(matrix); row++ {
		for column := 0; column < len(matrix[0]); column++ {
			// if the entry is not a roll, continue
			if matrix[row][column] != '@' {
				continue
			}

			// check the number of neighbours
			reachable := checkReachable(matrix, row, column)

			// if less than 4 neighbours, increment the total
			if reachable {
				reachableRolls++
			}
		}
	}

	fmt.Println(reachableRolls)
}

func checkReachable(m [][]rune, row, col int) bool {
	neighbours := findNeighbours(m, row, col)

	var neighbourCount int
	for _, n := range neighbours {
		if n == '@' {
			neighbourCount++
		}
	}

	return neighbourCount < 4
}

func findNeighbours(m [][]rune, row, col int) []rune {
	var neighbours []rune
	if row > 0 {
		// top middle
		neighbours = append(neighbours, m[row-1][col])
		if col > 0 {
			// top left
			neighbours = append(neighbours, m[row-1][col-1])
		}
		if col < len(m[0])-1 {
			// top right
			neighbours = append(neighbours, m[row-1][col+1])
		}
	}

	if row < len(m)-1 {
		// bottom middle
		neighbours = append(neighbours, m[row+1][col])
		if col > 0 {
			// bottom left
			neighbours = append(neighbours, m[row+1][col-1])
		}
		if col < len(m[0])-1 {
			// bottom right
			neighbours = append(neighbours, m[row+1][col+1])
		}
	}

	if col > 0 {
		// middle left
		neighbours = append(neighbours, m[row][col-1])
	}

	if col < len(m[0])-1 {
		// middle right
		neighbours = append(neighbours, m[row][col+1])
	}

	return neighbours
}
