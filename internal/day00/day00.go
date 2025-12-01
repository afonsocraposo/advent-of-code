package day00

import (
	"fmt"
	"github.com/afonsocraposo/advent-of-code/internal/utils/runner"
)

const year = 2025
const day = 0

func Main() {
	r := runner.New(year, day, part1, part2)
	r.TestPart1(1, 1)
	r.RunPart1(1)

	r.TestPart2(1, 2)
	r.RunPart2(1)
}

func part1(lines []string) string {
	solution := 0

	// TODO

	return fmt.Sprintf("%d", solution)
}

func part2(lines []string) string {
	solution := 0

	// TODO

	return fmt.Sprintf("%d", solution)
}
