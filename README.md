# Advent of Code 🎄✨

```
     Merry Christmas!
 ------------------------
   \
    \
     \   ,_---~~~~~----._
  _,,_,*^____      _____``*g*\"*,
 / __/ /'     ^.  /      \ ^@q   f
[  @f | @))    |  | @))   l  0 _/
 \`/   \~____ / __ \_____/    \
  |           _l__l_           I
  }          [______]           I
  ]            | | |            |
  ]             ~ ~             |
  |                            |
   |                           |
```

Welcome to **Advent of Code**! 🎅✨ This repo features my solutions to the daily programming puzzles, written in **Golang**, with a clean architecture for festive coding!

---

## Why Go? 🐹

- **Coolness**: I like Go and don't get the chance to use it on my day-to-day
- **Speed**: No waiting around for Santa's sleigh.
- **Simplicity**: Go keeps the solutions straightforward and delightful.
- **Concurrency**: Perfect for parallelizing Santa's busy workshop tasks.

---

## Repo Structure 📂

Here's how the magic is organized:

```
.
├── LICENSE               # Licensing information
├── mise.toml             # Go version management (1.25.4)
├── newday.sh             # Script to scaffold new day solutions
├── cmd
│   └── main.go           # Entry point for running solutions
├── go.mod                # Go module configuration
├── internal
│   ├── 2024              # Solutions organized by year
│   │   ├── all.go        # Aggregates all 2024 day solutions
│   │   ├── day01
│   │   │   └── day01.go  # Solution logic for Day 1
│   │   ├── day02
│   │   │   └── day02.go  # Solution logic for Day 2
│   │   └── ...           # Days 3-25
│   └── utils             # Shared utilities
│       ├── algorithms    # Common algorithms (Dijkstra, BFS, etc.)
│       ├── filereader    # File parsing utilities
│       ├── matrix        # Matrix and vector manipulation
│       ├── point         # 2D/3D point helpers
│       ├── set           # Set data structure
│       └── ...           # Other utilities
└── puzzles               # Puzzle input files (ignored in Git)
    └── 2024
        ├── day01
        │   ├── example.txt
        │   └── input.txt
        └── ...
```

---

## Dependencies 📦

This project uses the following Go packages:

- [`github.com/draffensperger/golp`](https://github.com/draffensperger/golp) - Linear programming solver

## Running a Solution 🏃

To solve the puzzles for a specific day, run the following command:

```bash
go run cmd/main.go <year> <day>
```

You can also omit arguments to use the current date, or provide just the day to use the current year.

### Examples

```bash
# Run Day 2 of 2024
go run cmd/main.go 2024 2

# Run Day 5 of current year
go run cmd/main.go 5

# Run today's puzzle
go run cmd/main.go
```

### Building and Running

```bash
# Build the binary
mise exec -- go build -o aoc ./cmd

# Run the binary
./aoc 2024 2
```

### Output

```plaintext
Advent of Code 2024 - Day 02
Part 1:
The solution is: 218
Part 2:
The solution is: 290
```

---

## Adding New Days 🗓️

Use the provided script to scaffold a new day:

```bash
./newday.sh <year> <day>
```

This will:

1. Create the day directory: `internal/<year>/day<day>/`
2. Generate a template `day<day>.go` file with the standard structure
3. Create puzzle input directories: `puzzles/<year>/day<day>/`

### Manual Setup

If you prefer to add a day manually:

1. Create a new directory under `internal/<year>` for the day:

   ```bash
   mkdir -p internal/2024/day03
   ```

2. Write your solution in `internal/2024/day03/day03.go`:

   ```go
   package day03

   import "github.com/afonsocraposo/advent-of-code/internal/utils/filereader"

   const year = 2024

   func Main() {
       lines := filereader.ReadLines(year, 3)
       // Your solution here
   }
   ```

3. Register the day in `internal/2024/all.go`:

   ```go
   import day03 "github.com/afonsocraposo/advent-of-code/internal/2024/day03"

   var Days = map[int]func(){
       // ...
       3: day03.Main,
       // ...
   }
   ```

---

## Go Version Management 🐹

This project uses [mise](https://mise.jdx.dev/) to manage the Go version (1.25.4). The version is specified in `mise.toml`.

To install mise and the correct Go version:

```bash
# Install mise (if not already installed)
curl https://mise.run | sh

# Install the Go version specified in mise.toml
mise install

# Run commands with the correct Go version
mise exec -- go build -o aoc ./cmd
```

---

## Architecture Notes 🏗️

### Year-Based Organization

Solutions are organized by year in `internal/<year>/`. Each year has an `all.go` file that aggregates all day solutions into a map, making it easy to add new years without cluttering `main.go`.

### Utilities

Common utilities are shared across all years in `internal/utils/`, including:

- **algorithms**: Dijkstra, BFS, flood fill, etc.
- **filereader**: Read puzzle inputs
- **matrix**: 2D grid operations
- **point**: 2D/3D coordinate helpers
- **set**: Set data structure
- And more!

---
