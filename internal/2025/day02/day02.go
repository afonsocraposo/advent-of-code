package day00

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/afonsocraposo/advent-of-code/internal/utils/runner"
)

const year = 2025
const day = 2

func Main() {
	r := runner.New(year, day, part1, part2)
	r.TestPart1(1, 1)
	r.RunPart1(1)

	r.TestPart2(1, 2)
	r.RunPart2(1)
}

func isInvalid1(n int) bool {
	str := strconv.Itoa(n)
	l := len(str)
	if l%2 != 0 {
		return false
	}
	if str[:l/2] == str[l/2:] {
		return true
	}
	return false
}

func isInvalid2(n int) bool {
	str := strconv.Itoa(n)
	l := len(str)

	for size := 1; size <= l/2; size++ {
		if l%size != 0 {
			continue
		}
		pattern := str[:size]
		invalid := true
		for i := size; i <= l-size; i += size {
			p := str[i : i+size]
			if pattern != p {
				invalid = false
				break
			}
		}
		if invalid {
			return true
		}
	}

	return false
}

func getInvalidIds(start int, end int, method int) []int {
	invalid := []int{}
	for i := start; i <= end; i++ {
		if method == 1 {
			if isInvalid1(i) {
				invalid = append(invalid, i)
			}
		} else {
			if isInvalid2(i) {
				invalid = append(invalid, i)
			}
		}
	}
	return invalid
}

func part1(lines []string) string {
	solution := 0

	line := lines[0]
	for _, r := range strings.Split(line, ",") {
		parts := strings.Split(r, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalln(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalln(err)
		}
		invalid := getInvalidIds(start, end, 1)
		for _, v := range invalid {
			solution += v
		}
	}

	return fmt.Sprintf("%d", solution)
}

func part2(lines []string) string {
	solution := 0

	line := lines[0]
	for _, r := range strings.Split(line, ",") {
		parts := strings.Split(r, "-")
		start, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatalln(err)
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatalln(err)
		}
		invalid := getInvalidIds(start, end, 2)
		for _, v := range invalid {
			solution += v
		}
	}

	return fmt.Sprintf("%d", solution)
}
