package day01

import (
	"fmt"
	"log"
	"strconv"

	"github.com/afonsocraposo/advent-of-code/internal/utils/runner"
)

const year = 2025
const day = 1

var examples = []int{1}

func Main() {
	r := runner.New(year, day, part1, part2)
	r.TestPart1(1, 1)
	r.RunPart1(1)

	r.TestPart2(1, 2)
	r.RunPart2(1)
}

func part1(lines []string) string {
	solution := 0
	dial := 50
	for _, line := range lines {
		rotations, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalln(err)
		}
		if line[0] == 'R' {
			dial += rotations % 100
		} else {
			dial += ((-rotations % 100) + 100)
		}
		dial %= 100
		if dial == 0 {
			solution++
		}
	}
	return fmt.Sprintf("%d", solution)
}

func part2(lines []string) string {
	solution := 0
	dial := 50
	for _, line := range lines {
		rotations, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalln(err)
		}
		if line[0] == 'R' {
			dial += rotations
			solution += dial / 100
		} else {
			dial -= rotations
			if dial < 0 && dial+rotations > 0 {
				// the dial went from positive to negative
				solution++
			}
			solution += dial / -100
		}
		if dial == 0 {
			solution++
		}
		dial = ((dial % 100) + 100) % 100
	}
	return fmt.Sprintf("%d", solution)
}
