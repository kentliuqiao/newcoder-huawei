package main

import (
	"fmt"
)

type point struct {
	x, y int
}

var directions = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

func (p point) add(r point) point {
	return point{r.x + p.x, r.y + p.y}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.x < 0 || p.x >= len(grid) {
		return 0, false
	}
	if p.y < 0 || p.y >= len(grid[p.x]) {
		return 0, false
	}

	return grid[p.x][p.y], true
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, dir := range directions {
			next := cur.add(dir)

			val, ok := next.at(maze)
			if !ok || val == 1 { // 碰壁
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 { // 老路
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.x][next.y] = curSteps + 1

			Q = append(Q, next)
		}
	}

	steps[0][0] = 0
	return steps
}

func path(steps [][]int, start, end point) []point {
	res := []point{end}

	Q := []point{end}

	cur := end
	last, _ := cur.at(steps)
	for len(Q) > 0 {

		cur := Q[0]
		Q = Q[1:]

		if cur == start {
			break
		}

		for _, dir := range directions {
			next := cur.add(dir)
			step, ok := next.at(steps)
			if ok && step == last-1 {
				res = append(res, next)
				Q = append(Q, next)
				last = step
			}

		}
	}

	return res
}

func main() {
	for {
		row, col := 0, 0
		if _, err := fmt.Scan(&row, &col); err != nil {
			break
		}
		maze := make([][]int, row)
		for i := range maze {
			maze[i] = make([]int, col)
		}
		for i := 0; i < row; i++ {
			for j := 0; j < col; j++ {
				fmt.Scan(&maze[i][j])
			}
		}

		steps := walk(maze, point{0, 0}, point{row - 1, col - 1})

		path := path(steps, point{0, 0}, point{row - 1, col - 1})
		path[len(path)-1] = point{0, 0}

		for i := len(path) - 1; i >= 0; i-- {
			fmt.Printf("(%d,%d)\n", path[i].x, path[i].y)
		}
	}
}
