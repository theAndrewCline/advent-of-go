package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
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

	timesDepthIncreased := 0

	for i, depth := range depths {
		if i != 0 {
			prevDepth := depths[i-1]

			if depth > prevDepth {
				timesDepthIncreased++
			}

		}
	}

	fmt.Println(timesDepthIncreased)
}
