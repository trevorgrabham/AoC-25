package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/trevorgrabham/AoC-25/day-9/part-2/point"
)

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
	//// tileMatrix ([][]byte)
	var (
		points      point.SortedPoints
		chunkSize   int
		wg          sync.WaitGroup
		maxDistance atomic.Uint64
	)
	numWorkers := 1
	// numWorkers := runtime.NumCPU()

	// split the input into lines
	lines := bytes.Split(inputData, []byte("\n"))

	// for each line
	var prev *point.Point
	for _, line := range lines {
		//// separate fields and append a new point to points
		fields := bytes.Split(line, []byte(","))
		xField, yField := fields[0], fields[1]

		x, err := strconv.Atoi(string(xField))
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(string(yField))
		if err != nil {
			panic(err)
		}

		point := &point.Point{X: x, Y: y, Prev: prev}

		if prev != nil { prev.Next = point }
		points.AddPoint(point)

		prev = point
	}
	points.PointsX[0].Prev, prev.Next = prev, points.PointsX[0]

	points.Sort()
	// for i := range points.PointsX { fmt.Printf("%s\t|\t%s\n", points.PointsX[i], points.PointsY[i]) }
	// fmt.Println()

	for i := range points.PointsX {
		points.PointsX[i].ComputeCornerType(points)
	}

	// calculate the work chunk size (num points / num workers)
	numWorkers = min(numWorkers, len(points.PointsX))
	chunkSize = len(points.PointsX) / numWorkers

	// create the work closure
	calcArea := func(start, end int) {
		//// setup a local max distance
		var localMax int

		//// for each point in the chunk
		for i := start; i < end; i++ {
			////// calculate the distance to every other point (even outside of the chunk)
			for j := 0; j < len(points.PointsX); j++ {
				if j == i {
					continue
				}

				//////// if distance is larger than local max, update it
				dist := points.PointsX[i].RectArea(points, *points.PointsX[j])
				if dist > localMax {
					localMax = dist

					// fmt.Printf("found new local max: %d\t\t%s - %s\n", dist, points.PointsX[i], points.PointsX[j])
				}
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
			if maxDistance.CompareAndSwap(globalMax, uint64(localMax)) {
				break
			}
		}

		wg.Done()
	}

	// for num workers
	for i := range numWorkers {
		//// add one to wg
		wg.Add(1)

		//// calculate start and end (i * chunkSize, (i+1)*chunkSize), if last chunk end = len(points)
		start, end := i*chunkSize, (i+1)*chunkSize
		if i >= numWorkers-1 {
			end = len(points.PointsX)
		}

		//// call the work closure
		go calcArea(start, end)
	}

	// wait for wg
	wg.Wait()

	fmt.Println(maxDistance.Load())
}
