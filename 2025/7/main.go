package main

import (
	"bufio"
	"io"
	"log/slog"
	"os"
)

const INPUT_FILE = "./2025/7/input.txt"

func main() {
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))

	r, err := os.Open(INPUT_FILE)
	if err != nil {
		l.Error("os.Open() failed", "error", err)
		os.Exit(1)
	}

	splits := run(r)

	l.Info("success", slog.Int("part1", splits))
}

func run(r io.Reader) int {
	startingBeam, splitters, rows := parseManifold(r)

	currBeams := make(map[int]struct{})
	currBeams[startingBeam] = struct{}{}

	splits := 0
	for row := 0; row < rows; row++ {
		newBeams := make(map[int]struct{})
		for currBeam := range currBeams {
			if _, ok := splitters[[2]int{row + 1, currBeam}]; !ok {
				newBeams[currBeam] = struct{}{}
				continue
			}
			newBeams[currBeam-1] = struct{}{}
			newBeams[currBeam+1] = struct{}{}
			splits++
		}
		currBeams = newBeams
	}

	return splits
}

func parseManifold(r io.Reader) (int, map[[2]int]struct{}, int) {
	startingBeam := 0
	splitters := make(map[[2]int]struct{})
	rows := 0

	grid := make([][]rune, 0)
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		grid = append(grid, []rune(s.Text()))
	}

	for col := 0; col < len(grid[0]); col++ {
		if grid[0][col] == 'S' {
			startingBeam = col
		}
	}

	rows = len(grid)
	for row := 1; row < rows; row++ {
		for col := 0; col < len(grid[row]); col++ {
			if grid[row][col] == '^' {
				splitters[[2]int{row, col}] = struct{}{}
			}
		}
	}

	return startingBeam, splitters, rows
}
