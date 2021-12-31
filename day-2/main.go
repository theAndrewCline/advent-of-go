package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseCourse() []string {
	fileContents, err := os.ReadFile("./input.txt")
	if err != nil {
		panic("Cannot read file")
	}

	fileString := strings.TrimSpace(string(fileContents))

	return strings.Split(fileString, "\n")

}

func main() {
	currentCourse := parseCourse()

	horizontalPosition := 0
	depth := 0
	aim := 0

	for _, direction := range currentCourse {
		parsedTuple := strings.Fields(direction)

		switch parsedTuple[0] {
		case "forward":
			distance, err := strconv.Atoi(parsedTuple[1])
			if err != nil {
				panic(err)
			}
			horizontalPosition += distance
			depth += aim * distance
		case "down":
			distance, err := strconv.Atoi(parsedTuple[1])
			if err != nil {
				panic(err)
			}
			aim += distance
		case "up":
			distance, err := strconv.Atoi(parsedTuple[1])
			if err != nil {
				panic(err)
			}
			aim -= distance
		}
	}

	fmt.Println(horizontalPosition * depth)
}
