package day08

import (
	"fmt"
	"log"
	"slices"
	"strconv"

	"github.com/afonsocraposo/advent-of-code/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code/internal/utils/point"
	"github.com/afonsocraposo/advent-of-code/internal/utils/runner"
	"github.com/afonsocraposo/advent-of-code/internal/utils/set"
)

const year = 2025
const day = 8

func Main() {
	r := runner.New(year, day, part1, part2)
	r.TestPart1(1, 1)
	r.RunPart1(1)

	r.TestPart2(1, 2)
	r.RunPart2(1)
}

type DistancePair struct {
	Point1   point.Point3D
	Point2   point.Point3D
	Distance float64
}

func compareDistancePair(a, b DistancePair) int {
	if a.Distance > b.Distance {
		return 1
	} else if a.Distance < b.Distance {
		return -1
	} else {
		return 0
	}
}

func pointToKey(p point.Point3D) string {
	return fmt.Sprintf("%d,%d,%d", p.X, p.Y, p.Z)
}

func part1(lines []string) string {
	solution := 0
	limit := 10
	if len(lines) > 20 {
		limit = 1000
	}

	n := len(lines)
	boxes := make([]point.Point3D, n)
	for i, line := range lines {
		v, err := matrix.ParseVector(line, ",")
		if err != nil {
			log.Fatalln(err)
		}

		boxes[i] = point.NewPoint3D(v.Get(0), v.Get(1), v.Get(2))
	}

	distances := []DistancePair{}
	for i := range n {
		for j := range i {
			if i == j {
				continue
			}
			a := boxes[i]
			b := boxes[j]
			d := a.Distance(b)
			distances = append(distances, DistancePair{a, b, d})
		}
	}

	slices.SortFunc(distances, compareDistancePair)

	connections := map[string]int{}
	circuitNr := 1
	c := 0
	for len(distances) > 0 {
		pair := distances[0]
		distances = distances[1:]
		key1 := pointToKey(pair.Point1)
		key2 := pointToKey(pair.Point2)
		c++
		if c > limit {
			break
		}

		circuit1, connected1 := connections[key1]
		circuit2, connected2 := connections[key2]
		if connected1 && connected2 {
			if circuit1 == circuit2 {
				continue
			} else {
				for k, v := range connections {
					if v == circuit2 {
						connections[k] = circuit1
					}
				}
				continue
			}
		}
		if connected1 && !connected2 {
			connections[key2] = circuit1
		} else if !connected1 && connected2 {
			connections[key1] = circuit2
		} else {
			connections[key1] = circuitNr
			connections[key2] = circuitNr
			circuitNr++
		}
	}
	boxesPerCircuit := map[int]int{}
	for _, circuit := range connections {
		if _, exists := boxesPerCircuit[circuit]; !exists {
			boxesPerCircuit[circuit] = 0
		}
		boxesPerCircuit[circuit] = boxesPerCircuit[circuit] + 1
	}
	s := set.NewSet()
	for _, boxes := range boxesPerCircuit {
		s.Add(fmt.Sprintf("%d", boxes))
	}
	solution = 1
	boxesCount := make([]int, s.Size())
	for i, v := range s.Values() {
		value, _ := strconv.Atoi(v)
		boxesCount[i] = value
	}
	slices.Sort(boxesCount)
	slices.Reverse(boxesCount)
	for _, v := range boxesCount[:3] {
		solution *= v
	}

	return fmt.Sprintf("%d", solution)
}

func part2(lines []string) string {
	solution := 0

	n := len(lines)
	boxes := make([]point.Point3D, n)
	for i, line := range lines {
		v, err := matrix.ParseVector(line, ",")
		if err != nil {
			log.Fatalln(err)
		}

		boxes[i] = point.NewPoint3D(v.Get(0), v.Get(1), v.Get(2))
	}

	distances := []DistancePair{}
	for i := range n {
		for j := range i {
			if i == j {
				continue
			}
			a := boxes[i]
			b := boxes[j]
			d := a.Distance(b)
			distances = append(distances, DistancePair{a, b, d})
		}
	}

	slices.SortFunc(distances, compareDistancePair)

	connections := map[string]int{}
	circuits := set.NewSet()
	connected := set.NewSet()
	circuitNr := 1
	c := 0
	for len(distances) > 0 {
		pair := distances[0]
		distances = distances[1:]
		key1 := pointToKey(pair.Point1)
		key2 := pointToKey(pair.Point2)
		c++
		connected.Add(key1)
		connected.Add(key2)

		circuit1, connected1 := connections[key1]
		circuit2, connected2 := connections[key2]
		if connected1 && connected2 {
			if circuit1 != circuit2 {
				for k, v := range connections {
					if v == circuit2 {
						connections[k] = circuit1
					}
				}
				circuits.Remove(fmt.Sprintf("%d", circuit2))
			}
		} else {
			if connected1 && !connected2 {
				connections[key2] = circuit1
			} else if !connected1 && connected2 {
				connections[key1] = circuit2
			} else {
				connections[key1] = circuitNr
				connections[key2] = circuitNr
				circuits.Add(fmt.Sprintf("%d", circuitNr))
				circuitNr++
			}
		}
		if connected.Size() == len(boxes) && circuits.Size() == 1 {
			solution = pair.Point1.X * pair.Point2.X
			break
		}
	}

	return fmt.Sprintf("%d", solution)
}
