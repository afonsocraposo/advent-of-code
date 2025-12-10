package day10

import (
	"fmt"
	"strings"

	"github.com/afonsocraposo/advent-of-code/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code/internal/utils/runner"
)

const year = 2025
const day = 10

func Main() {
	r := runner.New(year, day, part1, part2)
	r.TestPart1(1, 1)
	r.RunPart1(1)
	//
	// r.TestPart2(1, 2)
	// r.RunPart2(1)
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

func parseButton(button string, length int) matrix.Vector {
	v, err := matrix.ParseVector(button[1:len(button)-1], ",")
	if err != nil {
		panic(err)
	}
	result := matrix.NewEmptyVector(length)
	for _, i := range v.Values {
		result.Set(i, 1)
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

func part2(lines []string) string {
	solution := 0

	// TODO

	return fmt.Sprintf("%d", solution)
}
