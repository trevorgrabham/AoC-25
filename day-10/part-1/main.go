package main

import (
	"bytes"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"
)

type ProblemState struct {
	Distance       int
	PreviousButton int
}

type Problem struct {
	IndicatorLights     string
	ButtonWirings       [][]int
	JoltageRequirements []int
}

func (p *Problem) Solve() int {
	goal := strings.Repeat(".", len(p.IndicatorLights))

	maxStates := math.Pow(2, float64(len(p.IndicatorLights)))
	distances := make(map[string]ProblemState, int(maxStates))
	distances[p.IndicatorLights] = ProblemState{Distance: 0, PreviousButton: -1}

	stateQueue := make([]string, 0, int(maxStates)/2)
	stateQueue = append(stateQueue, p.IndicatorLights)
	for {
		currentState := stateQueue[0]
		current := distances[currentState]
		stateQueue = stateQueue[1:]
		for i, button := range p.ButtonWirings {
			if i == current.PreviousButton {
				continue
			}

			updatedState := currentState
			for _, wiring := range button {
				if updatedState[wiring] == '.' {
					updatedState = updatedState[:wiring] + "#" + updatedState[wiring+1:]
				} else {
					updatedState = updatedState[:wiring] + "." + updatedState[wiring+1:]
				}
			}

			if updatedState == goal {
				return current.Distance + 1
			}
			if updated, ok := distances[updatedState]; !ok || updated.Distance > current.Distance + 1 {
				distances[updatedState] = ProblemState{Distance: current.Distance + 1, PreviousButton: i}
				stateQueue = append(stateQueue, updatedState)
			} 
		}
	}
}

func (p Problem) String() string {
	var s strings.Builder
	s.WriteByte('[')
	s.WriteString(p.IndicatorLights)
	s.WriteString("] ")
	for _, wiring := range p.ButtonWirings {
		s.WriteByte('(')
		for i, light := range wiring {
			s.WriteString(strconv.Itoa(light))
			if i == len(wiring)-1 {
				continue
			}
			s.WriteByte(',')
		}
		s.WriteString(") ")
	}
	s.WriteByte('{')
	for i, req := range p.JoltageRequirements {
		s.WriteString(strconv.Itoa(req))
		if i == len(p.JoltageRequirements)-1 {
			continue
		}
		s.WriteByte(',')
	}
	s.WriteByte('}')

	return s.String()
}

func main() {
	inputData, err := os.ReadFile("../input/challenge.txt")
	if err != nil {
		log.Fatal(err)
	}
	inputData = bytes.TrimSpace(inputData)

	var (
		totalSum              atomic.Uint64
		wg                    sync.WaitGroup
		numWorkers, chunkSize int
		problems              []Problem
	)

	lines := bytes.Split(inputData, []byte("\n"))
	problems = make([]Problem, 0, len(lines))

	for _, line := range lines {
		indicatorLightsIndex := bytes.IndexByte(line, ']')
		if indicatorLightsIndex < 0 {
			panic(fmt.Errorf("unable to locate indicator lights: %s", string(line)))
		}
		joltageRequirementsIndex := bytes.IndexByte(line, '{')
		if joltageRequirementsIndex < 0 {
			panic(fmt.Errorf("unable to locate joltage requirements: %s", string(line)))
		}

		var p Problem
		p.IndicatorLights = string(line[1:indicatorLightsIndex:indicatorLightsIndex])

		buttonWirings := bytes.Split(bytes.TrimSpace(line[indicatorLightsIndex+1:joltageRequirementsIndex:joltageRequirementsIndex]), []byte(" "))
		p.ButtonWirings = make([][]int, 0, len(buttonWirings))
		for _, wiring := range buttonWirings {
			wiring = wiring[1 : len(wiring)-1 : len(wiring)-1]
			affectedButtons := bytes.Split(wiring, []byte(","))
			buttons := make([]int, 0, len(affectedButtons))
			for _, button := range affectedButtons {
				buttonNumber, err := strconv.Atoi(string(button))
				if err != nil {
					panic(fmt.Errorf("parsing button number: %v", err))
				}

				buttons = append(buttons, buttonNumber)
			}
			p.ButtonWirings = append(p.ButtonWirings, buttons)
		}

		joltageRequirements := bytes.Split(line[joltageRequirementsIndex+1:len(line)-1:len(line)-1], []byte(","))
		p.JoltageRequirements = make([]int, 0, len(joltageRequirements))
		for _, joltageRequired := range joltageRequirements {
			req, err := strconv.Atoi(string(joltageRequired))
			if err != nil {
				panic(fmt.Errorf("parsing joltage requirement: %v", err))
			}

			p.JoltageRequirements = append(p.JoltageRequirements, req)
		}

		problems = append(problems, p)
	}

	numWorkers = min(runtime.NumCPU(), len(problems))
	chunkSize = int(math.Round(float64(len(problems)) / float64(numWorkers)))

	for i := range numWorkers {
		wg.Add(1)

		start, end := i*chunkSize, (i+1)*chunkSize
		if i == numWorkers-1 {
			end = len(problems)
		}
		fmt.Printf("thread %d:\t[%d, %d)\n", i, start, end)

		go func(start, end int) {
			problems := problems[start:end:end]
			for _, problem := range problems {
				solution := problem.Solve()
				totalSum.Add(uint64(solution))
				fmt.Printf("thread %d: %s\nsolution: %d\n\n", i, problem.String(), solution)
			}
			wg.Done()
		}(start, end)
	}

	wg.Wait()

	fmt.Println(totalSum.Load())
}

/*
	split data by line
	split line into indicator lights, button wirings, and joltage requirements
	append data to problems
	prepare the worker pool
	for each worker
		prepare a map (indicator light state -> *distance from goal AND previous button pressed)
		prepare a queue for states to investigate and add the goal state
		forever
			grab a current state
			for each button (excluding the previously pressed one)
				calculate the updated state
				if the updated state is the start state (all off)
					add the current state length + 1 to our sum
				if updated state is new (since we are
					update the record for the updated state
					add the updated state to the queue
*/
