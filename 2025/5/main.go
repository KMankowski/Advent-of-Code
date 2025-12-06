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

func main() {
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))

	input, err := os.Open(INPUT_FILE)
	if err != nil {
		l.Error("os.Open() failed", "error", err)
		os.Exit(1)
	}

	freshIngredients, err := run(input)
	if err != nil {
		l.Error("run() failed", "error", err)
		os.Exit(1)
	}

	l.Info("success", slog.Int("ingredients", freshIngredients))
}

func run(r io.Reader) (int, error) {
	ranges, ingredients, err := parseRangesAndIngredients(r)
	if err != nil {
		return 0, fmt.Errorf("parseRangesAndIngredients() failed: %w", err)
	}

	freshIngredients := make([]int, 0)
	for _, ingredient := range ingredients {
		if isIngredientFresh(ranges, ingredient) {
			freshIngredients = append(freshIngredients, ingredient)
		}
	}

	return len(freshIngredients), nil
}

func isIngredientFresh(ranges [][2]int, ingredient int) bool {
	for _, r := range ranges {
		if ingredient >= r[0] && ingredient <= r[1] {
			return true
		}
	}
	return false
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
