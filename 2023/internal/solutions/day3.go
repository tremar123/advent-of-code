package solutions

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func Day3Part1(input string) any {
	lines := strings.Split(input, "\n")

	sum := 0
	currentNumber := ""
	currentNumberValid := false
	for lineIndex, line := range lines {
		lineOffsetMin := -1
		lineOffsetMax := 1
		if lineIndex == 0 {
			lineOffsetMin = 0
		}

		// 2 because strings.Split on \n return empty string as last element
		if lineIndex == len(lines)-2 {
			lineOffsetMax = 0
		}

		for charIndex, char := range line {
			if unicode.IsDigit(char) {
				currentNumber += string(char)

				if currentNumberValid {
					continue
				}

				charOffsetMin := -1
				charOffsetMax := 1

				if charIndex == 0 {
					charOffsetMin = 0
				}

				if charIndex == len(line)-1 {
					charOffsetMax = 0
				}

				for i := lineOffsetMin; i <= lineOffsetMax; i++ {
					for j := charOffsetMin; j <= charOffsetMax; j++ {
						ch := rune(lines[lineIndex+i][charIndex+j])
						if ch != '.' && !unicode.IsDigit(ch) {
							currentNumberValid = true
						}
					}
				}
			} else {
				if currentNumber == "" {
					continue
				}

				if currentNumberValid {
					num, _ := strconv.Atoi(currentNumber)
					sum += num
				}

				currentNumber = ""
				currentNumberValid = false
			}
		}
	}

	return sum
}

func Day3Part2(input string) any {
	lines := strings.Split(input, "\n")

	gearsMap := make(map[string][]int)
	currentNumber := ""
	currentAsteriskLocation := ""
	for lineIndex, line := range lines {
		lineOffsetMin := -1
		lineOffsetMax := 1
		if lineIndex == 0 {
			lineOffsetMin = 0
		}

		// 2 because strings.Split on \n return empty string as last element
		if lineIndex == len(lines)-2 {
			lineOffsetMax = 0
		}

		for charIndex, char := range line {
			if unicode.IsDigit(char) {
				currentNumber += string(char)

				if currentAsteriskLocation != "" {
					continue
				}

				charOffsetMin := -1
				charOffsetMax := 1

				if charIndex == 0 {
					charOffsetMin = 0
				}

				if charIndex == len(line)-1 {
					charOffsetMax = 0
				}

				for i := lineOffsetMin; i <= lineOffsetMax; i++ {
					for j := charOffsetMin; j <= charOffsetMax; j++ {
						ch := rune(lines[lineIndex+i][charIndex+j])
						if ch == '*' {
							currentAsteriskLocation = fmt.Sprintf("%dx%d", lineIndex+i, charIndex+j)
						}
					}
				}
			} else {
				if currentNumber == "" {
					continue
				}

				if currentAsteriskLocation != "" {
					num, _ := strconv.Atoi(currentNumber)
					gearsMap[currentAsteriskLocation] = append(gearsMap[currentAsteriskLocation], num)
				}

				currentNumber = ""
				currentAsteriskLocation = ""
			}
		}
	}

	sum := 0
	for _, nums := range gearsMap {
		if len(nums) == 2 {
			sum += nums[0] * nums[1]
		}
	}

	return sum
}
