package main

import (
	"bufio"
	"io"
	"log/slog"
	"os"
	"strconv"
)

const STARTING_POSITION = 50
const MATCHING_POSITION = 0
const INPUT_FILE = "./2025/1/input.txt"

func main() {
	// Default log level is Info, ignoring Debug statements
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	input, err := os.Open(INPUT_FILE)
	if err != nil {
		logger.Error("os.Open() failed", "error", err)
		os.Exit(1)
	}

	matches, passes, err := run(input, logger)
	if err != nil {
		logger.Error("run() failed", "error", err)
		os.Exit(1)
	}

	logger.Info("success", slog.Int("matches", matches), slog.Int("passes", passes))
}

func run(input io.Reader, logger *slog.Logger) (int, int, error) {
	rotations, err := parseRotations(input)
	if err != nil {
		return 0, 0, err
	}

	matches := countMatches(rotations, logger)

	// No need to log position twice
	passes := countPasses(rotations)

	return matches, passes, nil
}

func countPasses(rotations []int) int {
	passes := 0
	currPosition := STARTING_POSITION
	for _, rotation := range rotations {
		if (rotation < 0) && (currPosition+rotation < 0) {
			passes += ((((currPosition + rotation) / 100) - 1) * -1)
			if currPosition == 0 {
				passes--
			}
			if currPosition+rotation%100 == 0 {
				passes--
			}
		} else if (rotation > 0) && (currPosition+rotation > 99) {
			passes += ((currPosition + rotation) / 100)
		}

		currPosition += rotation
		currPosition %= 100
		if currPosition < 0 {
			currPosition += 100
		}
		if currPosition == 0 && rotation > 0 {
			passes--
		}
	}
	return passes
}

func countMatches(rotations []int, logger *slog.Logger) int {
	matches := 0
	currPosition := STARTING_POSITION
	for _, rotation := range rotations {
		currPosition += rotation
		currPosition %= 100
		if currPosition < 0 {
			currPosition += 100
		}
		logger.Debug("updated position", slog.Int("position", currPosition))
		if currPosition == MATCHING_POSITION {
			matches++
		}
	}
	return matches
}

// R50 L1000 -> [50, -1000]
func parseRotations(input io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	parsedRotations := make([]int, 0)

	for scanner.Scan() {
		rawRotation := scanner.Text()

		parsedRotation, err := strconv.Atoi(rawRotation[1:])
		if err != nil {
			return nil, err
		}

		if rawRotation[0] == 'L' {
			parsedRotation *= -1
		}

		parsedRotations = append(parsedRotations, parsedRotation)
	}

	return parsedRotations, nil
}
