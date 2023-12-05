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

func getPartsFromLine(prev string, curr string, next string) int {

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
		for !isAdjacent && i < len(prev) && i < end+1 {
			if i >= start-1 && isSymbol(prevR[i]) {
				isAdjacent = true
				break
			}
			i++
		}

		// check row below
		for !isAdjacent && j < len(next) && j < end+1 {
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

func getGearRatiosFromLine(prev string, curr string, next string) int {
	// If we find a gear, we check if there is a number:
	//   - to the left and right of the gear
	//   - above and below the gear

	re := regexp.MustCompile("[0-9]+")
	prevMatches := re.FindAllIndex([]byte(prev), -1)
	nextMatches := re.FindAllIndex([]byte(next), -1)

	j := 0 // index of prevMatch
	k := 0

	result := 0

	currR := []rune(curr)
	for i := 0; i < len(currR); i++ {
		if currR[i] == '*' {

			matches := []string{}

			// check the numbers in the upper row. Iterate until we find a number that starts after the current index + 1 because at
			// that point we know that the number is not adjacent to the gear.
			for j < len(prevMatches) && prevMatches[j][0] <= i+1 {
				if prevMatches[j][1] > i-1 {
					matches = append(matches, prev[prevMatches[j][0]:prevMatches[j][1]])
				}
				j++
			}

			for k < len(nextMatches) && nextMatches[k][0] <= i+1 {
				if nextMatches[k][1] > i-1 {
					matches = append(matches, next[nextMatches[k][0]:nextMatches[k][1]])
				}
				k++
			}

			// Not the most efficient solution

			// matches in the current row prefix.
			prevs := re.FindAllIndex([]byte(curr[0:i]), -1)
			// We need to check if the last match ends at the current index.
			if len(prevs) > 0 && prevs[len(prevs)-1][1] == i {
				matches = append(matches, curr[prevs[len(prevs)-1][0]:prevs[len(prevs)-1][1]])
			}

			// matches in the current row suffix.
			nexts := re.FindAllIndex([]byte(curr[i+1:]), -1)
			// We need to check if the first match starts at index 0 in the suffix.
			if len(nexts) > 0 && nexts[0][0] == 0 {
				matches = append(matches, curr[i+1:i+1+nexts[0][1]])
			}

			if len(matches) == 2 {
				num1, err1 := strconv.Atoi(matches[0])
				num2, err2 := strconv.Atoi(matches[1])
				if err1 != nil || err2 != nil {
					log.Fatal("Error converting string to integer")
				}
				result += num1 * num2
			}

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
	var partsSum int
	var gearSum int

	for scanner.Scan() {
		if prev == "" {
			prev = scanner.Text()
		} else if curr == "" {
			curr = scanner.Text()
			partsSum += getPartsFromLine("", prev, curr)
			gearSum += getGearRatiosFromLine("", prev, curr)
		} else {
			next = scanner.Text()
			partsSum += getPartsFromLine(prev, curr, next)
			gearSum += getGearRatiosFromLine(prev, curr, next)
			prev = curr
			curr = next
		}
	}

	partsSum += getPartsFromLine(prev, curr, "")
	gearSum += getGearRatiosFromLine(prev, curr, "")

	fmt.Println("Sum of parts: " + strconv.Itoa(partsSum))
	fmt.Println("Sum of gear ratios: " + strconv.Itoa(gearSum))
}
