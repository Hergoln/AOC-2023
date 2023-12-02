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

func summarizeFile_1stPuzzle(fileName string) int {
	file, err := os.Open(fileName)
	checkError(err)
	stream := bufio.NewScanner(file)
	sum := 0

	for stream.Scan() {
		firstNumber := -1
		lastNumber := -1
		text := stream.Text()
		for _, c := range text {
			if unicode.IsDigit(c) {
				if firstNumber < 0 {
					firstNumber = int(c)
				}
				lastNumber = int(c)
			}
		}

		intVar, _ := strconv.Atoi(string(firstNumber) + string(lastNumber))
		sum = sum + intVar
	}
	checkError(stream.Err())
	return sum
}

func

func main() {
	sum := summarizeFile_1stPuzzle("puzzle_1st_part_input")
	fmt.Println("Summarized value is: ", sum)
}