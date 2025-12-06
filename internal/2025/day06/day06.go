package day06

import (
	"fmt"
	"log"
	"strconv"

	"github.com/afonsocraposo/advent-of-code/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code/internal/utils/runner"
)

const year = 2025
const day = 6

func Main() {
	r := runner.New(year, day, part1, part2)
	r.TestPart1(1, 1)
	r.RunPart1(1)

	r.TestPart2(1, 2)
	r.RunPart2(1)
}

func part1(lines []string) string {
	solution := 0

	mat, err := matrix.ParseMatrix(lines[:len(lines)-1], " ")
	if err != nil {
		log.Fatalln(err)
	}

	v := matrix.ParseRuneVector(lines[len(lines)-1])
	ops := []int{}
	for _, val := range v.Values {
		if val != ' ' {
			ops = append(ops, val)
		}
	}

	_, n := mat.Size()
	for j := range n {
		col, _ := mat.Column(j)
		row := 0
		multiply := ops[j] == '*'
		if multiply {
			row = 1
		}
		for _, v := range col.Values {
			if multiply {
				row *= v
			} else {
				row += v
			}
		}
		solution += row
	}

	return fmt.Sprintf("%d", solution)
}

func part2(lines []string) string {
	solution := 0

	numStr := lines[:len(lines)-1]
	opStr := lines[len(lines)-1]

	numbers := [][]int{}
	col := 0
	colNum := []int{}
	for true {
		if col >= len(numStr[0]) {
			numbers = append(numbers, colNum)
			break
		}
		separator := true
		number := 0
		for _, row := range numStr {
			digit, err := strconv.Atoi(string(row[col]))
			if err != nil {
				continue
			}
			separator = false
			number = number*10 + digit
		}
		if separator {
			numbers = append(numbers, colNum)
			colNum = []int{}
		} else {
			colNum = append(colNum, number)
		}
		col++
	}

	v := matrix.ParseRuneVector(opStr)
	ops := []int{}
	for _, val := range v.Values {
		if val != ' ' {
			ops = append(ops, val)
		}
	}

	for i, op := range ops {
		multiply := op == '*'
		result := 0
		if multiply {
			result = 1
		}
		for _, val := range numbers[i] {
			if multiply {
				result *= val
			} else {
				result += val
			}
		}
		solution += result
	}

	return fmt.Sprintf("%d", solution)
}
