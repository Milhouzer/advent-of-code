package days

import (
	"adventofcode/src/utils"
	"fmt"
	"math"
)

type Day14 struct {
	DayN
	robots Robots
}

var (
	day14_pattern = "p=%d,%d v=%d,%d"
)

func (d *Day14) Preprocess(path string) error {
	robots := Robots{
		AreaX: 101,
		AreaY: 103,
	}
	lines := utils.ReadLines(path)
	for _, line := range lines {
		var posX, posY, velX, velY int
		_, err := fmt.Sscanf(line, day14_pattern, &posX, &posY, &velX, &velY)
		if err != nil {
			return err
		}

		pos := Vector2{
			I: float64(posX),
			J: float64(posY),
		}

		vel := Vector2{
			I: float64(velX),
			J: float64(velY),
		}
		robots.Add(pos, vel)
	}

	d.robots = robots

	return nil
}

func (d *Day14) Solve(path string) {
	min := math.MaxInt64
	itMin := 0
	robotsState := Robots{
		AreaX: d.robots.AreaX,
		AreaY: d.robots.AreaY,
	}
	for i := 0; i < 6475; i++ {
		d.robots.Tick()
		s := d.robots.SafetyFactor()
		if s < min {
			min = s
			itMin = i
			fmt.Printf("New minimum: %d at %d\r\n", min, itMin)
			robotsState.Positions = d.robots.Positions
		}
	}

	robotsState.Print()
	d.Pt1Sol = d.robots.SafetyFactor()
}

type Robots struct {
	Positions  []Vector2
	Velocities []Vector2
	AreaX      int
	AreaY      int
}

func (r *Robots) Add(pos, vel Vector2) {
	r.Positions = append(r.Positions, pos)
	r.Velocities = append(r.Velocities, vel)
}

func (r *Robots) Count() int {
	if len(r.Positions) != len(r.Velocities) {
		panic("robots pos/vel misaligned")
	}

	return len(r.Positions)
}

func (r *Robots) Tick() {
	for i := 0; i < r.Count(); i++ {
		r.Positions[i] = wrap(*r.Positions[i].Add(&r.Velocities[i]), 0, r.AreaX, 0, r.AreaY)
	}
}

func (r *Robots) Print() {
	matrix := make([][]rune, r.AreaY)
	for i := range matrix {
		matrix[i] = make([]rune, r.AreaX)
		for j := range matrix[i] {
			matrix[i][j] = '.'
		}
	}

	for _, pos := range r.Positions {
		if pos.I >= 0 && pos.I < float64(r.AreaX) && pos.J >= 0 && pos.J < float64(r.AreaY) {
			matrix[int(pos.J)][int(pos.I)] = '1'
		}
	}

	for _, row := range matrix {
		for _, cell := range row {
			fmt.Printf("%c ", cell)
		}
		fmt.Println()
	}
}

func (r *Robots) SafetyFactor() int {
	var quadrants [4]int
	for _, pos := range r.Positions {
		q := getQuadrant(pos, r.AreaX, r.AreaY)
		if q == -1 {
			continue
		}

		quadrants[q]++
	}

	// fmt.Printf("Quadrants: %v\r\n", quadrants)
	return quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3]
}

// 0 | 1
// -----
// 2 | 3
func getQuadrant(pos Vector2, sizeX, sizeY int) int {
	middleX, middleY := int(sizeX-1)/2, int(sizeY-1)/2
	if pos.I == float64(middleX) || pos.J == float64(middleY) {
		return -1
	}
	// top half
	if pos.J < float64(middleY) {
		if pos.I < float64(middleX) {
			return 0
		} else {
			return 1
		}
	} else {
		if pos.I < float64(middleX) {
			return 2
		} else {
			return 3
		}
	}
}

func wrap(v Vector2, xa, xb, ya, yb int) Vector2 {
	width := xb - xa
	height := yb - ya

	if width <= 0 || height <= 0 {
		panic("Invalid range: b must be greater than a and d must be greater than c")
	}

	v.I = float64(((int(v.I)-xa)%width + width) % width)
	v.J = float64(((int(v.J)-ya)%height + height) % height)

	return v
}
