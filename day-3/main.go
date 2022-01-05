package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reportBytes, err := os.ReadFile("./power-consumption-report.txt")
	if err != nil {
		panic("Cannot read report")
	}
	report := string(reportBytes)
	reportNumbers := strings.Fields(string(report))

	gammaBinarySlice := make([]string, 12)

	for pos := range gammaBinarySlice {
		zeros := 0
		ones := 0

		for _, binNum := range reportNumbers {
			if binNum[pos] == '0' {
				zeros++
			} else {
				ones++
			}
		}

		if zeros > ones {
			gammaBinarySlice[pos] = "0"
		} else {
			gammaBinarySlice[pos] = "1"
		}
	}

	epsilonBinarySlice := make([]string, 12)

	for pos := range epsilonBinarySlice {
		zeros := 0
		ones := 0

		for _, binNum := range reportNumbers {
			if binNum[pos] == '0' {
				zeros++
			} else {
				ones++
			}
		}

		if zeros < ones {
			epsilonBinarySlice[pos] = "0"
		} else {
			epsilonBinarySlice[pos] = "1"
		}
	}

	gammaRate, err := strconv.ParseInt(strings.Join(gammaBinarySlice, ""), 2, 64)
	if err != nil {
		panic("could not parse binary number")
	}

	epsilonRate, err := strconv.ParseInt(strings.Join(epsilonBinarySlice, ""), 2, 64)
	if err != nil {
		panic("could not parse binary number")
	}

	fmt.Println(gammaRate * epsilonRate)
}
