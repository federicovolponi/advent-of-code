package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const paperRoll = "@"

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}, {1, 1}, {1, -1}, {-1, 1}, {-1, -1}}

func readGrid() [][]string {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)

	var grid [][]string
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, strings.Split(line, ""))
	}
	return grid
}

func adjacentRollPapers(x, y int, grid [][]string) int {
	nAdj := 0
	for _, dir := range directions {
		dx, dy := dir[0], dir[1]
		newX := x + dx
		newY := y + dy
		if newX >= 0 && newX < len(grid[0]) && newY >= 0 && newY < len(grid) && grid[newX][newY] == paperRoll {
			nAdj++
		}
	}
	return nAdj
}

func countAccessibleForklift(grid [][]string) int {
	res := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == paperRoll && adjacentRollPapers(i, j, grid) < 4 {
				res++
			}
		}
	}
	return res
}

func countAccessibleForkliftPt2(grid [][]string) int {
	res := 0
	check := -1
	for check != res {
		check = res
		for i := range grid {
			for j := range grid[i] {
				if grid[i][j] == paperRoll && adjacentRollPapers(i, j, grid) < 4 {
					grid[i][j] = "X"
					res++
				}
			}
		}
	}
	return res
}

func main() {
	grid := readGrid()
	fmt.Println("First part solution: ", countAccessibleForklift(grid))
	fmt.Println("Second part solution: ", countAccessibleForkliftPt2(grid))
}
