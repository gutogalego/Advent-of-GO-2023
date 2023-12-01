package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func getLines(filePath string) ([]string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Create a slice to hold the lines
	var lines []string

	// Read each line
	for scanner.Scan() {
		lines = append(lines, scanner.Text()) // Append the line to the slice
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func problem1() {
	lines, _ := getLines("input.txt")
	fmt.Println(lines)

	sum := 0

	for _, line := range lines {
		var start rune
		var end rune

		for _, char := range line {
			if unicode.IsDigit(char) {
				if start == 0 {
					start = char
				}
				end = char

			}
		}
		combinedString := string(start) + string(end)

		combinedInt, _ := strconv.Atoi(combinedString)
		sum += combinedInt
	}
	fmt.Println(sum)

}

func main() {
	problem1()
}
