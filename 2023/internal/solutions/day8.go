package solutions

import (
	"math"
	"strings"
)

type Node struct {
	left  string
	right string
}

func Day8Part1(input string) any {
	blocks := strings.Split(input, "\n\n")
	instructions := blocks[0]
	lines := strings.Split(blocks[1], "\n")
	lines = lines[:len(lines)-1]
	nodes := make(map[string]Node, len(lines))

	for _, line := range lines {
		split := strings.Split(line, " = ")
		lr := strings.Trim(split[1], "()")
		lrSplit := strings.Split(lr, ", ")
		node := Node{
			left:  lrSplit[0],
			right: lrSplit[1],
		}

		nodes[split[0]] = node
	}

	steps := 0
	found := false
	current := "AAA"
	for !found {
		for _, instruction := range instructions {
			if instruction == 'L' {
				current = nodes[current].left
			} else if instruction == 'R' {
				current = nodes[current].right
			}

			steps++

			if current == "ZZZ" {
				found = true
				break
			}
		}
	}

	return steps
}

type CurrentNode struct {
	key   string
	found bool
	steps int
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return int(math.Abs(float64(a*b)) / float64(gcd(a, b)))
}

func lcmMultiple(numbers []int) int {
	result := numbers[0]
	for i := 1; i < len(numbers); i++ {
		result = lcm(result, numbers[i])
	}
	return result
}

func Day8Part2(input string) any {
	blocks := strings.Split(input, "\n\n")
	instructions := blocks[0]
	lines := strings.Split(blocks[1], "\n")
	lines = lines[:len(lines)-1]
	nodes := make(map[string]Node, len(lines))

	for _, line := range lines {
		split := strings.Split(line, " = ")
		lr := strings.Trim(split[1], "()")
		lrSplit := strings.Split(lr, ", ")
		node := Node{
			left:  lrSplit[0],
			right: lrSplit[1],
		}

		nodes[split[0]] = node
	}

	currentNodes := []CurrentNode{}

	for key := range nodes {
		if key[2] == 'A' {
			currentNodes = append(currentNodes, CurrentNode{
				key:   key,
				found: false,
			})
		}
	}

	for i := range currentNodes {
		for !currentNodes[i].found {
			for _, instruction := range instructions {
				if instruction == 'L' {
					currentNodes[i].key = nodes[currentNodes[i].key].left
				} else if instruction == 'R' {
					currentNodes[i].key = nodes[currentNodes[i].key].right
				}

				currentNodes[i].steps++

				if currentNodes[i].key[2] == 'Z' {
					currentNodes[i].found = true
					break
				}
			}
		}
	}

    nums := make([]int, len(currentNodes))

    for i, node := range currentNodes {
        nums[i] = node.steps
    }

    return lcmMultiple(nums)
}
