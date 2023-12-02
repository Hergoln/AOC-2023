package main

import (
	"fmt"
	"os"
	"bufio"
	"unicode"
	"strconv"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type handlerFunc func(string) int

func lineHandler_1stPuzzle(line string) int {
	firstNumber := -1
	lastNumber := -1
	for _, c := range line {
		if unicode.IsDigit(c) {
			if firstNumber < 0 {
				firstNumber = int(c)
			}
			lastNumber = int(c)
		}
	}

	intVar, _ := strconv.Atoi(string(firstNumber) + string(lastNumber))
	return intVar
}

maxLenOfValue := 5
numberSet := map[string]int{
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5
	"six" : 6,
	"seven": 7,
	"eight": 8,
	"nine": 9
}

func extractNumber(s string) (int, int) {
	if s[0].IsDigit {
		return strconv.Atoi(s[0]), 1
	}
	for k, v := range numberSet {
		if strings.Contains(s, v) {
			return v, len(s)
		}
	}
	return -1, 1
}

func lineHandler_2ndPuzzle(line string) int {
	firstNumber := -1
	lastNumber := -1
	for i := 0; i < len(line); {
		if firstNumber < 0 {
			number, len := extractNumber(line[i:maxLenOfValue])
			if number != 1 {
				firstNumber = number
				lastNumber = number
			}
			i = i + len
		} else {
			lastNumber, len = extractNumber(line[i:maxLenOfValue])
			i = i + len
		}
	}
	return strconv.Atoi(string(firstNumber) + string(lastNumber))
}

func summarizeFile(fileName string, handler handlerFunc) int {
	file, err := os.Open(fileName)
	checkError(err)
	stream := bufio.NewScanner(file)
	sum := 0

	for stream.Scan() {
		sum = sum + lineHandler_1stPuzzle(stream.Text())
	}
	checkError(stream.Err())
	return sum
}

func main() {
	sum := summarizeFile("example_input_part2", lineHandler_2ndPuzzle)
	fmt.Println("Summarized value is: ", sum)
}