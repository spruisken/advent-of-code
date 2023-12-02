package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func extractCalibrationValues(s string) int {
	// the problem statement wasn't clear as to whether
	// first and/or last could be 0 but I'm assuming they can't
	// I also didn't find any examples where they were 0 in the input file.
	var first, last int = 0, 0

	for _, c := range s {
		value, err := strconv.Atoi(string(c))
		if err == nil {
			if first == 0 {
				first = value
				last = first
			} else {
				last = value
			}
		}
	}

	return first*10 + last
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var sum int = 0

	for scanner.Scan() {
		line := scanner.Text()
		sum += extractCalibrationValues(line)
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println(sum)
}
