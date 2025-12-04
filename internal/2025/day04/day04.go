package day00

import (
	"fmt"
	"log"
	"strings"

	"github.com/afonsocraposo/advent-of-code/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code/internal/utils/runner"
)

const year = 2025
const day = 4

func Main() {
	r := runner.New(year, day, part1, part2)
	r.TestPart1(1, 1)
	r.RunPart1(1)

	r.TestPart2(1, 2)
	r.RunPart2(1)
}

func part1(lines []string) string {
	solution := 0

	linesN := make([]string, len(lines))
	for i, line := range lines {
		linesN[i] = strings.ReplaceAll(strings.ReplaceAll(line, "@", "1"), ".", "0")
	}
	mat, err := matrix.ParseMatrix(linesN, "")
	if err != nil {
		log.Fatalln(err)
	}
	kernel := matrix.NewMatrixWithValue(3, 3, 1)
	kernel.Set(1, 1, 0)

	convoluted, err := mat.Convolution(kernel)
	if err != nil {
		log.Fatalln(err)
	}

	m, n := mat.Size()
	for i := range m {
		for j := range n {
			value, err := mat.Get(i, j)
			if err != nil {
				log.Fatalln(err)
			}
			if value == 0 {
				continue
			}
			cv, err := convoluted.Get(i, j)
			if err != nil {
				log.Fatalln(err)
			}
			if cv < 4 {
				solution++
			}
		}
	}

	return fmt.Sprintf("%d", solution)
}

func part2(lines []string) string {
	solution := 0

	linesN := make([]string, len(lines))
	for i, line := range lines {
		linesN[i] = strings.ReplaceAll(strings.ReplaceAll(line, "@", "1"), ".", "0")
	}
	mat, err := matrix.ParseMatrix(linesN, "")
	if err != nil {
		log.Fatalln(err)
	}
	kernel := matrix.NewMatrixWithValue(3, 3, 1)
	kernel.Set(1, 1, 0)

	for true {
		convoluted, err := mat.Convolution(kernel)
		if err != nil {
			log.Fatalln(err)
		}

		m, n := mat.Size()
		removed := 0
		for i := range m {
			for j := range n {
				value, err := mat.Get(i, j)
				if err != nil {
					log.Fatalln(err)
				}
				if value == 0 {
					continue
				}
				cv, err := convoluted.Get(i, j)
				if err != nil {
					log.Fatalln(err)
				}
				if cv < 4 {
					removed++
					mat.Set(i, j, 0)
				}
			}
		}
		solution += removed
		if removed == 0 {
			break
		}
	}

	return fmt.Sprintf("%d", solution)
}
