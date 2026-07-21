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

func main() {
	inputData, err := os.ReadFile("../input/challenge.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := bytes.SplitSeq(bytes.TrimSpace(inputData), []byte("\n"))

	// vars
	//// operations
	//// numbers
	//// running sum
	//// problems
	numbers := make([][]int, 0)
	var (
		runningSum int
		operations []string
		problems []Problem
	)

	// get each line
	for line := range input {
		// separate the values
		values := strings.Fields(string(line))

		// check if they are operators ('+' or '*')
		if values[0] == "+" || values[0] == "*" {
		//// if they are set them as the operations and continue
			operations = values
			continue
		}

		// add them onto the numbers list 
		nums := make([]int, 0, len(values))
		for _, value := range values {
			num, err := strconv.Atoi(value)
			if err != nil { log.Fatalf("error converting value to number\n%v", err) }
			nums = append(nums, num)
		}
		numbers = append(numbers, nums)
	}

	// make a list of problems default init with len == operations
	problems = make([]Problem, len(operations))

	// for each row of numbers 
	for _, nums := range numbers {
	//// for each index of the row 
		for indx := range nums {
	////// add the number to the corresponding problem using index 
			problems[indx].numbers = append(problems[indx].numbers, nums[indx])
		}
	}

	// for each index of operations 
	for indx := range operations {
	//// add the operation to the corresponding problem using index
		switch operations[indx] {
		case "*":
			problems[indx].operation = multiply
			problems[indx].isMult = true
		case "+":
			problems[indx].operation = add
		default:
			log.Fatalf("error unknown operation\n%s", operations[indx])
		}
	}


	// solve each problem and add it to the runningSum
	for _, problem := range problems {
		runningSum += problem.Solution()
	}

	fmt.Println(runningSum)
}
