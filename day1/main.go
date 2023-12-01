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

type Stack struct {
	elements []rune
	couter   int
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func problem2() {

	digitMap := map[string]int{
		"zero":  0,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	lines, _ := getLines("input.txt")
	fmt.Println(lines)

	prefixSet := make(map[string]struct{})
	for word := range digitMap {
		for i := 1; i <= len(word); i++ {
			prefixSet[word[:i]] = struct{}{}
		}
	}

	reversedPrefixSet := make(map[string]struct{})
	for word := range digitMap {
		reversedWord := reverseString(word)
		for i := 1; i <= len(reversedWord); i++ {
			reversedPrefixSet[reversedWord[:i]] = struct{}{}
		}
	}

	fmt.Println(prefixSet)
	fmt.Println(reversedPrefixSet)

	first := -1
	lastDigit := -1
	sum := 0

	for _, str := range lines {
		for start, end := 0, 1; start < len(str); {
			if unicode.IsDigit(rune(str[start])) {
				first = int(str[start] - '0')
				break
			}
			if _, exists := prefixSet[str[start:end]]; exists {
				if _, exists := digitMap[str[start:end]]; exists {
					first = digitMap[str[start:end]]
					break
				}
				if end < len(str) {
					end++
				} else {
					start++
				}
			} else {
				start++
				if end <= start {
					end = start + 1
				}
			}
		}

		str := reverseString(str)

		for start, end := 0, 1; start < len(str); {
			if unicode.IsDigit(rune(str[start])) {
				lastDigit = int(str[start] - '0')
				break
			}
			if _, exists := reversedPrefixSet[str[start:end]]; exists {
				if _, exists := digitMap[reverseString(str[start:end])]; exists {
					lastDigit = digitMap[reverseString(str[start:end])]
					break
				}
				if end < len(str) {
					end++
				} else {
					start++
				}
			} else {
				start++
				if end <= start {
					end = start + 1
				}
			}
		}

		total := first*10 + lastDigit
		sum += total
	}

	fmt.Println("Sum: ", sum)
}

func main() {
	problem2()
}
