package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	dat, err := os.ReadFile("./test.txt")
	if err != nil {
		panic(err)
	}

	fileStr := string(dat)

	depthStrings := strings.Split(fileStr, "/n")

	depths := make([]int, len(depthStrings)-1)
	for i, ds := range depthStrings {
		if i != 0 {
			parsedDepth, err := strconv.Atoi(ds)

			if err != nil {
				panic(err)
			}

			depths[i-1] = parsedDepth
		}
	}
}
