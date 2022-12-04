package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	depthData, _ := os.ReadFile("../input.txt")
	depthString := strings.Split(string(depthData), "\n")

	depths := convertToInts(depthString)

	part1Results := part1(depths)
	part2Results := part2(depths)

	fmt.Printf("\npart 1 results: %v\n", part1Results)
	fmt.Printf("part 2 results: %v\n", part2Results)
}

func part1(depths []int) int {
	var depthIncrease, previousDepth int

	for _, depth := range depths {
		if previousDepth == 0 {
			previousDepth = depth

			continue
		}

		if depth > previousDepth {
			depthIncrease++
		}

		previousDepth = depth
	}

	return depthIncrease
}

func part2(depths []int) int {
	var counter int
	var previousSum int

	for i, depth := range depths {
		if previousSum == 0 {
			previousSum = depth
		}

		if i == len(depths)-3 {
			break
		}

		window := depths[i] + depths[i+1] + depths[i+2]

		if window > previousSum {
			counter++
		}

		previousSum = window
	}

	return counter
}

func convertToInts(list []string) []int {
	var intList []int

	for _, item := range list {
		n, _ := strconv.Atoi(item)
		intList = append(intList, n)
	}

	return intList
}
