package day010

import (
	"fmt"
	"log"
	"strconv"

	"github.com/afonsocraposo/advent-of-code/internal/utils/filereader"
)

const day = 01

var examples = []int{1}

func Main() {
	log.Printf("DAY %d\n", day)

	for part := 1; part <= 2; part++ {
		var partMethod func([]string) string
		if part == 1 {
			partMethod = part1
		} else {
			partMethod = part2
		}

		log.Printf("Part %d:\n", part)
		for _, example := range examples {
			exampleLines, err := filereader.ReadDayExample(2025, day, example)
			if err != nil {
				log.Fatalln(err)
			}
			expectedSolution, err := filereader.ReadDayExampleSolution(2025, day, example, part)
			if err != nil {
				continue
			}

			exampleSolution := partMethod(exampleLines)
			if exampleSolution != expectedSolution {
				log.Fatalf("WRONG solution for example %d. Expected: %s, Actual: %s\n", example, expectedSolution, exampleSolution)
			} else {
				log.Printf("The solution is CORRECT for example %d. Expected/actual: %s\n", example, exampleSolution)
			}

		}
		inputLines, err := filereader.ReadDayInput(2025, day, 1)
		if err != nil {
			log.Fatalln(err)
		}
		inputSolution := partMethod(inputLines)
		log.Printf("The solution for the input is: %s\n", inputSolution)
	}
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
