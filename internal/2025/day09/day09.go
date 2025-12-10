package day09

import (
	"fmt"
	"log"
	"slices"

	"github.com/afonsocraposo/advent-of-code/internal/utils/matrix"
	"github.com/afonsocraposo/advent-of-code/internal/utils/numbers"
	"github.com/afonsocraposo/advent-of-code/internal/utils/point"
	"github.com/afonsocraposo/advent-of-code/internal/utils/runner"
)

const year = 2025
const day = 9

func Main() {
	r := runner.New(year, day, part1, part2)
	r.TestPart1(1, 1)
	r.RunPart1(1)

	r.TestPart2(1, 2)
	r.RunPart2(1)
}

type PointsArea struct {
	Point1 point.Point
	Point2 point.Point
	Area   int
}

func part1(lines []string) string {
	solution := 0

	puzzle, err := matrix.ParseMatrix(lines, ",")
	if err != nil {
		log.Fatalln(err)
	}

	m, _ := puzzle.Size()
	tiles := make([]point.Point, m)
	for i, row := range puzzle.Rows {
		tiles[i] = point.NewPoint(row.Values[0], row.Values[1])
	}

	areas := []PointsArea{}
	for i, point1 := range tiles[:len(tiles)-1] {
		for _, point2 := range tiles[i+1:] {
			distance := point1.Distance(point2)
			area := (numbers.IntAbs(distance.I) + 1) * (numbers.IntAbs(distance.J) + 1)
			areas = append(areas, PointsArea{point1, point2, area})
		}
	}
	slices.SortFunc(areas, func(point1, point2 PointsArea) int {
		if point1.Area > point2.Area {
			return 1
		} else if point1.Area == point2.Area {
			return 0
		} else {
			return -1
		}
	})
	slices.Reverse(areas)
	solution = areas[0].Area

	return fmt.Sprintf("%d", solution)
}

func areaInRegion(mat matrix.Matrix, point1 point.Point, point2 point.Point) bool {
	minI := numbers.IntMin(point1.I, point2.I)
	minJ := numbers.IntMin(point1.J, point2.J)
	maxI := numbers.IntMax(point1.I, point2.I)
	maxJ := numbers.IntMax(point1.J, point2.J)
	for i := minI; i <= maxI; i++ {
		v1, err := mat.Get(i, minJ)
		if err != nil {
			panic(err)
		}
		if v1 == 0 {
			return false
		}

		v2, err := mat.Get(i, maxJ)
		if err != nil {
			panic(err)
		}
		if v2 == 0 {
			return false
		}
	}
	for j := minJ; j <= maxJ; j++ {
		v1, err := mat.Get(minI, j)
		if err != nil {
			panic(err)
		}
		if v1 == 0 {
			return false
		}

		v2, err := mat.Get(maxI, j)
		if err != nil {
			panic(err)
		}
		if v2 == 0 {
			return false
		}
	}
	return true
}

func containsIntersection(perimeter []point.Point, point1 point.Point, point2 point.Point) bool {
	minI := numbers.IntMin(point1.I, point2.I)
	minJ := numbers.IntMin(point1.J, point2.J)
	maxI := numbers.IntMax(point1.I, point2.I)
	maxJ := numbers.IntMax(point1.J, point2.J)

	for _, p := range perimeter {
		// exclude border
		if p.InsideBounds(minI+1, minJ+1, maxI-1, maxJ-1) {
			return true
		}
	}

	return false
}

func part2(lines []string) string {
	solution := 0

	puzzle, err := matrix.ParseMatrix(lines, ",")
	if err != nil {
		log.Fatalln(err)
	}

	m, _ := puzzle.Size()
	tiles := make([]point.Point, m)
	for i, row := range puzzle.Rows {
		tiles[i] = point.NewPoint(row.Values[1], row.Values[0])
	}

	perimeter := []point.Point{}
	tilesWithFirst := append(tiles, tiles[0])
	for t := range tiles {
		tile1 := tilesWithFirst[t]
		tile2 := tilesWithFirst[t+1]
		perimeter = append(perimeter, point.NewPoint(tile1.I, tile1.J))
		distance := tile1.Distance(tile2)
		dirI := 0
		dirJ := 0
		if distance.I != 0 {
			dirI = distance.I / numbers.IntAbs(distance.I)
		}
		if distance.J != 0 {
			dirJ = distance.J / numbers.IntAbs(distance.J)
		}
		tile1.Sum(point.NewPoint(dirI, dirJ))
		for tile1.I != tile2.I || tile1.J != tile2.J {
			perimeter = append(perimeter, point.NewPoint(tile1.I, tile1.J))
			tile1.Sum(point.NewPoint(dirI, dirJ))
		}
	}

	for i, point1 := range tiles[:len(tiles)-1] {
		for _, point2 := range tiles[i+1:] {
			distance := point1.Distance(point2)
			area := (numbers.IntAbs(distance.I) + 1) * (numbers.IntAbs(distance.J) + 1)
			if area < solution || containsIntersection(perimeter, point1, point2) {
				continue
			}
			if area > solution {
				solution = area
			}
		}
	}

	return fmt.Sprintf("%d", solution)
}
