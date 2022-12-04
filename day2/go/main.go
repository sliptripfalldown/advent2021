package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("../input.txt")
	commands := strings.Split(string(data), "\n")

	part1Results := part1(commands)
	part2Results := part2(commands)

	fmt.Printf("\npart 1 results: %v\n", part1Results)
	fmt.Printf("part 2 results: %v\n", part2Results)
}

func part1(commands []string) int {
	var horizontal, depth int

	for _, move := range commands {
		cmd := strings.Split(move, " ")
		action := cmd[0]
		value, _ := strconv.Atoi(cmd[1])

		switch action {
		case "forward":
			horizontal += value
		case "down":
			depth += value
		case "up":
			depth -= value
		}
	}

	return horizontal * depth
}

func part2(commands []string) int {
	var horizontal, depth, aim int

	for _, move := range commands {
		cmd := strings.Split(move, " ")
		action := cmd[0]
		value, _ := strconv.Atoi(cmd[1])

		switch action {
		case "forward":
			horizontal += value
			depth += aim * value
		case "down":
			aim += value
		case "up":
			aim -= value
		}
	}

	return horizontal * depth
}
