package main

import (
	"fmt"
)

func getPanel(grid [][]int, r, c int) []int {
	rr := r / 3
	cc := c / 3
	out := make([]int, 0, 9)
	for _, x := range grid[rr*3 : rr*3+3] {
		out = append(out, x[cc*3:cc*3+3]...)
	}
	return out
}

func getCol(grid [][]int, c int) []int {
	out := make([]int, 0, 9)
	for _, x := range grid {
		out = append(out, x[c])
	}
	return out
}

// solve the sudoku with backtracking algorithm
func solve(grid [][]int) [][]int {
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

	// Ref: https://dingo.sbs.arizona.edu/~sandiway/sudoku/examples.html
	inputGrid := [][]int{
		{0, 0, 0, 2, 6, 0, 7, 0, 1},
		{6, 8, 0, 0, 7, 0, 0, 9, 0},
		{1, 9, 0, 0, 0, 4, 5, 0, 0},
		{8, 2, 0, 1, 0, 0, 0, 4, 0},
		{0, 0, 4, 6, 0, 2, 9, 0, 0},
		{0, 5, 0, 0, 0, 3, 0, 2, 8},
		{0, 0, 9, 3, 0, 0, 0, 7, 4},
		{0, 4, 0, 0, 5, 0, 0, 3, 6},
		{7, 0, 3, 0, 1, 8, 0, 0, 0},
	}
	printGrid(inputGrid)
	out := solve(inputGrid)
	printGrid(out)
}
