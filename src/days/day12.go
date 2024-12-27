package days

import (
	"adventofcode/src/mathematics"
	"adventofcode/src/utils"
	"fmt"
)

type Vector2 = mathematics.Vector2

type Day12 struct {
	DayN
	tiles [][]byte
	areas map[area]struct{}
}

type area struct {
	Byte      byte
	Perimeter int
	Surface   int
	Corners   int
}

func (a *area) Cost() int {
	return a.Perimeter * a.Surface
}

func (a *area) Discount() int {
	return a.Surface * a.Corners
}

func (a *area) String() string {
	return fmt.Sprintf("Area %s: Surface: %d, Perimeter: %d, Corners: %d, Cost: %d, Discount: %d", string(a.Byte), a.Surface, a.Perimeter, a.Corners, a.Cost(), a.Discount())
}

type Corner struct {
	Pos  Vector2
	Rank int
}

var (
	gridDirs = []Vector2{
		{X: 0, Y: 1},
		{X: 0, Y: -1},
		{X: -1, Y: 0},
		{X: 1, Y: 0},
	}

	north     = Vector2{X: -1, Y: 0}
	south     = Vector2{X: 1, Y: 0}
	east      = Vector2{X: 0, Y: 1}
	west      = Vector2{X: 0, Y: -1}
	northEast = Vector2{X: -1, Y: 1}
	northWest = Vector2{X: -1, Y: -1}
	southEast = Vector2{X: 1, Y: 1}
	southWest = Vector2{X: 1, Y: -1}
)

func (d *Day12) Preprocess(path string) error {
	lines := utils.ReadLines(path)
	x, y := len(lines), len(lines[0])
	d.tiles = make([][]byte, x)
	for i := 0; i < x; i++ {
		d.tiles[i] = make([]byte, y)
		for j := 0; j < y; j++ {
			d.tiles[i][j] = lines[i][j]
		}
	}
	return nil
}

func (d *Day12) Solve(path string) {
	total := 0
	discount := 0
	d.areas = map[area]struct{}{}
	locallyExplored := make(map[Vector2]byte)
	corners := make(map[Corner]struct{})
	lastLen := 0
	for i := 0; i < len(d.tiles); i++ {
		for j := 0; j < len(d.tiles[i]); j++ {
			_, ok := locallyExplored[Vector2{X: float64(i), Y: float64(j)}]
			if ok {
				continue
			}

			// New area detected
			area := &area{
				Byte: d.tiles[i][j],
			}
			d.marchArea(float64(i), float64(j), d.tiles[i][j], locallyExplored, area, corners)
			newLen := len(corners)
			added := newLen - lastLen
			lastLen = newLen
			total += area.Cost()
			discount += area.Surface * added
			// fmt.Printf("%v\n", area)
			// fmt.Printf("Corners count: %d\n", len(corners))
			d.areas[*area] = struct{}{}
		}
	}
	d.Pt1Sol = total
	d.Pt2Sol = discount
}

func (d *Day12) marchArea(i, j float64, b byte, m map[Vector2]byte, area *area, corners map[Corner]struct{}) {

	// initial perimeter is 4
	perimeter := 4

	// increase surface by 1
	area.Surface++

	// mark init pos as visited
	pos := Vector2{X: i, Y: j}
	angles := d.angles(&pos, b)
	for _, corner := range angles {
		corners[corner] = struct{}{}
	}
	m[pos] = b

	for _, dir := range gridDirs {
		pos := Vector2{X: i + dir.X, Y: j + dir.Y}
		_, ok := m[pos]
		// position has already been visited and is inside the area
		if ok && m[pos] == b {
			perimeter--
			continue
		}

		// position is out of bounds
		if !(pos.X >= 0 && pos.X < float64(len(d.tiles)) && pos.Y >= 0 && pos.Y < float64(len(d.tiles[0]))) {
			continue
		}

		// position is inside the area, march and reduce perimeter
		if d.tiles[int(pos.X)][int(pos.Y)] == b {
			perimeter--
			d.marchArea(pos.X, pos.Y, b, m, area, corners)
		}
	}

	// add to area perimeter
	// area.Corners = len(corners)
	area.Perimeter += perimeter
}

func (d *Day12) angles(pos *Vector2, b byte) (corners []Corner) {

	n := pos.Add(&north)
	s := pos.Add(&south)
	e := pos.Add(&east)
	w := pos.Add(&west)

	ne := pos.Add(&northEast)
	se := pos.Add(&southEast)
	nw := pos.Add(&northWest)
	sw := pos.Add(&southWest)

	/*
	   oE	oE	oX
	   EX	EE	X*
	*/
	if !d.insideArea(e, b) && !d.insideArea(s, b) {
		corner := Corner{
			Pos:  *se,
			Rank: 0,
		}
		// fmt.Printf("0 Pos: %v Corner %s: %v\n", pos, string(b), corner)
		corners = append(corners, corner)
	} else if (d.insideArea(e, b) && d.insideArea(s, b)) && !d.insideArea(se, b) {
		corner := Corner{
			Pos:  *se,
			Rank: 0,
		}
		// fmt.Printf("1 Corner %s: %v\n", string(b), corner)
		corners = append(corners, corner)
	}

	/*
	   Eo	Eo	Xo
	   EE	XE	*X
	*/
	if !d.insideArea(s, b) && !d.insideArea(w, b) {
		corner := Corner{
			Pos:  *sw,
			Rank: 1,
		}
		// fmt.Printf("2 Corner %s: %v\n", string(b), corner)
		corners = append(corners, corner)
	} else if (d.insideArea(s, b) && d.insideArea(w, b)) && !d.insideArea(sw, b) {
		corner := Corner{
			Pos:  *sw,
			Rank: 1,
		}
		// fmt.Printf("3 Corner %s: %v\n", string(b), corner)
		corners = append(corners, corner)
	}

	/*
	   EE	XE	*X
	   Eo	Eo	Xo
	*/
	if !d.insideArea(n, b) && !d.insideArea(w, b) {
		corner := Corner{
			Pos:  *nw,
			Rank: 2,
		}
		// fmt.Printf("4 Corner %s: %v\n", string(b), corner)
		corners = append(corners, corner)
	} else if (d.insideArea(n, b) && d.insideArea(w, b)) && !d.insideArea(nw, b) {
		corner := Corner{
			Pos:  *nw,
			Rank: 2,
		}
		// fmt.Printf("5 Corner %s: %v\n", string(b), corner)
		corners = append(corners, corner)
	}

	/*
	   EE	XE	X*
	   oE	oE	oX
	*/
	if !d.insideArea(n, b) && !d.insideArea(e, b) {
		corner := Corner{
			Pos:  *ne,
			Rank: 3,
		}
		// fmt.Printf("6 Corner %s: %v\n", string(b), corner)
		corners = append(corners, corner)
	} else if (d.insideArea(n, b) && d.insideArea(e, b)) && !d.insideArea(ne, b) {
		corner := Corner{
			Pos:  *ne,
			Rank: 3,
		}
		// fmt.Printf("7 Corner %s: %v\n", string(b), corner)
		corners = append(corners, corner)
	}

	return
}

func (d *Day12) insideArea(pos *Vector2, b byte) bool {
	li, lj := len(d.tiles), len(d.tiles[0])
	if !pos.InBounds(li, lj) {
		return false
	}

	return d.tiles[int(pos.X)][int(pos.Y)] == b
}
