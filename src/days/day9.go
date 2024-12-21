package days

import (
	"adventofcode/src/utils"
	"fmt"
	"strconv"
	"strings"
)

type Day9 struct {
	DayN
}

var _ Day = (*Day9)(nil)

func (d *Day9) Solve(path string) {
	line := utils.ReadContent(path)
	if len(line)%2 == 0 {
		panic("line length must be odd.")
	}

	s := make([]int, len(line))
	for i := 0; i < len(line); i++ {
		s[i], _ = strconv.Atoi(string(line[i]))
	}

	// checksum1 := d.solvePt1(s, len(line))
	// d.Pt1Sol = checksum1

	checksum2 := d.solvePt2(s)
	d.Pt2Sol = checksum2
}

func getHoles(s []int) (holes, holes_index, data, data_index []int, decompiled_index int) {
	// .H.H.H.H.H.H.H.H.H.
	// 2333133121414131402
	// fill holes and holes_index, calculate s_q_last_index
	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			data = append(data, s[i])
			data_index = append(data_index, decompiled_index)
		} else {
			holes = append(holes, s[i])
			holes_index = append(holes_index, decompiled_index)
		}
		decompiled_index += s[i]
	}
	return holes, holes_index, data, data_index, decompiled_index
}

func (d *Day9) solvePt2(s []int) int {
	holes, holes_index, data, data_index, decompiled_index := getHoles(s)
	// log.Printf("%v\n%v\n%v\n%v\n", holes, holes_index, data, data_index)
	reconstructed := make([]int, decompiled_index)
	for i := len(data) - 1; i >= 0; i-- {
		d := data[i]
		// fmt.Printf("Process %d: %d\n", i, d)
		for j := 0; j < i+1; j++ {
			hole := holes[j]
			index := holes_index[j]
			// no hole was found to be filled
			if j == i {
				index := data_index[j]
				for k := 0; k < d; k++ {
					reconstructed[index+k] = i
				}
				// fmt.Printf("%v \n%v\n%v\n", reconstructed, holes, holes_index)
				break
			}

			if hole >= d {
				for k := 0; k < d; k++ {
					reconstructed[index+k] = i
				}

				holes[j] = holes[j] - d
				holes_index[j] = holes_index[j] + d
				// fmt.Printf("%v \n%v\n%v\n", reconstructed, holes, holes_index)
				break
			}
		}
	}

	checksum := 0
	for i := 0; i < len(reconstructed); i++ {
		checksum += i * reconstructed[i]
	}
	// fmt.Printf("%v", reconstructed)
	return checksum
}

func (d *Day9) solvePt1(s []int, lineLen int) int {
	sb := &strings.Builder{}
	p := 0
	q := lineLen - 1
	p_weight := 0
	q_weight := q / 2
	s_p := s[p]
	s_q := s[q]
	fillq := false
	checksum := 0
	occ := 0
	for {
		// Break condition
		if p >= q {
			for {
				checksum += occ * q_weight
				sb.WriteString(fmt.Sprintf("%d", q_weight))
				s_q--

				if s_q == 0 {
					break
				}
				occ++
			}
			break
		}
		if fillq {
			checksum += occ * q_weight
			sb.WriteString(fmt.Sprintf("%d", q_weight))
			s_q--
			s_p--
			if s_p == 0 {
				for {
					p++
					if p%2 == 0 {
						p_weight++
					}
					s_p = s[p]
					if s_p != 0 {
						fillq = false
						break
					}
				}
			}
			if s_q == 0 {
				for {
					q -= 2
					q_weight--
					s_q = s[q]
					if s_q != 0 {
						break
					}
				}
			}
		} else {
			checksum += occ * p_weight
			sb.WriteString(fmt.Sprintf("%d", p_weight))
			s_p--
			if s_p == 0 {
				for {
					p++
					if p%2 == 0 {
						p_weight++
					}
					s_p = s[p]
					fillq = !fillq
					if s_p != 0 {
						break
					}
				}
			}
		}
		occ++
	}

	// str := sb.String()

	// fmt.Printf("Compiled line of lenght %d: %s\n", len(line), line)
	// fmt.Printf("Compiled line of lenght %d: %s\n", len(line), str)
	// fmt.Printf("Compiled line of lenght %d: %s\n", len(line), line[:100])
	// fmt.Printf("Compiled line of lenght %d: %s\n", len(line), line[len(line)-100:])
	// fmt.Printf("Compiled str of lenght %d: %s\n", len(str), str[:100])
	// fmt.Printf("Compiled str of lenght %d: %s\n", len(str), str[len(str)-100:])

	return checksum
	// 6346871685398
	// 6373055193464

}
