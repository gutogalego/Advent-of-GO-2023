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

var directions = [][]int{
	{1, 0},
	{1, 1},
	{1, -1},
	{0, 1},
	{0, -1},
	{-1, 0},
	{-1, 1},
	{-1, -1},
}

func problem0() {
	lines, _ := getLines("example.txt")
	fmt.Println(lines)

	for idy, line := range lines {

		for idx, ch := range line {
			if unicode.IsNumber(ch) && ch != '.' {
				for _, dir := range directions {
					dir_y := 0
					if idy+dir[0] > 0 {
						dir_y = idy + dir[0]
						if dir_y > len(lines) {
							dir_y = len(lines) - 1
						}
					}

					dir_x := 0
					if idx+dir[1] > 0 {
						dir_x = idx + dir[1]
						if dir_x > len(line) {
							dir_x = len(line) - 1
						}
					}

					candidate_ch := rune(lines[dir_x][dir_y])

					if unicode.IsNumber(candidate_ch) {
						fmt.Printf("%c is a number\n", candidate_ch)
					}

				}
			}
		}
	}
}

func problem1() {
	lines, _ := getLines("input.txt")
	fmt.Println(lines)

	sum := 0

	idy := 0
	for idy < len(lines) {
		line := lines[idy]

		idx := 0
		num_start := -1
		num_end := -1
		for idx < len(line) {

			ch := line[idx]
			if unicode.IsNumber(rune(ch)) {
				if num_start == -1 {
					num_start = idx
				}
				num_end = idx
			}

			if (!unicode.IsNumber(rune(ch)) || idx+1 == len(line)) && num_start != -1 {

				valid := false
				for i := num_start; i <= num_end; i++ {

					for _, dir := range directions {
						dir_y := 0
						if idy+dir[0] > 0 {
							dir_y = idy + dir[0]
							if dir_y >= len(lines) {
								dir_y = len(lines) - 1
							}
						}

						dir_x := 0
						if i+dir[1] > 0 {
							dir_x = i + dir[1]
							if dir_x >= len(line) {
								dir_x = len(line) - 1
							}
						}

						candidate_ch := rune(lines[dir_y][dir_x])
						if !unicode.IsNumber(candidate_ch) && candidate_ch != '.' {
							valid = true
						}
					}

				}
				if valid {
					substr := line[num_start : num_end+1] // Substring from position 1 to 3

					// Convert substring to number
					num, _ := strconv.Atoi(substr)

					if num > 700 {
						fmt.Println(substr)
					}
					sum += num
				}

				num_start = -1
				num_end = -1
			}
			idx++
		}
		idy++
	}
	fmt.Println("The value of sum is:", sum)
}

func problem2() {
	lines, _ := getLines("input.txt")
	fmt.Println(lines)

	sum := 0

	idy := 0
	for idy < len(lines) {
		line := lines[idy]

		idx := 0
		num_start := -1
		num_end := -1
		for idx < len(line) {

			ch := line[idx]
			if unicode.IsNumber(rune(ch)) {
				if num_start == -1 {
					num_start = idx
				}
				num_end = idx
			}

			if (!unicode.IsNumber(rune(ch)) || idx+1 == len(line)) && num_start != -1 {

				valid := false
				for i := num_start; i <= num_end; i++ {

					for _, dir := range directions {
						dir_y := 0
						if idy+dir[0] > 0 {
							dir_y = idy + dir[0]
							if dir_y >= len(lines) {
								dir_y = len(lines) - 1
							}
						}

						dir_x := 0
						if i+dir[1] > 0 {
							dir_x = i + dir[1]
							if dir_x >= len(line) {
								dir_x = len(line) - 1
							}
						}

						candidate_ch := rune(lines[dir_y][dir_x])
						if !unicode.IsNumber(candidate_ch) && candidate_ch != '.' {
							valid = true
						}
					}

				}
				if valid {
					substr := line[num_start : num_end+1] // Substring from position 1 to 3

					// Convert substring to number
					num, _ := strconv.Atoi(substr)

					if num > 700 {
						fmt.Println(substr)
					}
					sum += num
				}

				num_start = -1
				num_end = -1
			}
			idx++
		}
		idy++
	}
	fmt.Println("The value of sum is:", sum)
}

func main() {
	problem2()
}
