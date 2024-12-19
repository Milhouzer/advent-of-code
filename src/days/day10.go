package days

import (
	"adventofcode/src/geometry"
	"adventofcode/src/utils"
	"fmt"
	"strconv"
)

type Day10 struct {
	utils.DayN
	world  [][]int
	bounds geometry.WorldBounds
}

var _ utils.Day = (*Day10)(nil)

var (
	hikeDir = []Vector3{
		geometry.Vector3{X: 1, Y: 0, Z: 0},
		geometry.Vector3{X: 0, Y: 1, Z: 0},
		geometry.Vector3{X: -1, Y: 0, Z: 0},
		geometry.Vector3{X: 0, Y: -1, Z: 0},
	}
)

func (d *Day10) Preprocess(path string) error {
	lines := utils.ReadFile(path)
	numRows := len(lines)
	numCols := len(lines[0])
	d.world = make([][]int, numRows)
	for i := 0; i < numRows; i++ {
		d.world[i] = make([]int, numCols)
		for j := 0; j < numCols; j++ {
			d.world[i][j], _ = strconv.Atoi(string(lines[i][j]))
		}
	}

	d.bounds = geometry.WorldBounds{
		Up:    float64(len(d.world)),
		Bot:   0,
		Right: float64(len(d.world[0])),
		Left:  0,
	}

	return nil
}

func (d *Day10) Solve(path string) {
	total := 0
	for i := 0; i < len(d.world); i++ {
		for j := 0; j < len(d.world[i]); j++ {
			if d.world[i][j] == 0 {
				ends := make(map[Vector3]struct{}, 0)
				vecs := d.isHikeTrail(float64(i), float64(j), 0, 9)
				for _, v := range vecs {
					ends[v] = struct{}{}
				}

				total += len(ends)
			}
		}
	}
	d.Pt1Sol = total

	total = 0
	for i := 0; i < len(d.world); i++ {
		for j := 0; j < len(d.world[i]); j++ {
			fmt.Printf("Hike %v\n", d.world[i][j])
			if d.world[i][j] == 0 {
				vecs := d.isHikeTrail(float64(i), float64(j), 0, 9)
				total += len(vecs)
			}
		}
	}

	d.Pt2Sol = total
}

func (d *Day10) isHikeTrail(i, j float64, start, target int) []Vector3 {
	if start == 9 {
		return nil
	}
	score := start + 1
	pos := &Vector3{X: i, Y: j, Z: 0}
	ends := make([]Vector3, 0)
	for {
		paths := make([]Vector3, 0)
		for _, dir := range hikeDir {
			nextPos := pos.Add(&dir)
			if !d.bounds.Contains(nextPos) {
				continue
			}

			height := d.world[int(nextPos.X)][int(nextPos.Y)]
			if height == score {
				paths = append(paths, dir)
			}
		}

		if score == target {
			for _, dir := range paths {
				ends = append(ends, *pos.Add(&dir))
			}
			return ends
		}

		if (len(paths)) == 0 {
			return nil
		} else {
			for _, dir := range paths {
				ends = append(ends, d.isHikeTrail(pos.X+dir.X, pos.Y+dir.Y, score, target)...)
			}
			return ends
		}
	}
}
