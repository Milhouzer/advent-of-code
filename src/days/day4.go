package days

import (
	"adventofcode/src/utils"
	"log"
	"time"
)

type Day4 struct {
	utils.DayN
	matrix [][]float32
}

var _ utils.Day = (*Day4)(nil)

/**************************************/
/* Elegant solution using convolution */
/* 				  WIP				  */
/**************************************/

// Problem can be solved using convolution matrices
var (
	Xb = 1 / float32(88)
	Mb = 1 / float32(77)
	Ab = 1 / float32(65)
	Sb = 1 / float32(83)

	XMASKernal = [3][3]float32{
		{1 / Mb, 0, 1 / Mb},
		{0, 1 / Ab, 0},
		{1 / Sb, 0, 1 / Sb},
	}

	TestMatrix = [][]float32{
		{Mb, Mb, Mb, Xb, Mb, Xb, Mb, Xb, Mb},
		{Xb, Ab, Xb, Ab, Xb, Ab, Sb, Ab, Mb},
		{Sb, Mb, Sb, Mb, Sb, Ab, Ab, Sb, Mb},
		{Xb, Xb, Ab, Xb, Sb, Ab, Ab, Sb, Mb},
		{Xb, Sb, Xb, Sb, Sb, Ab, Ab, Sb, Mb},
		{Mb, Mb, Mb, Xb, Sb, Ab, Ab, Sb, Mb},
		{Mb, Mb, Mb, Xb, Sb, Ab, Ab, Sb, Mb},
		{Mb, Mb, Mb, Xb, Sb, Ab, Ab, Sb, Mb},
		{Mb, Mb, Mb, Xb, Sb, Ab, Ab, Sb, Mb},
	}
)

// Convert data to byte matrix
func convertToMatrix(data []string) [][]float32 {
	numRows := len(data)
	numCols := len(data[0])
	matrix := make([][]float32, numRows)
	for i := 0; i < numRows; i++ {
		matrix[i] = make([]float32, numCols)
		for j := 0; j < numCols; j++ {
			matrix[i][j] = float32(data[i][j])
		}
	}

	return matrix
}

// 2D Convolution function designed specifically for finding "XMAS" patterns
//
// Convolution output using the test matrix: elements equal to exactly 5 are where X-MAS patterns are detected.
// This solution does not currently solve the problem
//
// [[5 4.7694807 5 4.8434815 5.276923 4.8100557 5.354845]
// [4.814261 5.0387273 4.871079 5.3724685 5.0796337 5.646154 5.065688]
// [4.480422 5 4.53724 5.244671 5.3892493 5.3892493 5.3225927]
// [4.9985924 4.5097404 4.9733806 5.062853 5.3892493 5.3892493 5.3225927]
// [4.75 4.7206817 4.619269 5.115564 5.3892493 5.3892493 5.3225927]
// [5 4.7402596 4.744269 5.062853 5.3892493 5.3892493 5.3225927]
// [5 4.7402596 4.744269 5.062853 5.3892493 5.3892493 5.3225927]
// [4.7402596 4.744269 5.062853 5.3892493 5.3892493 5.3225927]
// [5 4.7402596 4.744269 5.062853 5.3892493 5.3892493 5.3225927]]
func xmasConv(input [][]float32) [][]float32 {
	rows := len(input)
	cols := len(input[0])
	kRows := len(XMASKernal)
	kCols := len(XMASKernal[0])

	output := make([][]float32, rows-kRows+1)
	for i := range output {
		output[i] = make([]float32, cols-kCols+1)
	}

	for i := 0; i < len(output); i++ {
		for j := 0; j < len(output[i]); j++ {
			sum := float32(0)
			for ki := 0; ki < kRows; ki++ {
				for kj := 0; kj < kCols; kj++ {
					sum += input[i+ki][j+kj] * XMASKernal[ki][kj]
				}
			}
			output[i][j] = sum
		}
	}

	return output
}

/**********************************/
/* Brute force/Recursive solution */
/**********************************/

var (
	directions = [][2]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
		{1, -1},
		{1, 1},
		{-1, -1},
		{-1, 1},
	}
	word = "XMAS"
)

func (d *Day4) Preprocess(path string) error {
	lines := utils.ReadFile(path)
	d.matrix = convertToMatrix(lines)
	// log.Printf("%v", d.matrix)
	output := xmasConv(TestMatrix)
	log.Printf("output: %v", output)
	return nil
}

func (d *Day4) Solve(path string) {
	lines := utils.ReadFile(path)

	start := time.Now()
	count := 0
	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[0]); j++ {
			c := lines[i][j]
			if c == 'X' {
				for _, dir := range directions {
					if d.IsXMAS(lines, [2]int{i, j}, dir, 1) {
						count++
					}
				}
			}
		}
	}

	elapsed := time.Since(start)
	log.Printf("day4 pt1 took %s", elapsed)
	d.Pt1Sol = count

	start = time.Now()
	count = 0
	// avoid edges, 'A' can't be centered on edges, this allows us to not check boundaries in verification function
	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[0])-1; j++ {
			c := lines[i][j]
			if c == 'A' {
				if d.IsXPattern(lines, i, j) {
					count++
				}
			}
		}
	}

	elapsed = time.Since(start)
	log.Printf("day4 pt2 took %s", elapsed)
	d.Pt2Sol = count
}

func (d *Day4) IsXMAS(table []string, pos [2]int, dir [2]int, c int) bool {
	// next position in grid
	pos = [2]int{pos[0] + dir[0], pos[1] + dir[1]}

	// check if next pos is in table bounds
	if !d.IsInBounds(table, pos) {
		return false
	}

	res := table[pos[0]][pos[1]] == word[c]
	if res {
		// recursivity break condition
		if c == 3 {
			return true
		}
		return d.IsXMAS(table, pos, dir, c+1)
	}
	return false
}

// ------- 0 	  1 	   2  	  	3
//
// 0.1  M.S   	 S.M 	  M.M      S.S
// .p.  .A.  or  .A.  or  .A.  or  .A.
// 2.3	M.S		 S.M	  S.S	   M.M
func (d *Day4) IsXPattern(table []string, i, j int) bool {
	// 0 or 2
	if table[i-1][j-1] == 'M' && table[i+1][j+1] == 'S' {
		return (table[i+1][j-1] == 'M' && table[i-1][j+1] == 'S') || (table[i+1][j-1] == 'S' && table[i-1][j+1] == 'M')
	}

	// 1 or 3
	if table[i-1][j-1] == 'S' && table[i+1][j+1] == 'M' {
		return (table[i+1][j-1] == 'S' && table[i-1][j+1] == 'M') || (table[i+1][j-1] == 'M' && table[i-1][j+1] == 'S')
	}

	return false
}

func (d *Day4) IsInBounds(table []string, pos [2]int) bool {
	return (pos[0] >= 0 && pos[0] < len(table) && pos[1] >= 0 && pos[1] < len(table[0]))
}

func Add(a, b [2]int) [2]int {
	return [2]int{a[0] + b[0], a[1] + b[1]}
}
