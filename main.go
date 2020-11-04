package main

import (
	"fmt"
)

func getPanel(grid [][]int, r, c int) []int {
	if r > 2 || c > 2 {
		return nil
	}

	out := make([]int, 0, 9)
	for _, x := range grid[r*3 : r*3+3] {
		out = append(out, x[c*3:c*3+3]...)
	}

	return out
}

func getCol(grid [][]int, c int) []int {
	if c > 9 {
		return nil
	}

	out := make([]int, 0, 9)
	for _, x := range grid {
		out = append(out, x[c])
	}

	return out
}

// solve the sudoku with backtracking algorithm
func solve(grid [][]int) [][]int {
	// out := make([][]int)
	for i, x := range grid {
		for j, y := range x {
			if y == 0 {
				for _, n := range []int{1, 2, 3, 4, 5, 6, 7, 8, 9} {
					if validInput(grid, i, j, n) {
						grid[i][j] = n
						out := solve(grid)
						if out != nil {
							return grid // exit / termination condition
						}
						grid[i][j] = 0
					}
				}
				return nil // no possible soln found, tell the caller to try something different
			}
		}
	}
	return grid // all square filled
}

func contains(s []int, n int) bool {
	for _, e := range s {
		if n == e {
			return true
		}
	}
	return false
}

func validInput(grid [][]int, r, c, n int) bool {
	if r > 9 || c > 9 {
		return false
	}

	if n > 9 || n < 1 {
		return false
	}

	return !(contains(grid[r], n) || contains(getCol(grid, c), n) || contains(getPanel(grid, r, c), n))
}

func printGrid(grid [][]int) {
	fmt.Println("\nGrid:\n")
	for i, x := range grid {
		fmt.Printf("%d %d %d | %d %d %d | %d %d %d\n", x[0], x[1], x[2], x[3], x[4], x[5], x[6], x[7], x[8])
		if (i+1)%3 == 0 && i != 8 {
			fmt.Println("------+-------+------")
		}
	}
	fmt.Println("")
}

func main() {
	inputGrid := [][]int{
		{0, 2, 3, 4, 5, 6, 7, 8, 9},
		{2, 0, 4, 5, 6, 7, 8, 9, 1},
		{3, 4, 5, 6, 7, 8, 9, 1, 2},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{5, 6, 7, 8, 9, 1, 2, 3, 4},
		{6, 7, 8, 9, 1, 2, 3, 4, 5},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{8, 9, 1, 2, 3, 4, 5, 6, 7},
		{9, 1, 2, 3, 4, 5, 6, 7, 8},
	}
	out := solve(inputGrid)
	printGrid(out)
}
