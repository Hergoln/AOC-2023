package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

/*
1st Part
	Each line contains a set of multiple takes within a game.
	A game is to determine if given the number of cubes within each take it is possible to have
	a set of cubes like determined in config variable (12 red, 13 green, 14 blue).
	If it is possible that this confguration of cubes can be in the bag then return the number
	of the id of the game which is given at the beginning of the line ("Game <ID>:...").
	Sum up the IDs of possible games, this sum is answer for this puzzle.
	Example line from file:
	Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red

2nd Part
	Each game looks the exact same as before. This time we need to find
	the minimal set of cubes that would be possible for each game.
	Then for each game we multiple number of red, blue and green cubes and 
	sum the result for all games.
*/

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func retrieveStreamFromFile(filename string) *bufio.Scanner {
	file, err := os.Open(filename)
	checkError(err)
	return bufio.NewScanner(file)
}

var (
	config = map[string]int {
		"red": 12,
		"green": 13,
		"blue": 14,
	}
	gamePrefixLen = len("Game ")
)

func getNextTake(take string) string {
	sIndex := strings.Index(take, ";")
	if sIndex == -1 {
		return take
	}
	return take[:sIndex]
}

func isTakePossible(take string) bool {
	color := ""
	for len(take) > 0 {
		space := strings.Index(take[1:], " ")
		val, _ := strconv.Atoi(take[1:space+1])
		take = take[space+1:]

		commaIndex := strings.Index(take, ",")
		if commaIndex == -1 {
			color = take[1:]
			take = ""
		} else {
			color = take[1:commaIndex]
			take = take[commaIndex+1:]
		}
		if val > config[color] {
			
			return false
		}
	}
	return true
}

func isLinePossible(line string) bool {
	for len(line) > 0 {
		take := getNextTake(line[1:])
		if len(take) <= 0 {
			break
		}
		if !isTakePossible(take) {
			return false
		}
		line = line[len(take)+1:]
	}
	return true
}

func getGameIndex(line string) (int, int) {
	sIndex := strings.Index(line, ":")
	val, _ := strconv.Atoi(line[gamePrefixLen:sIndex])
	return val, sIndex
}

func calculateTake(take string, currentMax map[string]int) map[string]int {
	color := ""
	for len(take) > 0 {
		space := strings.Index(take[1:], " ")
		val, _ := strconv.Atoi(take[1:space+1])
		take = take[space+1:]

		commaIndex := strings.Index(take, ",")
		if commaIndex == -1 {
			color = take[1:]
			take = ""
		} else {
			color = take[1:commaIndex]
			take = take[commaIndex+1:]
		}
		if val > currentMax[color] {
			currentMax[color] = val
		}
	}
	return currentMax
}

func gameResult(line string) int {
	currentMax := map[string]int{
		"green": 1,
		"red": 1,
		"blue": 1,
	}
	for len(line) > 0 {
		take := getNextTake(line[1:])
		currentMax = calculateTake(take, currentMax)
		line = line[len(take)+1:]
	}
	return currentMax["green"] * currentMax["blue"] * currentMax["red"]
}

func firstPart() {
	stream := retrieveStreamFromFile("puzzle_input")
	sum := 0
	for stream.Scan() {
		line := stream.Text()
		gameId, prefixLen := getGameIndex(line)
		if isLinePossible(line[prefixLen:]) {
			sum = sum + gameId
		}
	}
	fmt.Println(sum)
}

func secondPart() {
	stream := retrieveStreamFromFile("puzzle_input")
	sum := 0
	for stream.Scan() {
		line := stream.Text()
		_, prefixLen := getGameIndex(line)
		sum = sum + gameResult(line[prefixLen:])
	}
	fmt.Println(sum)
}

func main() {
	secondPart()
}