package days

import (
	"adventofcode/src/geometry"
	"adventofcode/src/utils"
)

type Day8 struct {
	utils.DayN
	nodes     map[byte][]Vector3
	antinodes map[Vector3]struct{}
}

var _ utils.Day = (*Day8)(nil)

func (d *Day8) Solve(path string) {
	d.nodes = make(map[byte][]geometry.Vector3)
	d.antinodes = make(map[geometry.Vector3]struct{})
	lines := utils.ReadFile(path)
	x, y := len(lines), len(lines[0])
	bounds := geometry.WorldBounds{
		Up:    float64(y),
		Bot:   0,
		Left:  0,
		Right: float64(x),
	}

	for j := y - 1; j >= 0; j-- {
		for i := 0; i < x; i++ {
			s := lines[j][i]
			if s == DOT_SYMBOL {
				continue
			}
			worldPos := Vector3{X: float64(i), Y: float64(j), Z: 0}
			v, ok := d.nodes[s]
			if ok {
				d.nodes[s] = append(v, worldPos)
			} else {
				d.nodes[s] = []Vector3{worldPos}
			}
		}
	}

	d.solvePt1(bounds)
	d.Pt1Sol = len(d.antinodes)

	d.antinodes = make(map[geometry.Vector3]struct{})
	d.solvePt2(bounds)
	d.Pt2Sol = len(d.antinodes)
}

func (d *Day8) solvePt1(bounds geometry.WorldBounds) {
	for _, node := range d.nodes {
		for i := 0; i < len(node); i++ {
			for j := i + 1; j < len(node); j++ {
				n1, n2 := node[i], node[j]
				offset := n1.Subtract(&n2)
				an1, an2 := n1.Add(offset), n2.Subtract(offset)
				if bounds.Contains(an1) {
					d.antinodes[*an1] = struct{}{}
				}
				if bounds.Contains(an2) {
					d.antinodes[*an2] = struct{}{}
				}
			}
		}
	}
}

func (d *Day8) solvePt2(bounds geometry.WorldBounds) {
	var step *Vector3
	for _, node := range d.nodes {
		for i := 0; i < len(node); i++ {
			for j := i + 1; j < len(node); j++ {
				n1, n2 := node[i], node[j]
				step1, step2 := n1.Subtract(&n2), n2.Subtract(&n1)
				ray1, ray2 := geometry.LineTraceInBounds(bounds, n1, *step1, true), geometry.LineTraceInBounds(bounds, n2, *step2, true)
				for i := 0; i < len(ray1.Steps); i++ {
					step = ray1.Steps[i]
					d.antinodes[*step] = struct{}{}
				}
				for j := 0; j < len(ray2.Steps); j++ {
					step = ray2.Steps[j]
					d.antinodes[*step] = struct{}{}
				}
			}
		}
	}
}
