package main

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"os"
)

const INPUT_FILE = "./2025/4/input.txt"

type gridEntry int

const (
	roll gridEntry = iota
	empty
	cleared
)

func main() {
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))

	input, err := os.Open(INPUT_FILE)
	if err != nil {
		l.Error("os.Open() failed", "error", err)
		os.Exit(1)
	}

	rollsPart1, rollsPart2, err := run(input)
	if err != nil {
		l.Error("run() failed", "error", err)
		os.Exit(1)
	}

	l.Info("success", slog.Int("part1", rollsPart1), slog.Int("part2", rollsPart2))
}

func run(r io.Reader) (int, int, error) {
	// Part 1
	grid, err := parseGrid(r)
	if err != nil {
		return 0, 0, fmt.Errorf("parseGrid() error: %w", err)
	}
	rolls := len(getReachableRolls(grid))

	// Part 2
	rollsPart2 := iterateReachableRolls(grid)

	return rolls, rollsPart2, nil
}

func iterateReachableRolls(grid [][]gridEntry) int {
	rollsCleared := 0
	reachableRolls := getReachableRolls(grid)

	for len(reachableRolls) > 0 {
		clearReachableRolls(grid, reachableRolls)
		rollsCleared += len(reachableRolls)

		reachableRolls = getReachableRolls(grid)
	}
	return rollsCleared
}

func clearReachableRolls(grid [][]gridEntry, reachableRolls [][2]int) {
	for _, reachableRoll := range reachableRolls {
		grid[reachableRoll[0]][reachableRoll[1]] = cleared
	}
}

func getReachableRolls(grid [][]gridEntry) [][2]int {
	reachableRolls := make([][2]int, 0)
	for rowIndex := range grid {
		for colIndex := range grid[rowIndex] {
			if isReachableRoll(grid, rowIndex, colIndex) {
				reachableRolls = append(reachableRolls, [2]int{rowIndex, colIndex})
			}
		}
	}
	return reachableRolls
}

func isReachableRoll(grid [][]gridEntry, row, col int) bool {
	if !isRoll(grid, row, col) {
		return false
	}

	neighboringRolls := 0

	if isRoll(grid, row, col+1) {
		neighboringRolls++
	}
	if isRoll(grid, row, col-1) {
		neighboringRolls++
	}
	if isRoll(grid, row+1, col) {
		neighboringRolls++
	}
	if isRoll(grid, row+1, col+1) {
		neighboringRolls++
	}
	if isRoll(grid, row+1, col-1) {
		neighboringRolls++
	}
	if isRoll(grid, row-1, col) {
		neighboringRolls++
	}
	if isRoll(grid, row-1, col+1) {
		neighboringRolls++
	}
	if isRoll(grid, row-1, col-1) {
		neighboringRolls++
	}

	return neighboringRolls < 4
}

func isRoll(grid [][]gridEntry, row, col int) bool {
	if row < 0 || row >= len(grid) {
		return false
	}
	if col < 0 || col >= len(grid[row]) {
		return false
	}
	return grid[row][col] == roll
}

func parseGrid(r io.Reader) ([][]gridEntry, error) {
	parsedGrid := make([][]gridEntry, 0)

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		rawRow := s.Text()
		parsedRow := make([]gridEntry, 0)

		for _, char := range rawRow {
			switch char {
			case '.':
				parsedRow = append(parsedRow, empty)
			case '@':
				parsedRow = append(parsedRow, roll)
			default:
				return nil, fmt.Errorf("unrecognized input char %v in row %v", char, rawRow)
			}
		}

		parsedGrid = append(parsedGrid, parsedRow)
	}

	return parsedGrid, nil
}
