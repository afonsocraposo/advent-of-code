package day03

import (
	"fmt"
	"log"

	"github.com/afonsocraposo/advent-of-code/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code/internal/utils/runner"
)

const year = 2025
const day = 3

func Main() {
	r := runner.New(year, day, part1, part2)
	r.TestPart1(1, 1)
	r.RunPart1(1)

	r.TestPart2(1, 2)
	r.RunPart2(1)
}

func getLargestJoltage1(line string) int {
	vector, err := matrix.ParseVector(line, "")
	if err != nil {
		log.Fatalln(err)
	}
	m := -1
	indexes := []int{}
	for index, value := range vector.Values[:vector.Size()-1] {
		if value > m {
			m = value
			indexes = []int{index}
		} else if value == m {
			indexes = append(indexes, index)
		}
	}

	d1 := m
	d2 := 0
	for _, v := range vector.Values[indexes[0]+1:] {
		if v > d2 {
			d2 = v
		}
	}
	return d1*10 + d2
}

func getMaxAndPositions(values []int) (int, []int) {
	m := -1
	indexes := []int{}
	for index, value := range values {
		if value > m {
			m = value
			indexes = []int{index}
		} else if value == m {
			indexes = append(indexes, index)
		}
	}
	return m, indexes
}

func getLargestJoltage2(values []int, batteries []int) int {
	m, positions := getMaxAndPositions(values[:len(values)-(11-len(batteries))])
	batteries = append(batteries, m)

	solution := 0
	if len(batteries) == 12 {
		solution := 0
		for _, battery := range batteries {
			solution = solution*10 + battery
		}
		return solution
	}
	for _, pos := range positions {
		joltage := getLargestJoltage2(values[pos+1:], batteries)
		if joltage > solution {
			solution = joltage
		}
	}
	return solution
}

func part1(lines []string) string {
	solution := 0

	for _, line := range lines {
		joltage := getLargestJoltage1(line)
		solution += joltage
	}

	return fmt.Sprintf("%d", solution)
}

func part2(lines []string) string {
	solution := 0

	for _, line := range lines {
		vector, err := matrix.ParseVector(line, "")
		if err != nil {
			log.Fatalln(err)
		}
		joltage := getLargestJoltage2(vector.Values, []int{})
		solution += joltage
	}

	return fmt.Sprintf("%d", solution)
}
