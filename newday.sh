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
cp internal/2024/day00/day00.go $NEW_DAY
sed -i "" "s/package day00/package day${DAY}/g" $NEW_DAY
sed -i "" "s/DAY 0/DAY ${DAY#0}/g" $NEW_DAY
sed -i "" "s/filereader\.NewFromDayInput(2024, 0,/filereader.NewFromDayInput(${YEAR}, ${DAY#0},/g" $NEW_DAY
sed -i "" "s/filereader\.NewFromDayExample(2024, 0,/filereader.NewFromDayExample(${YEAR}, ${DAY#0},/g" $NEW_DAY
sed -i "" "s/filereader\.ReadDayInput(2024, 0,/filereader.ReadDayInput(${YEAR}, ${DAY#0},/g" $NEW_DAY
sed -i "" "s/filereader\.ReadDayExample(2024, 0,/filereader.ReadDayExample(${YEAR}, ${DAY#0},/g" $NEW_DAY
sed -i "" "s/filereader\.ReadDayExampleSolution(2024, 0,/filereader.ReadDayExampleSolution(${YEAR}, ${DAY#0},/g" $NEW_DAY

gofmt -w "internal/${YEAR}/day${DAY}/day${DAY}.go"
fi

MAIN_FILE="cmd/main.go"
IMPORT_STATEMENT="day${DAY} \\\"github.com/afonsocraposo/advent-of-code/internal/${YEAR}/day${DAY}\\\""
DAYS_ENTRY="${DAY#0}: day${DAY}.Main,"

if ! grep -q "day${DAY} \"github.com/afonsocraposo/advent-of-code/internal/${YEAR}/day${DAY}\"" "$MAIN_FILE"; then
  sed -i '' "/import (/a\\
	$IMPORT_STATEMENT
" "$MAIN_FILE"
fi

if ! grep -q "${DAY#0}: day${DAY}.Main," "$MAIN_FILE"; then
  sed -i '' "/var days${YEAR} = map\\[int\\]func()/,/^}/ {
    /^}/ i\\
	${DAYS_ENTRY}
  }" "$MAIN_FILE"
fi

gofmt -w "$MAIN_FILE"

echo "Setup completed for Year $YEAR, Day $DAY"
