package day00

import (
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"

	"github.com/afonsocraposo/advent-of-code/internal/utils/numbers"
	"github.com/afonsocraposo/advent-of-code/internal/utils/runner"
)

const year = 2025
const day = 5

func Main() {
	r := runner.New(year, day, part1, part2)
	r.TestPart1(1, 1)
	r.RunPart1(1)

	r.TestPart2(2, 1)
	r.RunPart2(2)
}

type Range struct {
	Start int
	End   int
}

func (r *Range) Includes(n int) bool {
	return n >= r.Start && n <= r.End
}

func (r *Range) Count() int {
	return r.End - r.Start + 1
}

type Evaluator struct {
	Ranges []Range
}

func (e *Evaluator) AddRange(start, end int, group bool) {
	if !group {
		e.Ranges = append(e.Ranges, Range{start, end})
	} else {
		for i, r := range e.Ranges {
			if r.Includes(start) {
				e.Ranges = slices.Delete(e.Ranges, i, i+1)
				e.AddRange(r.Start, numbers.IntMax(r.End, end), true)
				return
			}
			if r.Includes(end) {
				e.Ranges = slices.Delete(e.Ranges, i, i+1)
				e.AddRange(numbers.IntMin(r.Start, start), r.End, true)
				return
			}
		}
		e.Ranges = append(e.Ranges, Range{start, end})
	}
}

func (e *Evaluator) MergeRanges() {
	i := 0
	for i < len(e.Ranges) {
		r := e.Ranges[i]
		l1 := len(e.Ranges)
		e.Ranges = slices.Delete(e.Ranges, i, i+1)
		e.AddRange(r.Start, r.End, true)
		l2 := len(e.Ranges)
		if l1 == l2 {
			i++
		}
	}
}

func (e *Evaluator) IsFresh(n int) bool {
	for _, r := range e.Ranges {
		if r.Includes(n) {
			return true
		}
	}
	return false
}

func NewEvaluator() *Evaluator {
	return &Evaluator{[]Range{}}
}

func part1(lines []string) string {
	solution := 0

	freshList := true
	evaluator := NewEvaluator()
	for _, line := range lines {
		if line == "" {
			freshList = false
			continue
		}
		if freshList {
			parts := strings.Split(line, "-")
			start, err := strconv.Atoi(parts[0])
			if err != nil {
				log.Fatalln(err)
			}
			end, err := strconv.Atoi(parts[1])
			if err != nil {
				log.Fatalln(err)
			}
			evaluator.AddRange(start, end, false)
		} else {
			n, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalln(err)
			}
			if evaluator.IsFresh(n) {
				solution++
			}
		}
	}

	return fmt.Sprintf("%d", solution)
}

func part2(lines []string) string {
	solution := 0

	evaluator := NewEvaluator()
	for _, line := range lines {
		parts := strings.Split(line, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalln(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalln(err)
		}
		evaluator.AddRange(start, end, true)
		evaluator.MergeRanges()
	}
	for _, r := range evaluator.Ranges {
		solution += r.Count()
	}

	return fmt.Sprintf("%d", solution)
}
