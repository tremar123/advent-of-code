package solutions

import (
	"slices"
	"strconv"
	"strings"
)

type Card struct {
	id             int
	winningNumbers []int
	myNumbers      []int
}

// awful
func parseInput(input string) []Card {
	lines := strings.Split(input, "\n")
	cards := make([]Card, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			continue
		}

		card := Card{}
		split := strings.Split(line, ": ")
		card.id, _ = strconv.Atoi(strings.Fields(split[0])[1])
		numbers := strings.Split(split[1], " | ")
		for i := range numbers {
			numbers[i] = strings.Trim(numbers[i], " ")
		}
		winningNumbersStr := strings.Fields(numbers[0])
		myNumbersStr := strings.Fields(numbers[1])

		winningNumbers := make([]int, 0, len(winningNumbersStr))
		myNumbers := make([]int, 0, len(myNumbersStr))

		for _, num := range winningNumbersStr {
			n, _ := strconv.Atoi(num)
			winningNumbers = append(winningNumbers, n)
		}

		for _, num := range myNumbersStr {
			n, _ := strconv.Atoi(num)
			myNumbers = append(myNumbers, n)
		}

		card.winningNumbers = winningNumbers
		card.myNumbers = myNumbers

		cards = append(cards, card)
	}

	return cards
}

func Day4Part1(input string) any {
	cards := parseInput(input)

	sum := 0
	for _, card := range cards {
		points := 0

		for _, myNum := range card.myNumbers {
			if slices.Contains(card.winningNumbers, myNum) {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}

		sum += points
	}

	return sum
}

func Day4Part2(input string) any {
	cards := parseInput(input)

	counts := make([]int, len(cards))

	for i := range counts {
		counts[i] = 1
	}

	for i, card := range cards {
		winningNums := 0
		for _, myNum := range card.myNumbers {
			if slices.Contains(card.winningNumbers, myNum) {
				winningNums++
			}
		}

		for j := 1; j <= winningNums; j++ {
			counts[i+j] += 1 * counts[i]
		}
	}

	sum := 0
	for _, count := range counts {
		sum += count
	}

	return sum
}
