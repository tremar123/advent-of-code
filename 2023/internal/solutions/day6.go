package solutions

import (
	"strconv"
	"strings"
)

func Day6Part1(input string) any {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	for i := range lines {
		lines[i] = strings.Split(lines[i], ":")[1]
		lines[i] = strings.TrimSpace(lines[i])
	}

	timesStr := strings.Fields(lines[0])
	distancesStr := strings.Fields(lines[1])

	times := make([]int, 0, len(timesStr))
	distances := make([]int, 0, len(distancesStr))

	for _, s := range timesStr {
		time, _ := strconv.Atoi(s)
		times = append(times, time)
	}

	for _, s := range distancesStr {
		distance, _ := strconv.Atoi(s)
		distances = append(distances, distance)
	}

	result := 1

	for i := 0; i < len(times); i++ {
		wins := 0
		// speed is 1 milimeter per milisecond so hold == speed
		for hold := 0; hold <= times[i]; hold++ {
			travelled := (times[i] - hold) * hold

			if travelled > distances[i] {
				wins++
			}
		}

		result *= wins
	}

	return result
}

func Day6Part2(input string) any {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]

	for i := range lines {
		lines[i] = strings.Split(lines[i], ":")[1]
		lines[i] = strings.ReplaceAll(lines[i], " ", "")
	}

	time, _ := strconv.Atoi(lines[0])
	distance, _ := strconv.Atoi(lines[1])

	wins := 0
	for hold := 0; hold <= time; hold++ {
		travelled := (time - hold) * hold

		if travelled > distance {
			wins++
		}
	}

	return wins
}
