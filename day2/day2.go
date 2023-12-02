package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

	diceMap := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	lines, _ := getLines("input.txt")
	fmt.Println(lines)

	sum := 0

outerLoop:
	for id, line := range lines {
		line = strings.Split(line, ":")[1]
		semis := strings.Split(line, ";")
		for _, semi := range semis {
			diceNums := strings.Split(semi, ",")
			for _, diceNum := range diceNums {
				diceNum = strings.TrimSpace(diceNum)
				diceAndNum := strings.Split(diceNum, " ")
				dice := diceAndNum[1]
				num, _ := strconv.Atoi(diceAndNum[0])

				max := diceMap[dice]

				if num > max {
					continue outerLoop
				}
			}
		}
		sum += id + 1
	}

	fmt.Println("sum: ", sum)

}

func problem2() {

	diceMap := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	lines, _ := getLines("input.txt")
	fmt.Println(lines)

	sum := 0

	for _, line := range lines {
		line = strings.Split(line, ":")[1]
		semis := strings.Split(line, ";")
		for _, semi := range semis {
			diceNums := strings.Split(semi, ",")
			for _, diceNum := range diceNums {
				diceNum = strings.TrimSpace(diceNum)
				diceAndNum := strings.Split(diceNum, " ")
				dice := diceAndNum[1]
				num, _ := strconv.Atoi(diceAndNum[0])

				max := diceMap[dice]

				if num > max {
					diceMap[dice] = num
				}
			}
		}
		sum += diceMap["red"] * diceMap["green"] * diceMap["blue"]

		diceMap = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
	}

	fmt.Println("sum: ", sum)
}

func main() {
	problem2()
}
