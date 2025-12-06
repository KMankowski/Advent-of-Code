package main

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

const INPUT_FILE = "./2025/5/input.txt"

const (
	low int = iota
	high
)

func main() {
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))

	input, err := os.Open(INPUT_FILE)
	if err != nil {
		l.Error("os.Open() failed", "error", err)
		os.Exit(1)
	}

	freshIngredients, part2, err := run(input)
	if err != nil {
		l.Error("run() failed", "error", err)
		os.Exit(1)
	}

	l.Info("success", slog.Int("part1", freshIngredients), slog.Int("part2", part2))
}

func run(r io.Reader) (int, int, error) {
	ranges, ingredients, err := parseRangesAndIngredients(r)
	if err != nil {
		return 0, 0, fmt.Errorf("parseRangesAndIngredients() failed: %w", err)
	}

	// Part1
	freshIngredients := make([]int, 0)
	for _, ingredient := range ingredients {
		if isIngredientFresh(ranges, ingredient) {
			freshIngredients = append(freshIngredients, ingredient)
		}
	}

	// Part2
	normalizedRanges := normalizeRanges(ranges)

	totalOfRanges := 0
	for _, r := range normalizedRanges {
		totalOfRanges += r[high] - r[low] + 1
	}

	return len(freshIngredients), totalOfRanges, nil
}

func normalizeRanges(ranges [][2]int) [][2]int {
	normalizedRanges := make([][2]int, 0)
	for _, r := range ranges {
		normalizedRanges = normalizeRange(normalizedRanges, r)
	}
	return normalizedRanges
}

func normalizeRange(normalizedRanges [][2]int, newRange [2]int) [][2]int {
	for currRangeIndex, currRange := range normalizedRanges {
		if isInRange(currRange, newRange[low]) && isInRange(currRange, newRange[high]) {
			return normalizedRanges
		}
		if isInRange(currRange, newRange[low]) {
			newRange[low] = currRange[low]
			normalizedRanges = removeRange(normalizedRanges, currRangeIndex)
			return normalizeRange(normalizedRanges, newRange)
		}
		if isInRange(currRange, newRange[high]) {
			newRange[high] = currRange[high]
			normalizedRanges = removeRange(normalizedRanges, currRangeIndex)
			return normalizeRange(normalizedRanges, newRange)
		}
		if newRange[low] < currRange[low] && newRange[high] > currRange[high] {
			normalizedRanges = removeRange(normalizedRanges, currRangeIndex)
			return normalizeRange(normalizedRanges, newRange)
		}
	}
	normalizedRanges = append(normalizedRanges, newRange)
	return normalizedRanges
}

func removeRange(ranges [][2]int, index int) [][2]int {
	if ranges == nil || index >= len(ranges) || index < 0 {
		return ranges
	}
	if len(ranges) == 1 {
		return nil
	}
	ranges[index] = ranges[len(ranges)-1]
	return ranges[:len(ranges)-1]
}

func isIngredientFresh(ranges [][2]int, ingredient int) bool {
	for _, r := range ranges {
		if isInRange(r, ingredient) {
			return true
		}
	}
	return false
}

func isInRange(r [2]int, v int) bool {
	return v >= r[low] && v <= r[high]
}

func parseRangesAndIngredients(r io.Reader) ([][2]int, []int, error) {
	ranges := make([][2]int, 0)
	ingredients := make([]int, 0)

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		line := s.Text()

		// Line sep between ranges and ingredients
		if len(line) == 0 {
			continue
		}

		if !strings.Contains(line, "-") {
			// Ingredient
			ingredient, err := strconv.Atoi(line)
			if err != nil {
				return nil, nil, fmt.Errorf("strconv.Atoi() failed on ingredient: %w", err)
			}
			ingredients = append(ingredients, ingredient)
		} else {
			// Range
			rawRange := strings.Split(line, "-")
			if len(rawRange) != 2 {
				return nil, nil, fmt.Errorf("strings.Split() returned invalid rawRange: %v", rawRange)
			}

			low, err := strconv.Atoi(rawRange[0])
			if err != nil {
				return nil, nil, fmt.Errorf("strconv.Atoi() failed on low: %w", err)
			}

			high, err := strconv.Atoi(rawRange[1])
			if err != nil {
				return nil, nil, fmt.Errorf("strconv.Atoi() failed on high: %w", err)
			}

			ranges = append(ranges, [2]int{low, high})
		}
	}
	return ranges, ingredients, nil
}
