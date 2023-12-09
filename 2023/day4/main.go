package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	re := regexp.MustCompile(`\d+`)

	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	var result1 float64 = 0

	scanner := bufio.NewScanner(inputFile)

	cardPoints := make(map[int]int)

	for scanner.Scan() {
		curr := scanner.Text()
		// assuming no duplicates in the winning list and numbers I have.
		matches := re.FindAllString(curr, -1)

		presence := make(map[string]bool)

		var winning float64 = 0

		// skip the first number which is the card.
		for _, match := range matches[1:] {
			if presence[match] {
				winning++
			}
			presence[match] = true
		}

		if winning > 0 {
			result1 += math.Pow(2, winning-1)
		}

		card, _ := strconv.Atoi(matches[0])

		if _, exists := cardPoints[card]; !exists {
			cardPoints[card] = 1 // set current card.
		} else {
			cardPoints[card]++
		}

		for i := card + 1; i <= card+int(winning); i++ {
			if _, exists := cardPoints[i]; !exists {
				cardPoints[i] = 0
			}
			cardPoints[i] += cardPoints[card]
		}
	}

	fmt.Println("Part 1 Total Points: ", result1)

	result2 := 0

	for _, v := range cardPoints {
		result2 += v
	}
	fmt.Println("Part 2 Total Points: ", result2)
}
