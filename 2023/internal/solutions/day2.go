package solutions

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	RED   = "red"
	GREEN = "green"
	BLUE  = "blue"
)

type Cube struct {
	color string
	count int
}

type Game struct {
	id    int
	turns [][]Cube
}

func parseGames(input string) []Game {
	lines := strings.Split(input, "\n")
	games := make([]Game, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			break
		}
		game := Game{}
		words := strings.Split(line, " ")
		game.id, _ = strconv.Atoi(strings.TrimSuffix(words[1], ":"))

		currentTurn := []Cube{}
		for i := 2; i < len(words); i += 2 {
			count, err := strconv.Atoi(words[i])
			if err != nil {
				fmt.Fprint(os.Stderr, err.Error())
				panic("there should always be a number")
			}

			color := words[i+1]
			if color[len(color)-1] == ';' || color[len(color)-1] != ',' {
				currentTurn = append(currentTurn, Cube{
					count: count,
					color: strings.TrimSuffix(color, ";"),
				})
				game.turns = append(game.turns, currentTurn)
				currentTurn = []Cube{}
			} else {
				currentTurn = append(currentTurn, Cube{
					count: count,
					color: strings.TrimSuffix(color, ","),
				})
			}
		}
		games = append(games, game)
	}

	return games
}

func Day2Part1(input string) any {
	cubes := map[string]int{RED: 12, GREEN: 13, BLUE: 14}
	games := parseGames(input)
	sumOfPossibleGames := 0
	for _, game := range games {
		shouldGameCount := true
	turnLoop:
		for _, turn := range game.turns {
			for _, cube := range turn {
				if cubes[cube.color] < cube.count {
					shouldGameCount = false
					break turnLoop
				}
			}
		}
		if shouldGameCount {
			sumOfPossibleGames += game.id
		}
	}

	return sumOfPossibleGames
}

func Day2Part2(input string) any {
	games := parseGames(input)

	sum := 0
	for _, game := range games {
		cubesMax := map[string]int{}
		for _, turn := range game.turns {
			for _, cube := range turn {
				count := cubesMax[cube.color]
				if cube.count > count {
					cubesMax[cube.color] = cube.count
				}
			}
		}
		sum += cubesMax[RED] * cubesMax[GREEN] * cubesMax[BLUE]
	}

	return sum
}
