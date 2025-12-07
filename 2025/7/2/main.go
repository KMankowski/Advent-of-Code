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

	timelines := run(r)

	l.Info("success", slog.Int("part2", timelines))
}

func run(r io.Reader) int {
	particle, splitters, rows := parseManifold(r)

	cache := make(map[[2]int]int)

	return getTimelines(particle, splitters, rows, cache)
}

func getTimelines(particle [2]int, splitters map[[2]int]struct{}, rows int, cache map[[2]int]int) int {
	for {
		if particle[0] == rows {
			return 1
		}

		if timelines, ok := cache[particle]; ok {
			return timelines
		}

		nextCoordinate := [2]int{particle[0] + 1, particle[1]}
		_, isSplitterCollision := splitters[nextCoordinate]
		if isSplitterCollision {
			leftParticle := [2]int{particle[0], particle[1] - 1}
			rightParticle := [2]int{particle[0], particle[1] + 1}

			timelines := 0
			timelines += getTimelines(leftParticle, splitters, rows, cache)
			timelines += getTimelines(rightParticle, splitters, rows, cache)
			cache[particle] = timelines
			return timelines
		}

		particle = nextCoordinate
	}
}

func parseManifold(r io.Reader) ([2]int, map[[2]int]struct{}, int) {
	var startingBeam [2]int
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
			startingBeam = [2]int{0, col}
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
