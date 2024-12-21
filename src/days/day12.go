package days

import (
	"adventofcode/src/utils"
	"fmt"
	"math"
)

type Day12 struct {
	utils.DayN
	tiles [][]byte
	areas map[area]struct{}
}

type Vector2 struct {
	I int
	J int
}

// Add other to a [Vector2]
func (v *Vector2) Add(other *Vector2) *Vector2 {
	return &Vector2{I: v.I + other.I, J: v.J + other.J}
}

func (v *Vector2) Manhattan() int {
	return int(math.Abs(float64(v.I))) + int(math.Abs(float64(v.J)))
}

func (v *Vector2) IsZero() bool {
	return v.I == 0 && v.J == 0
}

func (v *Vector2) InBounds(i, j int) bool {
	return v.I >= 0 && v.I < i && v.J >= 0 && v.J < j
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
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
	}

	north     = Vector2{I: -1, J: 0}
	south     = Vector2{I: 1, J: 0}
	east      = Vector2{I: 0, J: 1}
	west      = Vector2{I: 0, J: -1}
	northEast = Vector2{I: -1, J: 1}
	northWest = Vector2{I: -1, J: -1}
	southEast = Vector2{I: 1, J: 1}
	southWest = Vector2{I: 1, J: -1}
)

func (d *Day12) Preprocess(path string) error {
	lines := utils.ReadFile(path)
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
			_, ok := locallyExplored[Vector2{I: i, J: j}]
			if ok {
				continue
			}

			// New area detected
			area := &area{
				Byte: d.tiles[i][j],
			}
			d.marchArea(i, j, d.tiles[i][j], locallyExplored, area, corners)
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

func (d *Day12) marchArea(i, j int, b byte, m map[Vector2]byte, area *area, corners map[Corner]struct{}) {

	// initial perimeter is 4
	perimeter := 4

	// increase surface by 1
	area.Surface++

	// mark init pos as visited
	pos := Vector2{I: i, J: j}
	angles := d.angles(&pos, b)
	for _, corner := range angles {
		corners[corner] = struct{}{}
	}
	m[pos] = b

	for _, dir := range gridDirs {
		pos := Vector2{I: i + dir.I, J: j + dir.J}
		_, ok := m[pos]
		// position has already been visited and is inside the area
		if ok && m[pos] == b {
			perimeter--
			continue
		}

		// position is out of bounds
		if !(pos.I >= 0 && pos.I < len(d.tiles) && pos.J >= 0 && pos.J < len(d.tiles[0])) {
			continue
		}

		// position is inside the area, march and reduce perimeter
		if d.tiles[pos.I][pos.J] == b {
			perimeter--
			d.marchArea(pos.I, pos.J, b, m, area, corners)
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

	return d.tiles[pos.I][pos.J] == b
}
