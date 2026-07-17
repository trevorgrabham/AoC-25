package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	inputData, err := os.ReadFile("input/example.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(inputData))
}
