package main

import (
	"fmt"
)

type Input struct {
	maze     [][]byte
	entrance []int
}

type TestSuit struct {
	name   string
	input  Input
	output int
}

func main() {
	caseTests := []TestSuit{
		{
			name: "first",
			input: Input{
				maze:     [][]byte{{'+', '+', '.', '+'}, {'.', '.', '.', '+'}, {'+', '+', '+', '.'}},
				entrance: []int{1, 2},
			},
			output: 1,
		},
		{
			name: "second",
			input: Input{
				maze:     [][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}},
				entrance: []int{1, 0},
			},
			output: 2,
		},
		{
			name: "third",
			input: Input{
				maze:     [][]byte{{'.', '+'}},
				entrance: []int{0, 0},
			},
			output: -1,
		},
	}

	for _, caseTest := range caseTests {
		result := nearestExit(caseTest.input.maze, caseTest.input.entrance)
		if result != caseTest.output {
			fmt.Printf("\033[31mF\033[0m")
			continue
		}
		fmt.Printf("\033[32m.\033[0m")
	}
}

func nearestExit(maze [][]byte, entrance []int) int {
	// x, y
	movements := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	queue := []int(entrance)

	rows, columns := len(maze)-1, len(maze[0])-1

	return -1
}
