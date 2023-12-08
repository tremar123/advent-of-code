package solutions

import (
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cards string
	bid   int
	rank  int
}

func Day7Part1(input string) any {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	highCard := []Hand{}
	onePair := []Hand{}
	twoPair := []Hand{}
	threeOfAKind := []Hand{}
	fullHouse := []Hand{}
	fourOfAKind := []Hand{}
	fiveOfAKind := []Hand{}

	for _, line := range lines {
		hand := Hand{}
		split := strings.Fields(line)
		hand.cards = split[0]
		hand.bid, _ = strconv.Atoi(split[1])

		mapOfChars := map[rune]int{}

		for _, ch := range hand.cards {
			mapOfChars[ch]++
		}

	switchLabel:
		switch len(mapOfChars) {
		case 5: // if there are 5 chars we know each is defferent
			highCard = append(highCard, hand)
		case 4: // if there are 4 chars we know one is repeated
			onePair = append(onePair, hand)
		case 3:
			// if any of chars doesn't repeat 3 times its two pair
			for _, count := range mapOfChars {
				if count == 3 {
					threeOfAKind = append(threeOfAKind, hand)
					break switchLabel
				}
			}

			twoPair = append(twoPair, hand)
		case 2:
			// if any of chars doesn't repeat 3 times its two pair
			for _, count := range mapOfChars {
				if count == 4 {
					fourOfAKind = append(fourOfAKind, hand)
					break switchLabel
				}
			}

			fullHouse = append(fullHouse, hand)
		case 1: // all the same
			fiveOfAKind = append(fiveOfAKind, hand)
		}
	}

	slices.SortFunc(highCard, cmpPart1)
	slices.SortFunc(onePair, cmpPart1)
	slices.SortFunc(twoPair, cmpPart1)
	slices.SortFunc(threeOfAKind, cmpPart1)
	slices.SortFunc(fullHouse, cmpPart1)
	slices.SortFunc(fourOfAKind, cmpPart1)
	slices.SortFunc(fiveOfAKind, cmpPart1)

	allHands := make([]Hand, 0, len(highCard)+len(onePair)+len(twoPair)+len(threeOfAKind)+len(fullHouse)+len(fourOfAKind)+len(fiveOfAKind))

	allHands = append(allHands, highCard...)
	allHands = append(allHands, onePair...)
	allHands = append(allHands, twoPair...)
	allHands = append(allHands, threeOfAKind...)
	allHands = append(allHands, fullHouse...)
	allHands = append(allHands, fourOfAKind...)
	allHands = append(allHands, fiveOfAKind...)

	result := 0
	for i, hand := range allHands {
		result += hand.bid * (i + 1)
	}

	return result
}

func cmpPart1(a Hand, b Hand) int {
	power := []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
	for i := 0; i < 5; i++ {
		aCard := slices.Index(power, rune(a.cards[i]))
		bCard := slices.Index(power, rune(b.cards[i]))
		if aCard < bCard {
			return -1
		} else if aCard > bCard {
			return 1
		}
	}

	return 0
}

func Day7Part2(input string) any {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	highCard := []Hand{}
	onePair := []Hand{}
	twoPair := []Hand{}
	threeOfAKind := []Hand{}
	fullHouse := []Hand{}
	fourOfAKind := []Hand{}
	fiveOfAKind := []Hand{}

	for _, line := range lines {
		hand := Hand{}
		split := strings.Fields(line)
		hand.cards = split[0]
		hand.bid, _ = strconv.Atoi(split[1])

		mapOfChars := map[rune]int{}

		for _, ch := range hand.cards {
			mapOfChars[ch]++
		}

		Jcount, containsJ := mapOfChars['J']

	switchLabel:
		switch len(mapOfChars) {
		case 5: // if there are 5 chars we know each is defferent
			if containsJ {
				onePair = append(onePair, hand)
				break switchLabel
			}
			highCard = append(highCard, hand)
		case 4: // if there are 4 chars we know one is repeated
			if containsJ {
				threeOfAKind = append(threeOfAKind, hand)
				break switchLabel
			}
			onePair = append(onePair, hand)
		case 3:
			// if any of chars doesn't repeat 3 times its two pair
			for _, count := range mapOfChars {
				if count == 3 {
					if containsJ {
						fourOfAKind = append(fourOfAKind, hand)
						break switchLabel
					}
					threeOfAKind = append(threeOfAKind, hand)
					break switchLabel
				}
			}

			if containsJ {
				if Jcount == 1 {
					fullHouse = append(fullHouse, hand)
					break switchLabel
				} else if Jcount == 2 {
					fourOfAKind = append(fourOfAKind, hand)
					break switchLabel
				}
			}
			twoPair = append(twoPair, hand)
		case 2:
			// if any of chars doesn't repeat 3 times its two pair
			for _, count := range mapOfChars {
				if count == 4 {
					if containsJ {
						fiveOfAKind = append(fiveOfAKind, hand)
						break switchLabel
					}
					fourOfAKind = append(fourOfAKind, hand)
					break switchLabel
				}
			}

			if containsJ {
				fiveOfAKind = append(fiveOfAKind, hand)
				break switchLabel
			}

			fullHouse = append(fullHouse, hand)
		case 1: // all the same
			fiveOfAKind = append(fiveOfAKind, hand)
		}
	}

	slices.SortFunc(highCard, cmpPart2)
	slices.SortFunc(onePair, cmpPart2)
	slices.SortFunc(twoPair, cmpPart2)
	slices.SortFunc(threeOfAKind, cmpPart2)
	slices.SortFunc(fullHouse, cmpPart2)
	slices.SortFunc(fourOfAKind, cmpPart2)
	slices.SortFunc(fiveOfAKind, cmpPart2)

	allHands := make([]Hand, 0, len(highCard)+len(onePair)+len(twoPair)+len(threeOfAKind)+len(fullHouse)+len(fourOfAKind)+len(fiveOfAKind))

	allHands = append(allHands, highCard...)
	allHands = append(allHands, onePair...)
	allHands = append(allHands, twoPair...)
	allHands = append(allHands, threeOfAKind...)
	allHands = append(allHands, fullHouse...)
	allHands = append(allHands, fourOfAKind...)
	allHands = append(allHands, fiveOfAKind...)

	result := 0
	for i, hand := range allHands {
		result += hand.bid * (i + 1)
	}

	return result
}

func cmpPart2(a Hand, b Hand) int {
	power := []rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}
	for i := 0; i < 5; i++ {
		aCard := slices.Index(power, rune(a.cards[i]))
		bCard := slices.Index(power, rune(b.cards[i]))
		if aCard < bCard {
			return -1
		} else if aCard > bCard {
			return 1
		}
	}

	return 0
}
