#!/usr/bin/env bash

template=$(cat << 'EOF'
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
EOF
)

if (( $# < 1 )); then 
  read -p "Day number: " day
else
  day="$1"
fi
mkdir -p "day-$day/part-1/input" "day-$day/part-2/input"

cd "day-$day/part-1"
echo "$template" > main.go 
go mod init "github.com/trevorgrabham/AoC-26/day-$day/part-1"
touch input/example.txt input/challenge.txt

cd "../part-2"
echo "$template" > main.go 
go mod init "github.com/trevorgrabham/AoC-26/day-$day/part-2"
touch input/example.txt input/challenge.txt

cd "../part-1"
vi input/example.txt
