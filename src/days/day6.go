package days

import (
	"adventofcode/src/geometry"
	"adventofcode/src/utils"
	"fmt"
	"slices"
)

type Vector3 = geometry.Vector3
type World = geometry.World
type Transform = geometry.Transform

type Day6 struct {
	utils.DayN
	world *World
	agent Transform
}

var _ utils.Day = (*Day6)(nil)

var forbidden = 0
var subAgent Transform

var newObstacles = make(map[Vector3]struct{})

const (
	POUND_SYMBOL        = '#'
	DOT_SYMBOL          = '.'
	AGENT_SYMBOL        = '^'
	TRAIL_SYMBOL        = '+'
	VIRTUAL_OBSTRUCTION = 'O'
	MAX_IT              = 150000
)

// Create problem associated world
func (d *Day6) Preprocess(path string) error {
	lines := utils.ReadFile(path)
	world, startPos := geometry.WorldFromFile(lines)

	d.world = world
	d.agent = Transform{
		Position: startPos,
		Rotation: Vector3{X: 0, Y: 1, Z: 0},
	}
	return nil
}

func (d *Day6) Solve(path string) {

	it := 0
	for {
		tick := march(d.world, &d.agent)
		if tick && it < MAX_IT {
			// d.world.DisplayAround(d.agent.Position, 150, 40)
		} else {
			break
		}
		it++
	}

	d.Pt1Sol = len(d.world.Visited())
	fmt.Println(d.Pt1Sol)
	d.Preprocess(path)

	it = 0
	for {
		tick := march_pt_2(d.world, &d.agent)
		if tick && it < MAX_IT {
			// d.world.DisplayAround(d.agent.Position, 150, 40)
		} else {
			break
		}
		it++
	}

	d.Pt2Sol = len(newObstacles)
	fmt.Printf("TOTAL: %d", forbidden)
}

func march(world *World, tr *Transform) bool {
	pos := tr.Position
	rot := tr.Rotation
	next := pos.Add(&rot)
	if !world.IsInBounds(next) {
		world.Visit(pos, rot)
		return false
	}

	world.Visit(pos, rot)
	if world.IsObstacle(*next) {
		tr.RotateRight()
		next = pos.Add(&tr.Rotation)
	}

	tr.Position = *next
	return true
}

func march_pt_2(world *World, tr *Transform) bool {
	pos := tr.Position
	rot := tr.Rotation
	next := pos.Add(&rot)
	if !world.IsInBounds(next) {
		world.Visit(pos, rot)
		return false
	}

	if !world.HasBeenVisited(*next) {
		subAgent = geometry.Transform{
			Position: pos,
			Rotation: rot,
		}
		if checkLoop(world, &subAgent) {
			newObstacles[pos] = struct{}{}
			forbidden++
		}
	}

	world.Visit(pos, rot)
	if world.IsObstacle(*next) {
		tr.RotateRight()
		next = pos.Add(&tr.Rotation)
	}

	tr.Position = *next
	return true
}

func checkLoop(world *World, tr *Transform) bool {
	encountered := make(map[Vector3][]Vector3)
	pos := tr.Position
	rot := tr.Rotation
	encountered[pos] = []Vector3{rot}
	tr.RotateRight()
	it := 0
	for {
		next := tr.Position.Add(&tr.Rotation)
		if !world.IsInBounds(next) {
			return false
		}

		if world.IsObstacle(*next) || (tr.Position == pos && it != 0) {
			tr.RotateRight()
			v, ok := encountered[*next]
			if ok {
				if slices.Contains(v, tr.Rotation) {
					world.ForbiddenPosition[pos] = struct{}{}
					return true
				}
				encountered[*next] = append(encountered[*next], tr.Rotation)
			} else {
				encountered[*next] = []Vector3{tr.Rotation}
			}
			next = tr.Position.Add(&tr.Rotation)
		}

		tr.Position = *next
		it++
	}
}
