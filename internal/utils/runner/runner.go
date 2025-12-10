package runner

import (
	"log"
	"time"

	"github.com/afonsocraposo/advent-of-code/internal/utils/filereader"
)

type Runner struct {
	year  int
	day   int
	part1 partFunction
	part2 partFunction
}

type partFunction func([]string) string

func New(year int, day int, part1 partFunction, part2 partFunction) Runner {
	return Runner{year, day, part1, part2}
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("⏱️ %s took %s\n", name, elapsed)
}

func (r *Runner) TestPart1(example int, solution int) {
	defer timeTrack(time.Now(), "TestPart1")
	log.Printf("\n🧪 Running part 1 logic for example %d and solution %d\n", example, solution)
	exampleLines, err := filereader.ReadDayExample(r.year, r.day, example)
	if err != nil {
		log.Fatalln(err)
	}
	expectedSolution, err := filereader.ReadDayExampleSolution(r.year, r.day, example, solution)
	if err != nil {
		log.Fatalln(err)
	}

	exampleSolution := r.part1(exampleLines)
	if exampleSolution != expectedSolution {
		log.Fatalf("❌ WRONG solution for example %d. Expected: %s, Actual: %s\n", example, expectedSolution, exampleSolution)
	} else {
		log.Printf("✅ The solution is CORRECT for example %d. Expected/actual: %s\n", example, exampleSolution)
	}
}

func (r *Runner) TestPart2(example int, solution int) {
	defer timeTrack(time.Now(), "TestPart2")
	log.Printf("\n🧪 Running part 2 logic for example %d and solution %d\n", example, solution)
	exampleLines, err := filereader.ReadDayExample(r.year, r.day, example)
	if err != nil {
		log.Fatalln(err)
	}
	expectedSolution, err := filereader.ReadDayExampleSolution(r.year, r.day, example, solution)
	if err != nil {
		log.Fatalln(err)
	}

	exampleSolution := r.part2(exampleLines)
	if exampleSolution != expectedSolution {
		log.Fatalf("❌ WRONG solution for example %d. Expected: %s, Actual: %s\n", example, expectedSolution, exampleSolution)
	} else {
		log.Printf("✅ The solution is CORRECT for example %d. Expected/actual: %s\n", example, exampleSolution)
	}
}

func (r *Runner) RunPart1(input int) {
	defer timeTrack(time.Now(), "RunPart1")
	log.Printf("\nℹ️ Running part 1 logic for input %d", input)
	inputLines, err := filereader.ReadDayInput(r.year, r.day, input)
	if err != nil {
		log.Fatalln(err)
	}
	inputSolution := r.part1(inputLines)
	log.Printf("☝️ The solution for the input is: %s\n", inputSolution)
}

func (r *Runner) RunPart2(input int) {
	defer timeTrack(time.Now(), "RunPart2")
	log.Printf("\nℹ️ Running part 2 logic for input %d", input)
	inputLines, err := filereader.ReadDayInput(r.year, r.day, input)
	if err != nil {
		log.Fatalln(err)
	}
	inputSolution := r.part2(inputLines)
	log.Printf("✌️ The solution for the input is: %s\n", inputSolution)
}
