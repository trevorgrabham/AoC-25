#!/usr/bin/env bash

next_script_template=$(cat << 'EOF'
  #!/usr/bin/env bash 
  
  if [[ $PWD =~ /day-([0-9]+)/part-2$ ]]; then 
    day=${BASH_REMATCH[1]}
    ((day++))
  
    cd ../..
    . scaffold.sh "$day"
  else 
    echo "invalid working dir $PWD"
  fi
EOF
)

go_template=$(cat << 'EOF'
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
echo "$go_template" > main.go 
go mod init "github.com/trevorgrabham/AoC-26/day-$day/part-1"
touch input/example.txt input/challenge.txt

cd "../part-2"
echo "$go_template" > main.go 
echo "$next_script_template" > next.sh
chmod 755 next.sh
go mod init "github.com/trevorgrabham/AoC-26/day-$day/part-2"
touch input/example.txt input/challenge.txt

cd "../part-1"
vi input/example.txt && vi input/challenge.txt && vi main.go
