package days

import (
	"adventofcode/src/mathematics"
	"adventofcode/src/utils"
	"fmt"
	"slices"
)

type Vector3 = mathematics.Vector3
type World = mathematics.World
type Transform = mathematics.Transform

type Day6 struct {
	DayN
	world *World
	agent Transform
}

var _ Day = (*Day6)(nil)

var forbidden = 0
var subAgent Transform

var newObstacles = make(map[Vector3]struct{})

const (
	MAX_IT = 150000
)

// Create problem associated world
func (d *Day6) Preprocess(path string) error {
	lines := utils.ReadLines(path)
	world, startPos := mathematics.WorldFromFile(lines)

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
		subAgent = mathematics.Transform{
			Position: pos,
			Rotation: rot,
		}
		obstacles := append(world.Obstacles(), *next)
		visited := make(map[Vector3][]Vector3)
		// for k, v := range world.Visited() {
		// 	visited[k] = v
		// }
		if checkLoop(obstacles, visited, world.Bounds, &subAgent) {
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

func checkLoop(obstacles []Vector3, visited map[Vector3][]Vector3, bounds mathematics.WorldBounds, tr *Transform) bool {
	for {
		next := tr.Position.Add(&tr.Rotation)
		if !bounds.Contains(next) {
			return false
		}

		v, ok := visited[tr.Position]
		if ok {
			if slices.Contains(v, tr.Rotation) {
				return true
			}
		}
		visited[tr.Position] = append(visited[tr.Position], tr.Rotation)
		if slices.Contains(obstacles, *next) {
			tr.RotateRight()
			next = tr.Position.Add(&tr.Rotation)
		}

		tr.Position = *next
	}
}
