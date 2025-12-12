package day12

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/afonsocraposo/advent-of-code/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code/internal/utils/runner"
)

const year = 2025
const day = 12

func Main() {
	r := runner.New(year, day, part1, nil)
	// r.TestPart1(1, 1)
	r.RunPart1(1)
}

type Region struct {
	M        int
	N        int
	Presents []int
}

func part1(lines []string) string {
	solution := 0

	isShape := false
	shape := []string{}
	shapes := []matrix.Matrix{}
	regions := []Region{}
	for _, line := range lines {
		if len(shapes) < 6 {
			if len(line) > 0 && (line[0] == '#' || line[0] == '.') {
				isShape = true
			} else {
				isShape = false
			}
			if isShape {
				shape = append(shape, line)
			} else {
				if len(shape) > 0 {
					mat := matrix.ParseRuneMatrix(shape)
					shapes = append(shapes, mat)
				}
				shape = []string{}
			}
		} else {
			parts := strings.Split(line, ": ")
			size := strings.Split(parts[0], "x")
			m := size[1]
			M, err := strconv.Atoi(m)
			if err != nil {
				panic(m)
			}
			n := size[0]
			N, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			presents := []int{}
			for v := range strings.SplitSeq(parts[1], " ") {
				number, err := strconv.Atoi(v)
				if err != nil {
					panic(err)
				}
				presents = append(presents, number)
			}
			regions = append(regions, Region{M, N, presents})
		}
	}

	for _, region := range regions {
		area := 0
		for _, count := range region.Presents {
			area += count * 9
		}
		available := region.M * region.N
		if area <= available {
			solution++
		}
	}

	return fmt.Sprintf("%d", solution)
}
