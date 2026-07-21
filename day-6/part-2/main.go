package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func add(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

type Problem struct {
	numbers []int 
	operation func(a, b int) int
	isMult bool
}

func (p Problem) Solution() (sum int) {
	if p.isMult { sum = 1 }
	for _, num := range p.numbers {
		sum = p.operation(sum, num)
	}
	return
}

func (p *Problem) SetOperation(operator string) {
	switch operator {
	case "+":
		p.operation = add
	case "*":
		p.operation = multiply
		p.isMult = true
	default:
		log.Fatalf("unknown operation %s\n", operator) 
	}
}

func main() {
	inputData, err := os.ReadFile("../input/challenge.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := bytes.Split(bytes.TrimSpace(inputData), []byte("\n"))

	// vars
	//// operations
	//// numberlist
	//// numbers
	//// running sum
	//// problems
	var (
		problemIndex, runningSum int
		operations []byte
		operationLine = input[len(input)-1:len(input):len(input)][0]
		numberList [][]byte
		problems []Problem
	)

	// parse operations
	ops := bytes.FieldsSeq(operationLine)
	for op := range ops {
		operations = append(operations, op[0])
	}

	// initialize problems
	problems = make([]Problem, len(operations))

	// get each line
	for _, line := range input[:len(input)-1:len(input)-1] {
		// initialize numberlist if not already
		if numberList == nil { numberList = make([][]byte, len(line)) }
		// for each character in the line (by index)
		for index := range line {
		//// append the character to numberList[index]
			numberList[index] = append(numberList[index], line[index])
		}
	}

	// set the first problem operation
	problems[0].SetOperation(string(operations[0]))

	// for each line in numberList 
	for _, line := range numberList {
		var value strings.Builder
	//// for each character in line (by index)
		for index := range line {
	////// if the character is blank, continue
			if line[index] == ' ' { continue }

	////// write the character to value
			value.WriteByte(line[index])
		}

	//// if value is empty 
		if value.String() == "" {
	////// increment problem index and set the new problems operation
			problemIndex++
			problems[problemIndex].SetOperation(string(operations[problemIndex]))

	//// else 
		} else {
	////// convert the number and add it to the current problem
			num, err := strconv.Atoi(value.String())
			if err != nil { log.Fatalf("error converting number\n%v", err) }

			problems[problemIndex].numbers = append(problems[problemIndex].numbers, num)
		}
	}

	// for each problem 
	for _, problem := range problems {
	//// solve it and add it to the runningSum
		runningSum += problem.Solution()
	}

	fmt.Println(runningSum)
}
