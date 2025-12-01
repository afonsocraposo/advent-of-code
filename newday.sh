#!/bin/sh

if [ -z "$1" ]; then
  YEAR=$(date +"%Y")
else
  YEAR="$1"
fi

if [ -z "$2" ]; then
  DAY=$(date +"%d")
else
  DAY=$(printf "%02d" "$2")
fi

mkdir -p "internal/${YEAR}/day${DAY}" "puzzles/${YEAR}/day${DAY}"

ALREADY_EXISTS=$(ls internal/${YEAR}/day${DAY} 2>/dev/null)

touch "internal/${YEAR}/day${DAY}/day${DAY}.go" "puzzles/${YEAR}/day${DAY}/example1.txt" "puzzles/${YEAR}/day${DAY}/example1.solution1.txt" "puzzles/${YEAR}/day${DAY}/input1.txt"

if [ -z "$ALREADY_EXISTS" ]; then
NEW_DAY="internal/${YEAR}/day${DAY}/day${DAY}.go"
cp internal/day00/day00.go $NEW_DAY
gofmt -w "internal/${YEAR}/day${DAY}/day${DAY}.go"
fi

echo "Setup completed for Year $YEAR, Day $DAY"
