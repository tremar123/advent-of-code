package main

import (
	"aoc23/internal/solutions"
	"flag"
	"fmt"
	"os"
)

func main() {
	day := flag.Int("day", 0, "Day of the advent")
	part := flag.Int("part", 1, "Puzzle number")
	inputPath := flag.String("input", "", "Path to input file")

	flag.Parse()

	if *day == 0 {
		fmt.Fprint(os.Stderr, "day flag is required")
		os.Exit(1)
	}

	if *part < 1 || *part > 2 {
		fmt.Fprint(os.Stderr, "There are only 2 puzzles")
	}

	file, err := os.ReadFile(*inputPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%q", err)
		os.Exit(2)
	}

	input := string(file)

	var output any

	switch *day {
	case 1:
		if *part == 1 {
			output = solutions.Day1Part1(input)
		} else {
			output = solutions.Day1Part2(input)
		}
	case 2:
		if *part == 1 {
			output = solutions.Day2Part1(input)
		} else {
			output = solutions.Day2Part2(input)
		}
	case 3:
		if *part == 1 {
			output = solutions.Day3Part1(input)
		} else {
			output = solutions.Day3Part2(input)
		}
	case 4:
		if *part == 1 {
			output = solutions.Day4Part1(input)
		} else {
			output = solutions.Day4Part2(input)
		}
	case 5:
		if *part == 1 {
			output = solutions.Day5Part1(input)
		} else {
			output = solutions.Day5Part2Parallel(input)
		}
	case 6:
		if *part == 1 {
			output = solutions.Day6Part1(input)
		} else {
			output = solutions.Day6Part2(input)
		}
	case 7:
		if *part == 1 {
			output = solutions.Day7Part1(input)
		} else {
			output = solutions.Day7Part2(input)
		}
	case 8:
		if *part == 1 {
			output = solutions.Day8Part1(input)
		} else {
			output = solutions.Day8Part2(input)
		}
	case 9:
		if *part == 1 {
			output = solutions.Day9Part1(input)
		} else {
			output = solutions.Day9Part2(input)
		}
	case 10:
		if *part == 1 {
			output = solutions.Day10Part1(input)
		} else {
			panic("unimplemented")
		}
	case 11:
		if *part == 1 {
			panic("unimplemented")
		} else {
			panic("unimplemented")
		}
	case 12:
		if *part == 1 {
			panic("unimplemented")
		} else {
			panic("unimplemented")
		}
	case 13:
		if *part == 1 {
			panic("unimplemented")
		} else {
			panic("unimplemented")
		}
	case 14:
		if *part == 1 {
			panic("unimplemented")
		} else {
			panic("unimplemented")
		}
	case 15:
		if *part == 1 {
			panic("unimplemented")
		} else {
			panic("unimplemented")
		}
	case 16:
		if *part == 1 {
			panic("unimplemented")
		} else {
			panic("unimplemented")
		}
	case 17:
		if *part == 1 {
			panic("unimplemented")
		} else {
			panic("unimplemented")
		}
	case 18:
		if *part == 1 {
			panic("unimplemented")
		} else {
			panic("unimplemented")
		}
	case 19:
		if *part == 1 {
			panic("unimplemented")
		} else {
			panic("unimplemented")
		}
	case 20:
		if *part == 1 {
			panic("unimplemented")
		} else {
			panic("unimplemented")
		}
	case 21:
		if *part == 1 {
			panic("unimplemented")
		} else {
			panic("unimplemented")
		}
	case 22:
		if *part == 1 {
			panic("unimplemented")
		} else {
			panic("unimplemented")
		}
	case 23:
		if *part == 1 {
			panic("unimplemented")
		} else {
			panic("unimplemented")
		}
	case 24:
		if *part == 1 {
			panic("unimplemented")
		} else {
			panic("unimplemented")
		}
	case 25:
		if *part == 1 {
			panic("unimplemented")
		} else {
			panic("unimplemented")
		}
	}

	fmt.Println(output)
}
