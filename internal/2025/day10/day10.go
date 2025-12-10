package day10

import (
	"fmt"
	"math"
	"slices"
	"strings"

	"github.com/afonsocraposo/advent-of-code/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code/internal/utils/runner"
	"github.com/draffensperger/golp"
)

const year = 2025
const day = 10

func Main() {
	r := runner.New(year, day, part1, part2)
	r.TestPart1(1, 1)
	r.RunPart1(1)

	r.TestPart2(1, 2)
	r.RunPart2(1)
}

func parseIndicator(indicator string) matrix.Vector {
	n := len(indicator) - 2
	v := matrix.NewEmptyVector(n)
	for i, c := range indicator[1 : len(indicator)-1] {
		if c == '#' {
			v.Set(i, 1)
		}
	}
	return v
}

func parseButton(buttons string, length int) matrix.Vector {
	v, err := matrix.ParseVector(buttons[1:len(buttons)-1], ",")
	if err != nil {
		panic(err)
	}
	result := matrix.NewEmptyVector(length)
	for _, i := range v.Values {
		result.Set(i, 1)
	}
	return result
}

func parseJoltage(joltage string, length int) matrix.Vector {
	v, err := matrix.ParseVector(joltage[1:len(joltage)-1], ",")
	if err != nil {
		panic(err)
	}
	result := matrix.NewEmptyVector(length)
	for i, joltage := range v.Values {
		result.Set(i, joltage)
	}
	return result
}

func checkIndicator(indicator matrix.Vector, output matrix.Vector) bool {
	for i := range indicator.Size() {
		if indicator.Values[i] != output.Values[i]%2 {
			return false
		}
	}
	return true
}

func part1(lines []string) string {
	solution := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		indicatorString := parts[0]
		buttonsString := parts[1 : len(parts)-1]

		indicator := parseIndicator(indicatorString)

		buttons := make([]matrix.Vector, len(buttonsString))
		for i, b := range buttonsString {
			button := parseButton(b, indicator.Size())
			buttons[i] = button
		}

		m := matrix.NewMatrix(buttons)
		t, err := m.Transpose()
		if err != nil {
			panic(err)
		}

		combinations := matrix.GenerateBinarySequences(len(buttons))
		nc, _ := combinations.Size()
		buttonPresses := make([]int, nc)
		for i, row := range combinations.Rows {
			presses := row.Reduce(func(a, b int) int { return a + b }, 0)
			buttonPresses[i] = presses
		}

		tc, _ := combinations.Transpose()

		result, err := t.Multiply(*tc)
		if err != nil {
			panic(err)
		}

		minPresses := 10
		for i, presses := range buttonPresses {
			col, err := result.Column(i)
			if err != nil {
				panic(err)
			}
			if presses > minPresses || !checkIndicator(indicator, *col) {
				continue
			}
			minPresses = presses
		}
		solution += minPresses

	}

	return fmt.Sprintf("%d", solution)
}

func solve(mat matrix.Matrix, result matrix.Vector) int {
	_, n := mat.Size() // number of variables
	lp := golp.NewLP(0, n)

	// minimize a+b+c+d...
	obj := slices.Repeat([]float64{-1}, n)
	lp.SetObjFn(obj)
	lp.SetMaximize()

	// add constraints (each row of the matrix)
	for i, row := range mat.Rows {
		fl := make([]float64, row.Size())
		for j, v := range row.Values {
			fl[j] = float64(v)
		}
		lp.AddConstraint(fl, golp.EQ, float64(result.Values[i]))
	}

	// non negative variables
	for i := 1; i < n; i++ {
		lp.SetBounds(i, 0, 10000) // variable i has lower bound 0
		lp.SetInt(i, true)
	}

	lp.Solve()
	sol := 0
	for _, v := range lp.Variables() {
		sol += int(math.Round(v))
	}

	return sol
}

func part2(lines []string) string {
	solution := 0

	for _, line := range lines {
		parts := strings.Split(line, " ")
		indicatorString := parts[0]
		buttonsString := parts[1 : len(parts)-1]
		joltageString := parts[len(parts)-1]

		indicator := parseIndicator(indicatorString)

		buttons := make([]matrix.Vector, len(buttonsString))
		for i, b := range buttonsString {
			button := parseButton(b, indicator.Size())
			buttons[i] = button
		}
		joltage := parseJoltage(joltageString, indicator.Size())

		m := matrix.NewMatrix(buttons)
		t, err := m.Transpose()
		if err != nil {
			panic(err)
		}
		result := solve(*t, joltage)
		solution += result
	}

	return fmt.Sprintf("%d", solution)
}
