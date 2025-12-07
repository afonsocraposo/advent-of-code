package day00

import (
	"fmt"
	"log"
	"slices"

	"github.com/afonsocraposo/advent-of-code/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code/internal/utils/runner"
)

const year = 2025
const day = 7

func Main() {
	r := runner.New(year, day, part1, part2)
	r.TestPart1(1, 1)
	r.RunPart1(1)

	r.TestPart2(1, 2)
	r.RunPart2(1)
}

func part1(lines []string) string {
	solution := 0

	mat := matrix.ParseRuneMatrix(lines)
	m, n := mat.Size()
	j := slices.Index(mat.Rows[0].Values, 'S')
	beams := []int{j}
	for i := 1; i < m; i++ {
		newBeams := map[int]bool{}
		for _, j := range beams {
			v, err := mat.Get(i, j)
			if err != nil {
				log.Fatalln(err)
			}
			if v == '^' {
				solution++
				if j-1 >= 0 {
					newBeams[j-1] = true
				}
				if j+1 < n {
					newBeams[j+1] = true
				}
			} else {
				newBeams[j] = true
			}
		}
		beams = make([]int, len(newBeams))

		i := 0
		for j := range newBeams {
			beams[i] = j
			i++
		}
	}

	return fmt.Sprintf("%d", solution)
}

func part2(lines []string) string {
	solution := 0

	mat := matrix.ParseRuneMatrix(lines)
	m, n := mat.Size()
	j := slices.Index(mat.Rows[0].Values, 'S')
	timelines := make([]int, n)
	timelines[j] = 1
	for i := 1; i < m; i++ {
		for j, v := range timelines {
			if j == 0 {
				continue
			}
			char, err := mat.Get(i, j)
			if err != nil {
				log.Fatalln(err)
			}
			if char != '^' {
				continue
			}
			if j-1 >= 0 {
				timelines[j-1] += v
			}
			if j+1 < n {
				timelines[j+1] += v
			}
			timelines[j] = 0
		}
	}

	for _, value := range timelines {
		solution += value
	}

	return fmt.Sprintf("%d", solution)
}
