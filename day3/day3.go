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

func removeDuplicates(elements []int) []int {
	// Map to keep track of unique elements
	seen := make(map[int]bool)

	// Slice to hold the unique elements
	unique := []int{}

	for _, value := range elements {
		if _, ok := seen[value]; !ok {
			seen[value] = true
			unique = append(unique, value)
		}
	}

	return unique
}

func reverseRunes(runes []rune) []rune {
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return runes
}

func problem2() {
	lines, _ := getLines("input.txt")
	fmt.Println(lines)

	sum := 0

	idy := 0
	for idy < len(lines) {
		line := lines[idy]

		idx := 0

		for idx < len(line) {

			ch := line[idx]

			if ch == '*' {
				nums := []int{}

				for _, dir := range directions {
					dir_y := 0
					if idy+dir[0] > 0 {
						dir_y = idy + dir[0]
						if dir_y >= len(lines) {
							dir_y = len(lines) - 1
						}
					}

					dir_x := 0
					if idx+dir[1] > 0 {
						dir_x = idx + dir[1]
						if dir_x >= len(line) {
							dir_x = len(line) - 1
						}
					}

					candidate_ch := rune(lines[dir_y][dir_x])
					if unicode.IsNumber(candidate_ch) {
						fmt.Printf("%c", candidate_ch)
						rightNumeric := true
						leftNumeric := true
						dir_right := dir_x + 1
						dir_left := dir_x - 1
						rightChars := []rune{}
						leftChars := []rune{}

						for dir_left >= 0 && leftNumeric {
							left_ch := rune(lines[dir_y][dir_left])
							dir_left--
							if unicode.IsNumber(left_ch) {
								leftChars = append(leftChars, left_ch)
							} else {
								leftNumeric = false
							}
						}

						for dir_right < len(line) && rightNumeric {
							right_ch := rune(lines[dir_y][dir_right])
							dir_right++
							if unicode.IsNumber(right_ch) {
								rightChars = append(rightChars, right_ch)
							} else {
								rightNumeric = false
							}

						}

						// fmt.Println("left", leftChars)
						// fmt.Println("candidate_ch", candidate_ch)
						// fmt.Println("rightChars", rightChars)
						leftChars = reverseRunes(leftChars)
						allElements := append(leftChars, candidate_ch)
						allElements = append(allElements, rightChars...)
						resultString := string(allElements)

						resultInt, _ := strconv.Atoi(resultString)

						nums = append(nums, resultInt)

					}

				}
				nums = removeDuplicates(nums)

				if len(nums) == 2 {
					fmt.Println("nums", nums)
					sum += nums[0] * nums[1]
				}

				// Identify numbers around

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
