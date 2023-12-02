package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	ID     int
	Rounds []Round
}

type Round struct {
	Red   int
	Green int
	Blue  int
}

type Bag = Round

func parseGame(line string) Game {
	re := regexp.MustCompile(`Game (\d+): (.*)`)
	matches := re.FindStringSubmatch(line)
	id, _ := strconv.Atoi(matches[1])
	roundsStr := strings.Split(matches[2], ";")

	var rounds []Round

	for _, roundStr := range roundsStr {
		round := Round{}
		for _, colorStr := range strings.Split(roundStr, ",") {
			colorStr = strings.TrimSpace(colorStr)
			parts := strings.Split(colorStr, " ")
			count, _ := strconv.Atoi(parts[0])
			switch parts[1] {
			case "red":
				round.Red = count
			case "green":
				round.Green = count
			case "blue":
				round.Blue = count
			}
		}
		rounds = append(rounds, round)
	}

	return Game{
		ID:     id,
		Rounds: rounds,
	}
}

func getIdIfPossible(bag Bag, game Game) int {
	for _, round := range game.Rounds {
		if round.Red > bag.Red || round.Green > bag.Green || round.Blue > bag.Blue {
			return 0
		}
	}
	return game.ID
}

func getFewestCubesPower(bag Bag, game Game) int {
	fewestCubeBag := Bag{0, 0, 0}
	for _, round := range game.Rounds {
		fewestCubeBag.Red = max(round.Red, fewestCubeBag.Red)
		fewestCubeBag.Green = max(round.Green, fewestCubeBag.Green)
		fewestCubeBag.Blue = max(round.Blue, fewestCubeBag.Blue)
	}
	return fewestCubeBag.Blue * fewestCubeBag.Green * fewestCubeBag.Red
}

func main() {
	inputFile, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer inputFile.Close()

	scanner := bufio.NewScanner(inputFile)

	bag := Bag{Red: 12, Green: 13, Blue: 14}

	possibleGamesIdSum := 0
	fewestCubesPowerSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		game := parseGame(line)
		possibleGamesIdSum += getIdIfPossible(bag, game)
		fewestCubesPowerSum += getFewestCubesPower(bag, game)
	}

	fmt.Println("Sum of ids of possible games: " + strconv.Itoa(possibleGamesIdSum))
	fmt.Println("Sum of fewest cubes power: " + strconv.Itoa(fewestCubesPowerSum))
}
