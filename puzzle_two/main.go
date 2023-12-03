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

func main() {
	sum := summarizeFile("puzzle_2nd_part_input", lineHandler_2ndPuzzle)
	fmt.Println("Summarized value is: ", sum)
}