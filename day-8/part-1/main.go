package main

import (
	"bytes"
	"cmp"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
)

type Distance struct {
	distance       float64
	firstJunction  *Junction
	secondJunction *Junction
}

func (d Distance) String() string {
	return fmt.Sprintf("dist: %f\t\t%s\t%s", d.distance, d.firstJunction.String(), d.secondJunction.String())
}

type Junction struct {
	ID        int
	CircuitID int
	x         float64
	y         float64
	z         float64
}

func (j *Junction) Distance(other *Junction) Distance {
	dx := j.x - other.x
	dy := j.y - other.y
	dz := j.z - other.z
	return Distance{distance: math.Sqrt(dx*dx + dy*dy + dz*dz), firstJunction: j, secondJunction: other}
}

func (j *Junction) String() string {
	return fmt.Sprintf("%.0f, %.0f, %.0f\t(id: %d\tcircuit: %d)", j.x, j.y, j.z, j.ID, j.CircuitID)
}

func main() {
	inputData, err := os.ReadFile("../input/challenge.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputData = bytes.TrimSpace(inputData)

	// vars
	//// junctionID (int)
	//// circuitID (int)
	//// junctions ([]*Junction)
	//// distances ([]Distance)
	//// circuits ([][]*Junction)
	var (
		junctionID, circuitID int
		junctions             []*Junction
		distances             []Distance
		circuits              = make([][]*Junction, 1000)
	)

	// split input by line
	lines := bytes.Split(inputData, []byte("\n"))

	// prepare junctions (number of lines) and distances ( (len(junctions) * len(junctions)-1) / 2 )
	junctions = make([]*Junction, 0, len(lines))
	numJunctions := len(junctions)
	distances = make([]Distance, 0, numJunctions*(numJunctions-1)/2)

	// for each line
	for _, line := range lines {
		//// separate fields by ','
		coords := bytes.Split(line, []byte(","))

		//// convert each field to a float64
		x, err := strconv.ParseFloat(string(coords[0]), 64)
		if err != nil {
			panic(err)
		}

		y, err := strconv.ParseFloat(string(coords[1]), 64)
		if err != nil {
			panic(err)
		}

		z, err := strconv.ParseFloat(string(coords[2]), 64)
		if err != nil {
			panic(err)
		}

		//// set and increment junctionID, set circuitID = -1
		//// append the values to junctions
		junctions = append(junctions, &Junction{ID: junctionID, CircuitID: -1, x: x, y: y, z: z})
		junctionID++
	}

	// for all junctions (excluding last)
	for i := 0; i < len(junctions)-2; i++ {
		//// for all other junctions after
		for j := i + 1; j < len(junctions); j++ {
			////// calculate a Distance and append to distances
			distances = append(distances, junctions[i].Distance(junctions[j]))
		}
	}

	// sort distances
	slices.SortFunc(distances, func(a, b Distance) int {
		return cmp.Compare(a.distance, b.distance)
	})

	// grab the first 1000 distances
	for i := range 1000 {
		//// get the circuitID for each junction
		first, second := distances[i].firstJunction, distances[i].secondJunction

		//// if they are both -1
		if first.CircuitID == -1 && second.CircuitID == -1 {
			////// update the junctions circuitID's to circuitID
			first.CircuitID = circuitID
			second.CircuitID = circuitID

			////// append both junctionIDs to circuits[circuitID]
			circuits[circuitID] = append(circuits[circuitID], first, second)

			////// inc circuitID
			circuitID++

			//// elif the first has a circuit, the second doesn't
		} else if second.CircuitID == -1 {
			////// update the second circuit to the first's circuitID
			second.CircuitID = first.CircuitID

			////// append the second ID to circuits
			circuits[first.CircuitID] = append(circuits[first.CircuitID], second)

			//// elif the first doesn't have a circuit, the second does
		} else if first.CircuitID == -1 {
			////// update the first circuit to the seconds's circuitID
			first.CircuitID = second.CircuitID

			////// append the first ID to circuits
			circuits[second.CircuitID] = append(circuits[second.CircuitID], first)

			//// if they are both equal, continue
		} else if first.CircuitID == second.CircuitID {
			continue

			//// else both are set
		} else {
			secondCircuit := second.CircuitID

			////// for each junction in circuits[second.CircuitID]
			for _, junc := range circuits[second.CircuitID] {
				//////// set their circuitID to first.CircuitID
				junc.CircuitID = first.CircuitID

				//////// append it to circuits[first.CircuitID]
				circuits[first.CircuitID] = append(circuits[first.CircuitID], junc)
			}

			////// set circuits[second.CircuitID] = nil
			circuits[secondCircuit] = nil
		}
	}

	// sort circuits by length
	slices.SortFunc(circuits, func(a, b []*Junction) int {
		return cmp.Compare(len(b), len(a))
	})

	fmt.Println(len(circuits[0]) * len(circuits[1]) * len(circuits[2]))
}
