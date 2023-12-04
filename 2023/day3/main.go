package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

// The goal is to include any number in the sum if it is adjacent to a symbol, even diagonally.
//
// The input file contains various symbols such as $, +, =, -, etc.
// We can visualize the input as a 2D matrix of characters:
// Each symbol can be represented by its row and column number, e.g., [1,5].
// Each number can be represented by its row number and a column range (since numbers can be multiple digits long), e.g., [2, 3-7].
// For a given number [2, 3-7], we need to check:
//   - If there is a symbol adjacent on the left or right, e.g., at [2, 2] or [2, 8]?
//   - If there is a symbol adjacent above or below, e.g., at [1, 2-8] or [3, 2-8]?
// To compute this efficiently, we have two options:
// 1. Scan three rows at a time. Read the first and last row to identify the symbols, and read the middle row to identify the numbers. Check if these numbers should be included in the sum.
// 2. Alternatively, we can scan three columns at a time, identify symbols and note any numeric digits adjacent to them. Then, scan through the input to find the numbers that contain these numeric digits. This approach might be less efficient but simpler.
// I will proceed with option 1, but  will scan all three rows simultaneously instead of one at a time.

func isSymbol(c rune) bool {
	return c != '.' && (c < '0' || c > '9')
}

func processLines(prev string, curr string, next string) int {

	prevR := []rune(prev)
	currR := []rune(curr)
	nextR := []rune(next)

	re := regexp.MustCompile("[0-9]+")
	matches := re.FindAllIndex([]byte(curr), -1)

	currr := []rune(curr)

	i := 0
	j := 0
	var result int

	for _, match := range matches {
		start := match[0]
		end := match[1]
		number, _ := strconv.Atoi(curr[start:end])

		// check left and right
		isAdjacent := start > 0 && isSymbol(currR[start-1]) || end < len(currr) && isSymbol(currR[end])

		// check row above
		for isAdjacent == false && i < len(prev) && i < end+1 {
			if i >= start-1 && isSymbol(prevR[i]) {
				isAdjacent = true
				break
			}
			i++
		}

		// check row below
		for isAdjacent == false && j < len(next) && j < end+1 {
			if j >= start-1 && isSymbol(nextR[j]) {
				isAdjacent = true
				break
			}
			j++
		}

		if isAdjacent {
			result += number
		}
	}

	return result
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	var prev, curr, next string
	var result int

	for scanner.Scan() {
		if prev == "" {
			prev = scanner.Text()
		} else if curr == "" {
			curr = scanner.Text()
			result += processLines("", prev, curr)
		} else {
			next = scanner.Text()
			result += processLines(prev, curr, next)
			prev = curr
			curr = next
		}
	}

	result += processLines(prev, curr, "")

	fmt.Println(result)
}
