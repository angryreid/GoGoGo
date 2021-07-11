package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)
	maze := make([][]int, row) // Slice 6 * 5
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}
	return maze

}

type point struct {
	i, j int
}

var dirs = [4]point{
	{-1, 0},
	{0, -1},
	{1, 0},
	{0, 1},
}

func (p point) add(r point) point {
	return point{p.j + r.j, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}
	return grid[p.i][p.j], true
}

func walk(maze [][]int, start point, end point) {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	nextPointQueue := []point{start}

	for len(nextPointQueue) > 0 {
		cur := nextPointQueue[0]
		nextPointQueue = nextPointQueue[1:]

		for _, dir := range dirs {
			next := cur.add(dir)
			// maze at next is 0
			// and steps at next is 0
			// and next != start

		}
	}

}

func main() {
	maze := readMaze("high/bfs/maze/maze.in")
	for _, row := range maze {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
	walk(maze,
		point{0, 0},
		point{len(maze) - 1, len(maze[0]) - 1})
}
