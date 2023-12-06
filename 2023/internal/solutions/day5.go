package solutions

import (
	"strconv"
	"strings"
	"sync"
)

func Day5Part1(input string) any {
	blocks := strings.Split(input, "\n\n")
	operations := parseBlocks(blocks)
	seedsList := strings.Split(blocks[0], ": ")[1]
	seedsStr := strings.Fields(seedsList)
	seeds := make([]int, len(seedsStr))

	for i := range seedsStr {
		seeds[i], _ = strconv.Atoi(seedsStr[i])
	}

	minValue := 0
	minValueSet := false

	for _, seed := range seeds {
		calculateLocation(&seed, operations)

		if seed < minValue || !minValueSet {
			minValue = seed
			minValueSet = true
		}
	}

	return minValue
}

func Day5Part2(input string) any {
	blocks := strings.Split(input, "\n\n")
	operations := parseBlocks(blocks)
	seedsList := strings.Split(blocks[0], ": ")[1]
	seedsStr := strings.Fields(seedsList)
	seeds := make([][2]int, 0, len(seedsStr))

	for i := 0; i < len(seedsStr); i += 2 {
		start, _ := strconv.Atoi(seedsStr[i])
		rang, _ := strconv.Atoi(seedsStr[i+1])

		seeds = append(seeds, [2]int{start, rang})
	}

	minValue := 0
	minValueSet := false

	for _, pair := range seeds {
		for i := pair[0]; i <= pair[0]+pair[1]; i++ {
			seed := i
			calculateLocation(&seed, operations)

			if seed < minValue || !minValueSet {
				minValue = seed
				minValueSet = true
			}
		}
	}

	return minValue
}

func Day5Part2Parallel(input string) any {
	blocks := strings.Split(input, "\n\n")
	operations := parseBlocks(blocks)
	seedsList := strings.Split(blocks[0], ": ")[1]
	seedsStr := strings.Fields(seedsList)
	seeds := make([][2]int, 0, len(seedsStr))

	for i := 0; i < len(seedsStr); i += 2 {
		start, _ := strconv.Atoi(seedsStr[i])
		rang, _ := strconv.Atoi(seedsStr[i+1])

		seeds = append(seeds, [2]int{start, rang})
	}

	overallMinValue := 0
	overallMinValueSet := false
	overallMinValueMutex := sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, pair := range seeds {
		go func(pair [2]int) {
			wg.Add(1)
			minValue := 0
			minValueSet := false

			for i := pair[0]; i <= pair[0]+pair[1]; i++ {
				seed := i
				calculateLocation(&seed, operations)

				if seed < minValue || !minValueSet {
					minValue = seed
					minValueSet = true
				}
			}

			overallMinValueMutex.Lock()
			if minValue < overallMinValue || !overallMinValueSet {
				overallMinValue = minValue
				overallMinValueSet = true
			}
			overallMinValueMutex.Unlock()

			wg.Done()
		}(pair)
	}

	wg.Wait()

	return overallMinValue
}

func calculateLocation(seed *int, operations [][][3]int) {
	for _, line := range operations {
		for _, numbers := range line {
			if *seed >= numbers[1] && *seed <= numbers[1]+numbers[2] {
				*seed = *seed - numbers[1] + numbers[0]
				break
			}
		}
	}
}

func parseBlocks(blocks []string) [][][3]int {
	// 1st level represents convert map
	// 2st level represents lines of map
	// 3rd numbers
	operations := make([][][3]int, len(blocks)-1)

	for blockIndex := 1; blockIndex < len(blocks); blockIndex++ {
		lines := strings.Split(blocks[blockIndex], "\n")

		for lineIndex := 1; lineIndex < len(lines); lineIndex++ {
			if lines[lineIndex] == "" {
				continue
			}

			numbersStr := strings.Fields(lines[lineIndex])
			numbers := [3]int{}

			for i := range numbersStr {
				numbers[i], _ = strconv.Atoi(numbersStr[i])
			}

			operations[blockIndex-1] = append(operations[blockIndex-1], numbers)
		}
	}

	return operations
}
