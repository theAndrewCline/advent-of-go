package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseFileForDepths() []int {
	dat, err := os.ReadFile("depths.txt")
	if err != nil {
		panic(err)
	}

	fileStr := strings.TrimSpace(string(dat))

	depthStrings := strings.Split(fileStr, "\n")

	depths := make([]int, len(depthStrings))
	for i, ds := range depthStrings {
		if i != 0 {
			parsedDepth, err := strconv.Atoi(ds)
			if err != nil {
				panic(err)
			}

			depths[i] = parsedDepth
		}
	}

	return depths
}

func sumSet(depths []int, depth int, index int) int {
	return depth + depths[index+1] + depths[index+2]
}

func main() {
	timesDepthIncreased := 0

	depths := parseFileForDepths()

	for i, depth := range depths {
		if i+3 < len(depths) && i != 0 {
			currentSet := sumSet(depths, depth, i)
			previousSet := sumSet(depths, depths[i-1], i-1)

			if currentSet > previousSet {
				timesDepthIncreased++
			}
		}
	}

	fmt.Println(timesDepthIncreased)
}
