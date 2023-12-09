package solutions

import (
	"strconv"
	"strings"
)

func Day9Part1(input string) any {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	readings := make([][]int, 0, len(lines))

	for _, line := range lines {
		numsStr := strings.Fields(line)
		// reserve one more cap for our new number
		nums := make([]int, 0, len(numsStr)+1)
		for _, numStr := range numsStr {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}
		readings = append(readings, nums)
	}

	sum := 0
	for _, reading := range readings {
		diffs := [][]int{reading}
		for i := 0; i < len(diffs); i++ {
			newDiffs := make([]int, 0, len(diffs))
			for j := 0; j < len(diffs[i])-1; j++ {
				newDiffs = append(newDiffs, diffs[i][j+1]-diffs[i][j])
			}

			allZero := true
			for _, num := range newDiffs {
				if num != 0 {
					allZero = false
				}
			}

			if allZero {
				newDiffs = append(newDiffs, 0)
			}

			diffs = append(diffs, newDiffs)

			if allZero {
				break
			}
		}

		for i := len(diffs) - 2; i >= 0; i-- {
			newNum := diffs[i][len(diffs[i])-1] + diffs[i+1][len(diffs[i+1])-1]
			diffs[i] = append(diffs[i], newNum)
		}

		sum += diffs[0][len(diffs[0])-1]
	}

	return sum
}

func Day9Part2(input string) any {
	lines := strings.Split(input, "\n")
	lines = lines[:len(lines)-1]
	readings := make([][]int, 0, len(lines))

	for _, line := range lines {
		numsStr := strings.Fields(line)
		// reserve one more cap for our new number
		nums := make([]int, 0, len(numsStr)+1)
		for _, numStr := range numsStr {
			num, _ := strconv.Atoi(numStr)
			nums = append(nums, num)
		}
		readings = append(readings, nums)
	}

	sum := 0
	for _, reading := range readings {
		diffs := [][]int{reading}
		for i := 0; i < len(diffs); i++ {
			newDiffs := make([]int, 0, len(diffs))
			for j := 0; j < len(diffs[i])-1; j++ {
				newDiffs = append(newDiffs, diffs[i][j+1]-diffs[i][j])
			}

			allZero := true
			for _, num := range newDiffs {
				if num != 0 {
					allZero = false
				}
			}

			if allZero {
				newDiffs = append(newDiffs, 0)
			}

			diffs = append(diffs, newDiffs)

			if allZero {
				break
			}
		}

		for i := len(diffs) - 2; i >= 0; i-- {
			newNum := diffs[i][0] - diffs[i+1][0]
			diffs[i] = append([]int{newNum}, diffs[i]...)
		}

		sum += diffs[0][0]
	}

	return sum
}
