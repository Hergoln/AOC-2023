package main

import (
	"fmt"
	"os"
	"bufio"
	"unicode"
	"strconv"
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
	sum := summarizeFile("puzzle_1st_part_input", lineHandler_1stPuzzle)
	fmt.Println("Summarized value is: ", sum)
}