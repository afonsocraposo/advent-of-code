package main

import (
	"log"
	"os"
	"strconv"
	"time"

	year2024 "github.com/afonsocraposo/advent-of-code/internal/2024"
)

var yearDays = map[int]map[int]func(){
	2024: year2024.Days,
}

func main() {
	log.SetFlags(0)

	args := os.Args[1:]

	var year, day int
	var err error

	if len(args) == 0 {
		year = time.Now().Year()
		day = time.Now().Day()
		log.Printf("Using current date: Year %d, Day %02d\n", year, day)
	} else if len(args) == 1 {
		year = time.Now().Year()
		day, err = strconv.Atoi(args[0])
		if err != nil {
			log.Fatalln("Invalid day:", err)
		}
		log.Printf("Using current year %d with specified day %02d\n", year, day)
	} else {
		year, err = strconv.Atoi(args[0])
		if err != nil {
			log.Fatalln("Invalid year:", err)
		}
		day, err = strconv.Atoi(args[1])
		if err != nil {
			log.Fatalln("Invalid day:", err)
		}
	}

	log.Printf("Advent of Code %d - Day %02d\n", year, day)

	yearMap, yearExists := yearDays[year]
	if !yearExists {
		log.Fatalf("Year %d is not implemented yet\n", year)
	}

	dayFunc, dayExists := yearMap[day]
	if !dayExists {
		log.Fatalf("Day %02d for year %d is not implemented yet\n", day, year)
	}

	dayFunc()
}
