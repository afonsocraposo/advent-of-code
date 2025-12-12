package year2025

import (
	day01 "github.com/afonsocraposo/advent-of-code/internal/2025/day01"
	day02 "github.com/afonsocraposo/advent-of-code/internal/2025/day02"
	day03 "github.com/afonsocraposo/advent-of-code/internal/2025/day03"
	day04 "github.com/afonsocraposo/advent-of-code/internal/2025/day04"
	day05 "github.com/afonsocraposo/advent-of-code/internal/2025/day05"
	day06 "github.com/afonsocraposo/advent-of-code/internal/2025/day06"
	day07 "github.com/afonsocraposo/advent-of-code/internal/2025/day07"
	day08 "github.com/afonsocraposo/advent-of-code/internal/2025/day08"
	day09 "github.com/afonsocraposo/advent-of-code/internal/2025/day09"
	// day10 "github.com/afonsocraposo/advent-of-code/internal/2025/day10" // disabled because of golp dependency
	day11 "github.com/afonsocraposo/advent-of-code/internal/2025/day11"
	day12 "github.com/afonsocraposo/advent-of-code/internal/2025/day12"
)

var Days = map[int]func(){
	1: day01.Main,
	2: day02.Main,
	3: day03.Main,
	4: day04.Main,
	5: day05.Main,
	6: day06.Main,
	7: day07.Main,
	8: day08.Main,
	9: day09.Main,
	// 10: day10.Main, // disabled because of golp dependecy
	11: day11.Main,
	12: day12.Main,
}
