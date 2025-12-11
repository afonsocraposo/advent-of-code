package day11

import (
	"fmt"
	"strings"

	"github.com/afonsocraposo/advent-of-code/internal/utils/runner"
)

const year = 2025
const day = 11

func Main() {
	r := runner.New(year, day, part1, part2)
	r.TestPart1(1, 1)
	r.RunPart1(1)

	r.TestPart2(2, 2)
	r.RunPart2(1)
}

type Devices map[string][]string

func countPaths(devices Devices, current string) int {
	if current == "out" {
		return 1
	}
	outputs, ok := devices[current]
	if !ok {
		panic("Device doesn't exist on list of devices")
	}
	result := 0
	for _, output := range outputs {
		result += countPaths(devices, output)
	}
	return result
}

func part1(lines []string) string {
	solution := 0

	devices := Devices{}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		device := strings.Replace(parts[0], ":", "", 1)
		outputs := parts[1:]
		devices[device] = outputs
	}

	solution = countPaths(devices, "you")

	return fmt.Sprintf("%d", solution)
}

func paramsToKey(current string, dac, fft bool) string {
	return fmt.Sprintf("%s:%t:%t", current, dac, fft)
}

func countPathsWithDacAndFft(devices Devices, current string, dac, fft bool, cache *map[string]int) int {
	if current == "out" {
		if dac && fft {
			return 1
		} else {
			return 0
		}
	}

	key := paramsToKey(current, dac, fft)
	c, ok := (*cache)[key]
	if ok {
		return c
	}

	outputs, ok := devices[current]
	if !ok {
		panic("Device doesn't exist on list of devices")
	}
	result := 0
	for _, output := range outputs {
		v := 0
		if output == "dac" {
			v = countPathsWithDacAndFft(devices, output, true, fft, cache)
		} else if output == "fft" {
			v = countPathsWithDacAndFft(devices, output, dac, true, cache)
		} else {
			v = countPathsWithDacAndFft(devices, output, dac, fft, cache)
		}
		result += v
	}
	(*cache)[key] = result

	return result
}

func part2(lines []string) string {
	solution := 0

	devices := Devices{}

	for _, line := range lines {
		parts := strings.Split(line, " ")
		device := strings.Replace(parts[0], ":", "", 1)
		outputs := parts[1:]
		devices[device] = outputs
	}

	cache := map[string]int{}
	solution = countPathsWithDacAndFft(devices, "svr", false, false, &cache)

	return fmt.Sprintf("%d", solution)
}
