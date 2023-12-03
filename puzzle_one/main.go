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

var (
	maxLenOfValue int = 5
	numberSet = map[string]int{
		"one": 1,
		"two": 2,
		"three": 3,
		"four": 4,
		"five": 5,
		"six" : 6,
		"seven": 7,
		"eight": 8,
		"nine": 9,
	}
)

func extractNumber(s string) int {
	for i := 0; i < len(s); i++ {
		for k, v := range numberSet {
			if strings.Contains(s[0:i+1], k) {
				return v
			}
		}
		if unicode.IsDigit(rune(s[i])) {
			val, _ := strconv.Atoi(string(s[i]))
			return val
		}
	}
	return 0
}

func computeEndOfSlice(lenOfLine, currentIndex, maxLen int) int {
	if currentIndex + maxLen > lenOfLine {
		return lenOfLine
	}
	return currentIndex + maxLen
}

func lineHandler_2ndPuzzle(line string) int {
	firstNumber := 0
	lastNumber := 0
	for i := 0; i < len(line); i++ {
		endOfSlice := computeEndOfSlice(len(line), i, maxLenOfValue)
		if firstNumber <= 0 {
			number := extractNumber(line[i:endOfSlice])
			if number > 0 {
				firstNumber = number
				lastNumber = number
			}
		} else {
			number := extractNumber(line[i:endOfSlice])
			if number > 0 {
				lastNumber = number
			}
		}
	}
	val, _ := strconv.Atoi(strconv.Itoa(firstNumber) + strconv.Itoa(lastNumber))
	return val
}

func summarizeFile(fileName string, handler handlerFunc) int {
	file, err := os.Open(fileName)
	checkError(err)
	stream := bufio.NewScanner(file)
	sum := 0

	for stream.Scan() {
		sum = sum + handler(stream.Text())
	}
	checkError(stream.Err())
	return sum
}

func main() {
	sum := summarizeFile("puzzle_2nd_part_input", lineHandler_2ndPuzzle)
	fmt.Println("Summarized value is: ", sum)
}