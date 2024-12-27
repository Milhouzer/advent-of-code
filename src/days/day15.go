package days

import (
	"adventofcode/src/mathematics"
	"adventofcode/src/utils"
	"fmt"
	"strings"
)

type Day15 struct {
	DayN
	warehouse *WarehouseWoes
}

type Element int

const (
	OBSTACLE Element = iota
	BOX
	EMPTY
	BOX_LEFT
	BOX_RIGHT
)

type WarehouseWoes struct {
	Elements map[Vector2]Element
	Current  Vector2
	Commands string
	Size     Vector2
}

func (w *WarehouseWoes) Print() {
	grid := make([][]rune, int(w.Size.Y))
	for i := range grid {
		grid[i] = make([]rune, int(w.Size.X))
		for j := range grid[i] {
			grid[i][j] = '.'
		}
	}

	for pos, element := range w.Elements {
		switch element {
		case BOX:
			grid[int(pos.X)][int(pos.Y)] = 'O'
		case OBSTACLE:
			grid[int(pos.X)][int(pos.Y)] = '#'
		}
	}

	grid[int(w.Current.X)][int(w.Current.Y)] = '@'

	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func (w *WarehouseWoes) PrintUpscaled() {
	grid := make([][]rune, int(w.Size.X))
	for i := range grid {
		grid[i] = make([]rune, int(w.Size.Y))
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = '.'
		}
	}

	for pos, element := range w.Elements {
		switch element {
		case BOX_LEFT:
			grid[int(pos.X)][int(pos.Y)] = '['
		case BOX_RIGHT:
			grid[int(pos.X)][int(pos.Y)] = ']'
		case OBSTACLE:
			grid[int(pos.X)][int(pos.Y)] = '#'
		}
	}

	grid[int(w.Current.X)][int(w.Current.Y)] = '@'

	for _, row := range grid {
		fmt.Println(string(row))
	}
}

func (w *WarehouseWoes) AreaType(pos Vector2) Element {
	elem, v := w.Elements[pos]
	if v {
		return elem
	}

	return EMPTY
}

// pushing a row of boxes is equivalent to teleporting the first box at the end of the line if the end of the line is empty.
func (w *WarehouseWoes) PushBox(pos, dir Vector2) bool {
	elem, v := w.Elements[pos]
	if !(v && elem == BOX) {
		return true
	}

	canPush := true
	shouldMove := true
	nextPos := pos.Add(&dir)
	for {
		areaType := w.AreaType(*nextPos)
		switch areaType {
		case OBSTACLE:
			shouldMove = false
			canPush = false
		case BOX:
			nextPos = nextPos.Add(&dir)
		case EMPTY:
			delete(w.Elements, pos)
			w.Elements[*nextPos] = BOX
			canPush = false
		}

		if !canPush {
			return shouldMove
		}
	}
}

func (w *WarehouseWoes) PushBoxUpscaled(pos, dir Vector2, pushed *[]Vector2) bool {
	elem, v := w.Elements[pos]
	if !(v && (elem == BOX_LEFT || elem == BOX_RIGHT)) {
		return false
	}

	var otherPart *Vector2
	if elem == BOX_LEFT {
		right := Moves['>']
		otherPart = pos.Add(&right)
	} else if elem == BOX_RIGHT {
		left := Moves['<']
		otherPart = pos.Add(&left)
	}

	if dir == Moves['^'] || dir == Moves['v'] {
		nextPos := pos.Add(&dir)
		otherNextPos := otherPart.Add(&dir)

		el1, v1 := w.Elements[*nextPos]
		el2, v2 := w.Elements[*otherNextPos]
		// fmt.Printf("Bot or up move: %v, %v, %v, %v, %v: %v, %v\r\n", dir, pos, otherPart, nextPos, otherNextPos, el1, el2)

		// No boxes/obstacles above, box can be pushed directly without consequences
		if !v1 && !v2 {
			*pushed = append(*pushed, pos)
			*pushed = append(*pushed, *otherPart)
			return true
		}

		if v1 && !v2 {
			// Boxes cannot be pushed at all
			if el1 == OBSTACLE {
				return false
			}
			// Boxes are aligned :
			// [][]			[][]
			// []     or  	[]
			// @.			.@
			// if el1 == elem {
			*pushed = append(*pushed, pos)
			*pushed = append(*pushed, *otherPart)
			return w.PushBoxUpscaled(*nextPos, dir, pushed)
			// }
		}

		if !v1 && v2 {
			// Boxes cannot be pushed at all
			if el2 == OBSTACLE {
				return false
			}

			// Boxes are aligned :
			// .[]	 		[].
			// []     or   	 []
			// @.			 .@
			// if el2 == elem {
			*pushed = append(*pushed, pos)
			*pushed = append(*pushed, *otherPart)
			return w.PushBoxUpscaled(*otherNextPos, dir, pushed)
			// }
		}

		if v1 && v2 {
			// Boxes cannot be pushed at all
			if el1 == OBSTACLE || el2 == OBSTACLE {
				return false
			}

			// Boxes are aligned :
			// [][]			[][]
			//  []		or 	 []
			//  @.			 .@
			*pushed = append(*pushed, pos)
			*pushed = append(*pushed, *otherPart)
			return w.PushBoxUpscaled(*nextPos, dir, pushed) && w.PushBoxUpscaled(*otherNextPos, dir, pushed)
		}

	} else if dir == Moves['>'] || dir == Moves['<'] {
		// Get next pos after upscaled box
		nextPos := pos.Add(&dir)
		otherNextPos := nextPos.Add(&dir)
		el1, v1 := w.Elements[*otherNextPos]

		// No boxes/obstacles in line, box can be pushed directly without consequences
		// .[]<boxes>@ or @<boxes>[].
		if !v1 {
			*pushed = append(*pushed, pos)
			*pushed = append(*pushed, *otherPart)
			// fmt.Printf("Pushed0: %v\r\n", pushed)
			return true
		} else {
			// Boxes cannot be pushed at all
			// #[]<boxes>@ or @<boxes>[]#
			if el1 == OBSTACLE {
				// fmt.Printf("Pushed: %v\r\n", pushed)
				return false
			}

			// A box is present after this one
			// [][]<boxes>@ or @<boxes>[][]
			if el1 == elem {
				*pushed = append(*pushed, pos)
				*pushed = append(*pushed, *otherPart)
				// fmt.Printf("Pushed 2: %v\r\n", pushed)
				return w.PushBoxUpscaled(*nextPos, dir, pushed)
			}
			// fmt.Printf("Pushed prout: %v\r\n", pushed)
		}
	}

	fmt.Printf("Pushed end: %v\r\n", pushed)
	return false
}

func (w *WarehouseWoes) CalculateCoords() int {
	total := float64(0)
	for pos, elem := range w.Elements {
		if elem != BOX {
			continue
		}

		total += 100*pos.X + pos.Y
	}

	return int(total)
}

func (w *WarehouseWoes) CalculateUpscaledCoords() int {
	total := float64(0)
	for pos, elem := range w.Elements {
		if elem == BOX_LEFT {
			total += pos.X*100 + pos.Y
		}
	}

	return int(total)
}

var (
	Moves map[rune]Vector2 = map[rune]mathematics.Vector2{
		MOVE_EAST:  Vector2{X: 0, Y: 1},  // >
		MOVE_NORTH: Vector2{X: -1, Y: 0}, // v
		MOVE_WEST:  Vector2{X: 0, Y: -1}, // <
		MOVE_SOUTH: Vector2{X: 1, Y: 0},  // ^
	}

	MovesUpscaled map[rune]Vector2 = map[rune]mathematics.Vector2{
		MOVE_EAST:  Vector2{X: 0, Y: 0.5},  // >
		MOVE_NORTH: Vector2{X: -0.5, Y: 0}, // v
		MOVE_WEST:  Vector2{X: 0, Y: -0.5}, // <
		MOVE_SOUTH: Vector2{X: 0.5, Y: 0},  // ^
	}
)

func (d *Day15) Preprocess(path string) error {
	warehouse := &WarehouseWoes{}
	warehouse.Elements = make(map[mathematics.Vector2]Element)
	lines := utils.ReadLines(path)
	cmdIndex := 0
	for x, line := range lines {
		if line == "" {
			cmdIndex = x
			break
		}

		for y, b := range line {
			pos := Vector2{X: float64(x), Y: float64(y)}
			switch b {
			case WAREHOUSE_OBSTACLE:
				warehouse.Elements[pos] = OBSTACLE
			case WAREHOUSE_BOX:
				warehouse.Elements[pos] = BOX
			case WAREHOUSE_ROBOT:
				warehouse.Current = pos
			}
		}
	}

	cmdLine := lines[cmdIndex:]
	warehouse.Size = mathematics.Vector2{X: float64(cmdIndex), Y: float64(len(lines[0]))}

	warehouse.Commands = strings.Join(cmdLine, "")

	d.warehouse = warehouse
	return nil
}

func (d *Day15) PreprocessUpscaled(path string) error {
	warehouse := &WarehouseWoes{}
	warehouse.Elements = make(map[mathematics.Vector2]Element)
	lines := utils.ReadLines(path)
	cmdIndex := 0
	for x, line := range lines {
		if line == "" {
			cmdIndex = x
			break
		}

		for y, b := range line {
			posLeft := Vector2{X: float64(x), Y: float64(2 * y)}
			posRight := Vector2{X: float64(x), Y: float64(2*y + 1)}
			switch b {
			case WAREHOUSE_OBSTACLE:
				warehouse.Elements[posLeft] = OBSTACLE
				warehouse.Elements[posRight] = OBSTACLE
			case WAREHOUSE_BOX:
				warehouse.Elements[posLeft] = BOX_LEFT
				warehouse.Elements[posRight] = BOX_RIGHT
			case WAREHOUSE_ROBOT:
				warehouse.Current = posLeft
			}
		}
	}

	cmdLine := lines[cmdIndex:]
	warehouse.Size = mathematics.Vector2{X: float64(cmdIndex), Y: float64(len(lines[0]) * 2)}

	warehouse.Commands = strings.Join(cmdLine, "")

	d.warehouse = warehouse
	return nil
}

func (d *Day15) Solve(path string) {
	fmt.Println("PART 1: Initial state:")
	d.warehouse.PrintUpscaled()
	fmt.Println("-----------------------")
	for _, cmd := range d.warehouse.Commands {
		move := Moves[cmd]
		nextPos := d.warehouse.Current.Add(&move)
		areaType := d.warehouse.AreaType(*nextPos)
		switch areaType {
		case EMPTY:
			d.warehouse.Current = *nextPos
		case BOX:
			shouldMove := d.warehouse.PushBox(*nextPos, move)
			if shouldMove {
				d.warehouse.Current = *nextPos
			}
		case OBSTACLE:
		}

		// fmt.Println("Command: ", string(cmd))
		// d.warehouse.Print()
		// fmt.Println("-----------------------")
	}

	d.Pt1Sol = d.warehouse.CalculateCoords()

	d.PreprocessUpscaled(path)
	fmt.Println("PART 2: Initial state:")
	d.warehouse.PrintUpscaled()
	fmt.Println("-----------------------")

	for _, cmd := range d.warehouse.Commands {
		// if i == 60 {
		// 	return
		// }
		dir := Moves[cmd]
		nextPos := d.warehouse.Current.Add(&dir)
		areaType := d.warehouse.AreaType(*nextPos)
		moved := make([]Vector2, 0)
		switch areaType {
		case EMPTY:
			d.warehouse.Current = *nextPos
		case BOX_LEFT:
			shouldMove := d.warehouse.PushBoxUpscaled(*nextPos, dir, &moved)
			if shouldMove {
				d.warehouse.Current = *nextPos
				moved_unique := Unique(moved)
				m := make(map[Vector2]Element, 0)
				for _, v := range moved_unique {
					el := d.warehouse.Elements[v]
					new := v.Add(&dir)
					m[*new] = el
					delete(d.warehouse.Elements, v)
				}
				for p, v := range m {
					d.warehouse.Elements[p] = v
				}
				// fmt.Printf("Moved: %v\r\n", moved_unique)
				// d.warehouse.PrintUpscaled()
			}
		case BOX_RIGHT:
			shouldMove := d.warehouse.PushBoxUpscaled(*nextPos, dir, &moved)

			if shouldMove {
				moved_unique := Unique(moved)
				m := make(map[Vector2]Element, 0)
				for _, v := range moved_unique {
					el := d.warehouse.Elements[v]
					new := v.Add(&dir)
					m[*new] = el
					delete(d.warehouse.Elements, v)
				}
				for p, v := range m {
					d.warehouse.Elements[p] = v
				}
				d.warehouse.Current = *nextPos
				// fmt.Printf("Moved: %v\r\n", moved_unique)
			}
		case OBSTACLE:
		}
		// ##....[][]@
		// fmt.Println("Command: ", string(cmd))
		// d.warehouse.Print()
		// fmt.Println("-----------------------")
		// if i >= 125 && i <= 150 {
		// d.warehouse.PrintUpscaled()
		// }
	}

	d.warehouse.PrintUpscaled()
	d.Pt2Sol = d.warehouse.CalculateUpscaledCoords()
}

func Unique[T comparable](arr []T) []T {
	unique := make([]T, 0)
	m := make(map[T]struct{}, 0)
	for _, v := range arr {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
			unique = append(unique, v)
		}
	}

	return unique
}
