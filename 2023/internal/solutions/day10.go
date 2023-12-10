package solutions

import (
	"math"
	"strings"
)

type Point struct {
	x int
	y int
}

// not the best solution i guess but whatever it works
func Day10Part1(input string) any {
	sketch := strings.Split(input, "\n")
	sketch = sketch[:len(sketch)-1]
	var currentPoint Point

	for y := 0; y < len(sketch); y++ {
		for x := 0; x < len(sketch[y]); x++ {
			if sketch[y][x] == 'S' {
				currentPoint = Point{x, y}
			}
		}
	}

	var nextPoint Point

	if currentPoint.y != 0 {
		switch rune(sketch[currentPoint.y-1][currentPoint.x]) {
		case '|':
			currentPoint = Point{currentPoint.x, currentPoint.y - 1}
			nextPoint = Point{currentPoint.x, currentPoint.y - 1}
			goto neverUseGoto
		case '7':
			currentPoint = Point{currentPoint.x, currentPoint.y - 1}
			nextPoint = Point{currentPoint.x - 1, currentPoint.y}
			goto neverUseGoto
		case 'F':
			currentPoint = Point{currentPoint.x, currentPoint.y - 1}
			nextPoint = Point{currentPoint.x + 1, currentPoint.y}
			goto neverUseGoto
		}
	}

	if currentPoint.y != len(sketch)-1 {
		switch rune(sketch[currentPoint.y+1][currentPoint.x]) {
		case '|':
			currentPoint = Point{currentPoint.x, currentPoint.y + 1}
			nextPoint = Point{currentPoint.x, currentPoint.y + 1}
			goto neverUseGoto
		case 'L':
			currentPoint = Point{currentPoint.x, currentPoint.y + 1}
			nextPoint = Point{currentPoint.x + 1, currentPoint.y}
			goto neverUseGoto
		case 'J':
			currentPoint = Point{currentPoint.x, currentPoint.y + 1}
			nextPoint = Point{currentPoint.x - 1, currentPoint.y}
			goto neverUseGoto
		}
	}

	if currentPoint.x != 0 {
		switch rune(sketch[currentPoint.y][currentPoint.x-1]) {
		case '-':
			currentPoint = Point{currentPoint.x - 1, currentPoint.y}
			nextPoint = Point{currentPoint.x - 1, currentPoint.y}
			goto neverUseGoto
		case 'L':
			currentPoint = Point{currentPoint.x - 1, currentPoint.y}
			nextPoint = Point{currentPoint.x, currentPoint.y - 1}
			goto neverUseGoto
		case 'F':
			currentPoint = Point{currentPoint.x - 1, currentPoint.y}
			nextPoint = Point{currentPoint.x, currentPoint.y + 1}
			goto neverUseGoto
		}
	}

	if currentPoint.x != len(sketch[0])-1 {
		switch rune(sketch[currentPoint.y][currentPoint.x+1]) {
		case '-':
			currentPoint = Point{currentPoint.x + 1, currentPoint.y}
			nextPoint = Point{currentPoint.x + 1, currentPoint.y}
			goto neverUseGoto
		case '7':
			currentPoint = Point{currentPoint.x + 1, currentPoint.y}
			nextPoint = Point{currentPoint.x, currentPoint.y + 1}
			goto neverUseGoto
		case 'J':
			currentPoint = Point{currentPoint.x + 1, currentPoint.y}
			nextPoint = Point{currentPoint.x, currentPoint.y - 1}
			goto neverUseGoto
		}
	}

neverUseGoto:
	steps := 1
	for {
		if rune(sketch[nextPoint.y][nextPoint.x]) == 'S' {
			break
		}

		x := nextPoint.x - currentPoint.x
		y := nextPoint.y - currentPoint.y

		currentPoint = nextPoint

		if x == -1 {
			// comming from right
			switch rune(sketch[nextPoint.y][nextPoint.x]) {
			case 'L':
				nextPoint = Point{currentPoint.x, currentPoint.y - 1}
			case '-':
				nextPoint = Point{currentPoint.x - 1, currentPoint.y}
			case 'F':
				nextPoint = Point{currentPoint.x, currentPoint.y + 1}
			}
		} else if x == 1 {
			// comming from left
			switch rune(sketch[nextPoint.y][nextPoint.x]) {
			case 'J':
				nextPoint = Point{currentPoint.x, currentPoint.y - 1}
			case '-':
				nextPoint = Point{currentPoint.x + 1, currentPoint.y}
			case '7':
				nextPoint = Point{currentPoint.x, currentPoint.y + 1}
			}
		} else if y == -1 {
			// comming from bottom
			switch rune(sketch[nextPoint.y][nextPoint.x]) {
			case 'F':
				nextPoint = Point{currentPoint.x + 1, currentPoint.y}
			case '|':
				nextPoint = Point{currentPoint.x, currentPoint.y - 1}
			case '7':
				nextPoint = Point{currentPoint.x - 1, currentPoint.y}
			}
		} else if y == 1 {
			// comming from top
			switch rune(sketch[nextPoint.y][nextPoint.x]) {
			case 'L':
				nextPoint = Point{currentPoint.x + 1, currentPoint.y}
			case '|':
				nextPoint = Point{currentPoint.x, currentPoint.y + 1}
			case 'J':
				nextPoint = Point{currentPoint.x - 1, currentPoint.y}
			}
		}

		steps++
	}

	return math.Round(float64(steps) / 2.0)
}
