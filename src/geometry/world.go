package geometry

import (
	"adventofcode/src/utils"
	"cmp"
	"fmt"
	"math"
	"slices"
	"strings"
)

var (
	WorldUp = &Vector3{X: 0, Y: 0, Z: 1}
)

const (
	POUND_SYMBOL        = '#'
	DOT_SYMBOL          = '.'
	AGENT_SYMBOL        = '^'
	TRAIL_SYMBOL        = '+'
	VIRTUAL_OBSTRUCTION = 'O'
)

type Segment struct {
	Origin Vector3
	End    Vector3
}

func (s *Segment) Direction() *Vector3 {
	v := s.End.Subtract(&s.Origin)
	return v.Normalize()
}

// Given three collinear points p, q, r, the function checks if
// point q lies on line segment 'pr'
func onSegment(p, q, r Vector3) bool {
	if q.X <= math.Max(p.X, r.X) && q.X >= math.Min(p.X, r.X) &&
		q.Y <= math.Max(p.Y, r.Y) && q.Y >= math.Min(p.Y, r.Y) {
		return true
	}

	return false
}

// To find orientation of ordered triplet (p, q, r).
// The function returns following values
// 0 --> p, q and r are collinear
// 1 --> Clockwise
// 2 --> Counterclockwise
func orientation(p, q, r Vector3) int {
	val := (q.Y-p.Y)*(r.X-q.X) - (q.X-p.X)*(r.Y-q.Y)

	if val == 0 {
		return 0
	}

	if val > 0 {
		return 1
	}

	return 2
}

func doIntersect(p1, q1, p2, q2 Vector3) bool {
	// special cases
	o1 := orientation(p1, q1, p2)
	o2 := orientation(p1, q1, q2)
	o3 := orientation(p2, q2, p1)
	o4 := orientation(p2, q2, q1)

	// General case
	if o1 != o2 && o3 != o4 {
		return true
	}

	// p1, q1 and p2 are collinear and p2 lies on segment p1q1
	if o1 == 0 && onSegment(p1, p2, q1) {
		return true
	}

	// p1, q1 and q2 are collinear and q2 lies on segment p1q1
	if o2 == 0 && onSegment(p1, q2, q1) {
		return true
	}

	// p2, q2 and p1 are collinear and p1 lies on segment p2q2
	if o3 == 0 && onSegment(p2, p1, q2) {
		return true
	}

	// p2, q2 and q1 are collinear and q1 lies on segment p2q2
	if o4 == 0 && onSegment(p2, q1, q2) {
		return true
	}

	return false
}

type WorldBounds struct {
	Up    float64
	Bot   float64
	Left  float64
	Right float64
}

func (b *WorldBounds) Contains(pos *Vector3) bool {
	return pos.X >= b.Left && pos.X < b.Right && pos.Y >= b.Bot && pos.Y < b.Up
}

type World struct {
	sizeX int
	sizeY int
	// Visited maps visited positions with directions
	visited   map[Vector3][]Vector3
	obstacles []Vector3
	Bounds    WorldBounds
}

func WorldFromFile(lines []string) (*World, Vector3) {
	sX := len(lines[0])
	sY := len(lines)
	world := &World{
		sizeX:     sX,
		sizeY:     sY,
		visited:   make(map[Vector3][]Vector3),
		obstacles: make([]Vector3, 0),
		Bounds: WorldBounds{
			Up:    float64(sY),
			Bot:   0,
			Left:  0,
			Right: float64(sX),
		},
	}

	startPos := Vector3{}

	for i := 0; i < sY; i++ {
		for j := 0; j < len(lines[i]); j++ {
			worldPos := Vector3{X: float64(j), Y: float64(sY - i - 1), Z: 0}
			s := lines[i][j]
			if s == POUND_SYMBOL {
				world.AddObstacle(worldPos)
			}
			if s == AGENT_SYMBOL {
				startPos = worldPos
			}
		}
	}

	return world, startPos
}

func (w *World) Obstacles() []Vector3 {
	return w.obstacles
}

func (w *World) AddObstacle(worldPos Vector3) {
	w.obstacles = append(w.obstacles, worldPos)
}

func (w *World) RemoveObstacle(worldPos Vector3) {
	w.obstacles = slices.DeleteFunc(w.obstacles, func(v Vector3) bool { return v.Equals(&worldPos) })
}

func (w *World) SizeX() int {
	return w.sizeX
}

func (w *World) SizeY() int {
	return w.sizeY
}

func (w *World) Visit(pos, rot Vector3) {
	w.visited[pos] = append(w.visited[pos], rot)
}

func (w *World) Visited() map[Vector3][]Vector3 {
	return w.visited
}

func (w *World) HasBeenVisited(pos Vector3) bool {
	_, ok := w.visited[pos]
	return ok
}

func (w *World) IsInBounds(pos *Vector3) bool {
	return w.Bounds.Contains(pos)
}

func (w *World) IsObstacle(pos Vector3) bool {
	return slices.Contains(w.obstacles, pos)
}

func Clamp[T cmp.Ordered](value, min, max T) T {
	if value < min {
		return min
	} else if value > max {
		return max
	}
	return value
}

func (w *World) DisplayAround(center Vector3, sX, sY int) {
	// if len(w.ForbiddenPosition) < 1800 {
	// 	return
	// }
	utils.ClearScreen()
	up := Clamp(int(center.Y)+sY/2, 0, w.SizeY())
	bot := Clamp(int(center.Y)-sY/2, 0, w.SizeY())
	left := Clamp(int(center.X)-sX/2, 0, w.SizeX())
	right := Clamp(int(center.X)+sX/2, 0, w.SizeX())

	// forbidden := slices.Collect(maps.Keys(w.ForbiddenPosition))
	for i := bot; i < bot+sY && i < up; i++ {
		line := []rune(strings.Repeat(" ", sX))
		for j := left; j < left+sX && j < right; j++ {
			pos := Vector3{X: float64(j), Y: float64(i), Z: 0}
			if v, ok := w.visited[pos]; ok {
				switch v[0] {
				case Vector3{1, 0, 0}:
					line[j] = '>'
				case Vector3{-1, 0, 0}:
					line[j] = '<'
				case Vector3{0, 1, 0}:
					line[j] = 'v'
				case Vector3{0, -1, 0}:
					line[j] = '^'
				}
			}
			// if slices.Contains(forbidden, pos) {
			// 	line[j] = 'O'
			// }
			if slices.Contains(w.obstacles, pos) {
				line[j] = '#'
			}
		}
		fmt.Println(string(line))
	}
	// fmt.Println(crosses)
}

type LineTraceResult struct {
	Hit   bool
	Trace Segment
	Steps []*Vector3
}

func LineTraceInBounds(bounds WorldBounds, pos, step Vector3, returnSteps bool) (res LineTraceResult) {
	currentPos := &pos
	if !bounds.Contains(currentPos) {
		res.Trace = Segment{
			Origin: *currentPos,
			End:    *currentPos,
		}
		return
	}
	if returnSteps {
		res.Steps = append(res.Steps, currentPos)
	}

	for {
		nextPos := currentPos.Add(&step)
		if !bounds.Contains(nextPos) {
			res.Trace = Segment{
				Origin: pos,
				End:    *currentPos,
			}
			return
		}
		currentPos = nextPos

		if returnSteps {
			res.Steps = append(res.Steps, currentPos)
		}

		res.Hit = true
	}
}
