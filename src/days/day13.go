package days

import (
	"adventofcode/src/mathematics"
	"adventofcode/src/utils"
	"fmt"
	"math/big"
)

type Day13 struct {
	DayN
	systems []system
}

type system struct {
	Buttons mathematics.Matrix2x2
	Prize   Vector2
}

type bigSystem struct {
	Buttons mathematics.Matrix2x2BigFloat
	Prize   mathematics.Vector2BigFloat
}

func (s *system) String() string {
	return fmt.Sprintf("%v, [%.0f, %.0f]", s.Buttons, s.Prize, s.Prize.Y)
}

func (s *system) FromString(lines []string) (system, error) {
	var buttonA_X, buttonA_Y, buttonB_X, buttonB_Y, prize_X, prize_Y float64

	_, err := fmt.Sscanf(lines[0], "Button A: X+%v, Y+%v", &buttonA_X, &buttonA_Y)
	if err != nil {
		fmt.Printf("Error parsing input %s: %v\r\n", lines[0], err)
		return system{}, err
	}
	_, err = fmt.Sscanf(lines[1], "Button B: X+%v, Y+%v", &buttonB_X, &buttonB_Y)
	if err != nil {
		fmt.Printf("Error parsing input %s: %v\r\n", lines[1], err)
		return system{}, err
	}
	_, err = fmt.Sscanf(lines[2], "Prize: X=%v, Y=%v", &prize_X, &prize_Y)
	if err != nil {
		fmt.Printf("Error parsing input %s: %v\r\n", lines[2], err)
		return system{}, err
	}

	return system{
		Buttons: mathematics.Matrix2x2{
			X1: buttonA_X,
			X2: buttonB_X,
			Y1: buttonA_Y,
			Y2: buttonB_Y,
		},
		Prize: mathematics.Vector2{
			X: prize_X,
			Y: prize_Y,
		},
	}, nil
}

func (s *system) Solve() (float64, error) {
	inverted, err := s.Buttons.Invert()
	if err != nil {
		return 0, err
	}

	sol := inverted.MulVector2(s.Prize)
	if !utils.IsFloatAnInt(sol.X) || !utils.IsFloatAnInt(sol.Y) {
		return 0, fmt.Errorf("system has no solution in |N")
	}

	return sol.X*3 + sol.Y, nil
}

func (s *bigSystem) Solve() (*big.Float, error) {
	inverted, err := s.Buttons.Invert()
	if err != nil {
		return nil, err
	}

	sol := inverted.MulVector2(s.Prize)
	a, va := mathematics.ToIntIfNear(sol.X)
	b, vb := mathematics.ToIntIfNear(sol.Y)
	if !(a && b) {
		return nil, fmt.Errorf("system has no solution in |N: %f, %f", sol.X, sol.Y)
	}

	fmt.Printf("Viable system: %f, %f", sol.X, sol.Y)
	res := new(big.Float).Mul(va, big.NewFloat(3))
	res = new(big.Float).Add(res, vb)
	return res, nil
}

var _ Day = (*Day13)(nil)

func (d *Day13) Solve(path string) {
	lines := utils.ReadLinesInPacks(path, 4)

	tokens := float64(0)
	for i := 0; i < len(lines); i++ {
		s := system{}
		s, err := s.FromString(lines[i])
		if err != nil {
			panic(fmt.Sprintf("can't parse system: %v", err))
		}
		sol, err := s.Solve()
		if err == nil {
			tokens += sol
			d.systems = append(d.systems, s)
		}
	}

	d.Pt1Sol = int(tokens)

	d.systems = make([]system, 0)
	offset, _ := new(big.Float).SetString("10000000000000")
	tokens2 := new(big.Float)
	for i := 0; i < len(lines); i++ {
		system := system{}
		system, err := system.FromString(lines[i])
		if err != nil {
			panic(fmt.Sprintf("can't parse system: %v", err))
		}
		bigsystem := bigSystem{
			Buttons: mathematics.BigFromMatrix2x2(system.Buttons),
			Prize: mathematics.Vector2BigFloat{
				X: new(big.Float).Add(offset, big.NewFloat(system.Prize.X)).SetPrec(128),
				Y: new(big.Float).Add(offset, big.NewFloat(system.Prize.Y)).SetPrec(128),
			},
		}

		sol, err := bigsystem.Solve()
		if err != nil {
			fmt.Printf("system %v (%v) cannot be solved: %v\r\n", bigsystem, sol, err)
		} else {
			tokens2 = new(big.Float).Add(tokens2, sol)
			fmt.Printf("System %.0f, %.0f solved, added: %.0f, %.0f\r\n", bigsystem.Prize.X, bigsystem.Prize.Y, sol, bigsystem.Buttons.X1)
		}
	}
	// 30626309651571 too low
	// 59769020323405 too low
	// 104878770860279 NO
	// 105620095782133 NO
	// 105620095782380
	// 165389115996228 too high ??
	// 165389116105505 ???
	// 165389116105952 too high
	fmt.Printf("Pt2 REAL ANSWER: %.0f\r\n", tokens2)
}
