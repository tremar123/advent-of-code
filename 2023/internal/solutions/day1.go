package solutions

import (
	"strconv"
	"strings"
	"unicode"
)

func Day1Part1(input string) any {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		firstDigit := ""
		lastDigit := ""
		for _, ch := range line {
			if unicode.IsDigit(ch) {
				if firstDigit == "" {
					firstDigit = string(ch)
				}
				lastDigit = string(ch)
			}
		}
		value, _ := strconv.Atoi(firstDigit + lastDigit)
		sum += value
	}

	return sum
}

func Day1Part2(input string) any {
	lines := strings.Split(input, "\n")

	sum := 0
	for _, line := range lines {
		firstDigit := ""
		lastDigit := ""
		for runeOnLineIndex := 0; runeOnLineIndex < len(line); runeOnLineIndex++ {
			words := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
			if unicode.IsDigit(rune(line[runeOnLineIndex])) {
				if firstDigit == "" {
					firstDigit = string(line[runeOnLineIndex])
				}
				lastDigit = string(line[runeOnLineIndex])
			} else {
				for _, word := range words {
					endOfWordIndex := runeOnLineIndex + len(word)
					if endOfWordIndex > len(line) {
						continue
					}
					str := line[runeOnLineIndex:endOfWordIndex]
					switch str {
					case "one":
						if firstDigit == "" {
							firstDigit = "1"
						}
						lastDigit = "1"
						break
					case "two":
						if firstDigit == "" {
							firstDigit = "2"
						}
						lastDigit = "2"
						break
					case "three":
						if firstDigit == "" {
							firstDigit = "3"
						}
						lastDigit = "3"
						break
					case "four":
						if firstDigit == "" {
							firstDigit = "4"
						}
						lastDigit = "4"
						break
					case "five":
						if firstDigit == "" {
							firstDigit = "5"
						}
						lastDigit = "5"
						break
					case "six":
						if firstDigit == "" {
							firstDigit = "6"
						}
						lastDigit = "6"
						break
					case "seven":
						if firstDigit == "" {
							firstDigit = "7"
						}
						lastDigit = "7"
						break
					case "eight":
						if firstDigit == "" {
							firstDigit = "8"
						}
						lastDigit = "8"
						break
					case "nine":
						if firstDigit == "" {
							firstDigit = "9"
						}
						lastDigit = "9"
						break
					}
				}
			}
		}
		value, _ := strconv.Atoi(firstDigit + lastDigit)
		sum += value
	}

	return sum
}
