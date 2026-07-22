package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"sync/atomic"
)

func abs(a int) int { 
	if a < 0 { a = a * -1 } 
	return a 
}

type Point struct {
	x int 
	y int 
}

func (p Point) String() string { return fmt.Sprintf("(%d, %d)", p.x, p.y) }

func (p Point) RectArea(other Point) int { 
	dx := abs(p.x - other.x) + 1
	dy := abs(p.y - other.y) + 1 
	return dx * dy
}

func main() {
	inputData, err := os.ReadFile("../input/challenge.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputData = bytes.TrimSpace(inputData)

	// vars 
	//// points ([]Point)
	//// maxWorkers (int)
	//// chunkSize (int)
	//// wg (sync.WaitGroup)
	//// maxDistance (atomic.Uint64)
	var (
		points []Point
		chunkSize int
		wg sync.WaitGroup
		maxDistance atomic.Uint64
	)
	numWorkers := runtime.NumCPU()

	// split the input into lines 
	lines := bytes.Split(inputData, []byte("\n"))

	// prepare points
	points = make([]Point, 0, len(lines))

	// for each line 
	for _, line := range lines {
	//// separate fields and append a new point to points 
		fields := bytes.Split(line, []byte(","))
		xField, yField := fields[0], fields[1]

		x, err := strconv.Atoi(string(xField))
		if err != nil { panic(err) }

		y, err := strconv.Atoi(string(yField))
		if err != nil { panic(err) }

		points = append(points, Point{x: x, y: y})
	}

	// calculate the work chunk size (num points / num workers)
	numWorkers = min(numWorkers, len(points))
	chunkSize = len(points) / numWorkers

	// create the work closure
	calcArea := func(start, end int) {
	//// setup a local max distance 
		var localMax int

	//// for each point in the chunk 
		for i := start; i < end; i++ {
	////// calculate the distance to every other point (even outside of the chunk)
			for j := 0; j < len(points); j++ {
				if j == i { continue }

	//////// if distance is larger than local max, update it
				dist := points[i].RectArea(points[j])
				if dist > localMax { localMax = dist }
			}
		}

		for {
	//// read global max
			globalMax := maxDistance.Load()

	//// if local max < global max break
			if localMax < int(globalMax) {
				break
			}

	//// CAS 
			if maxDistance.CompareAndSwap(globalMax, uint64(localMax)) { break }
		}

		wg.Done()
	}

	// for num workers
	for i := range numWorkers {
	//// add one to wg 
		wg.Add(1)

	//// calculate start and end (i * chunkSize, (i+1)*chunkSize), if last chunk end = len(points)
		start, end := i * chunkSize, (i+1) * chunkSize
		if i >= numWorkers - 1 { end = len(points) }

	//// call the work closure
		go calcArea(start, end)
	}

	// wait for wg 
	wg.Wait()

	fmt.Println(maxDistance.Load())
}
