package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("../input.txt")
	commandStrings := strings.Split(string(data), "\n")

	part1Results := part1(commandStrings)
	part2Results := part2(commandStrings)

	fmt.Printf("\npart 1 results: %v\n", part1Results)
	fmt.Printf("part 2 results: %v\n", part2Results)
}

func part1(list []string) int64 {

	type report struct {
		high int
		low  int
	}

	gamma := make(map[int]*report)
	epsilon := make(map[int]*report)

	for i := 0; i <= 11; i++ {
		gamma[i] = &report{}
		epsilon[i] = &report{}
	}

	for _, line := range list {
		for i, char := range line {
			switch string(char) {
			case "0":
				gamma[i].high += 1
				epsilon[i].high += 1
			case "1":
				gamma[i].low += 1
				epsilon[i].low += 1
			}
		}
	}

	var gammaString bytes.Buffer
	var epslionString bytes.Buffer
	for i := 0; i <= 11; i++ {
		if gamma[i].high > gamma[i].low {
			gammaString.WriteString(fmt.Sprintf("%d", 1))
			epslionString.WriteString(fmt.Sprintf("%d", 0))
		} else {
			epslionString.WriteString(fmt.Sprintf("%d", 1))
			gammaString.WriteString(fmt.Sprintf("%d", 0))

		}
	}

	gammaInt, _ := strconv.ParseInt(gammaString.String(), 2, 64)
	epsilonInt, _ := strconv.ParseInt(epslionString.String(), 2, 64)

	return gammaInt * epsilonInt
}

func part2(list []string) int {

	return len(list)
}
