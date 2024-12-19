package days

import (
	"adventofcode/src/utils"
	"fmt"
)

type Day12 struct {
	utils.DayN
	areas [][]byte
}

type Vector2 struct {
	I int
	J int
}

var (
	gridDirs = []Vector2{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
)

func (d *Day12) Preprocess(path string) error {
	lines := utils.ReadFile(path)
	x, y := len(lines), len(lines[0])
	d.areas = make([][]byte, x)
	for i := 0; i < x; i++ {
		d.areas[i] = make([]byte, y)
		for j := 0; j < y; j++ {
			d.areas[i][j] = lines[i][j]
		}
	}
	return nil
}

func (d *Day12) Solve(path string) {
	// explored := make(map[Vector2]struct{})
	// for i := 0; i < len(d.areas); i++ {
	// 	for j := 0; j < len(d.areas[i]); j++ {
	// 		_, ok := explored[Vector2{I: i, J: j}]
	// 		if ok {
	// 			continue
	// 		}

	// 		locallyExplored := make(map[Vector2]struct{})
	// 		d.computeMetrics(i, j, d.areas[i][j], locallyExplored)
	// 		fmt.Printf("%v", locallyExplored)
	// 	}
	// }

	locallyExplored := make(map[Vector2]struct{})
	d.computeMetrics(0, 0, d.areas[0][0], locallyExplored)
	fmt.Printf("%v", locallyExplored)
}

func (d *Day12) computeMetrics(i, j int, b byte, m map[Vector2]struct{}) (int, int) {
	perimeter, area := 4, 0
	for _, dir := range gridDirs {
		pos := Vector2{I: i + dir.I, J: j + dir.J}
		_, ok := m[pos]
		if ok || !(pos.I >= 0 && pos.I < len(d.areas) && pos.J >= 0 && pos.J < len(d.areas[0])) {
			continue
		}

		m[pos] = struct{}{}
		if d.areas[pos.I][pos.J] == b {
			d.computeMetrics(pos.I, pos.J, b, m)
			perimeter--
			area += a
		}
	}

	return perimeter, area
}
