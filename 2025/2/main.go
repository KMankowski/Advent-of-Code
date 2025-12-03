package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

const INPUT_FILE = "./2025/2/input.txt"

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	input, err := os.Open(INPUT_FILE)
	if err != nil {
		logger.Error("os.Open() failed", "error", err)
		os.Exit(1)
	}

	sum, err := run(input, logger)
	if err != nil {
		logger.Error("run() failed", "error", err)
		os.Exit(1)
	}

	logger.Info("success", slog.Int("sum", sum))
}

func run(reader io.Reader, logger *slog.Logger) (int, error) {
	ranges, err := parseRanges(reader)
	if err != nil {
		return 0, fmt.Errorf("parseRanges() failed: %w", err)
	}

	sum := 0
	for _, r := range ranges {
		invalidIDs := getInvalidIDs(r)
		logger.Info("Invalid IDs", "range", r, "ids", invalidIDs)

		for _, id := range invalidIDs {
			sum += id
		}
	}

	return sum, nil
}

func getInvalidIDs(r [2]int) []int {
	min := r[0]
	max := r[1]

	invalidIDs := make([]int, 0)
	for currID := min; currID <= max; currID++ {
		if isPart1Invalid(currID) {
			invalidIDs = append(invalidIDs, currID)
		}
	}
	return invalidIDs
}

func isPart1Invalid(id int) bool {
	stringID := strconv.Itoa(id)
	if len(stringID)%2 != 0 {
		return false
	}
	secondHalfStart := len(stringID) / 2
	for firstHalfCurr := 0; firstHalfCurr < secondHalfStart; firstHalfCurr++ {
		if stringID[firstHalfCurr] != stringID[firstHalfCurr+secondHalfStart] {
			return false
		}
	}
	return true
}

func parseRanges(r io.Reader) ([][2]int, error) {
	csvReader := csv.NewReader(r)

	// doesn't matter how many comma-separated values in input
	csvReader.FieldsPerRecord = -1

	rawRanges, err := csvReader.Read()
	if err != nil {
		return nil, fmt.Errorf("csvReader.Read() failed: %w", err)
	}

	parsedRanges := make([][2]int, 0)
	for _, rawRange := range rawRanges {
		rawRangeValues := strings.Split(rawRange, "-")
		if len(rawRangeValues) != 2 {
			return nil, fmt.Errorf("strings.Split() returned a len() of %v on the []string %v", len(rawRangeValues), rawRangeValues)
		}

		var parsedRange [2]int
		parsedRange[0], err = strconv.Atoi(rawRangeValues[0])
		if err != nil {
			return nil, fmt.Errorf("strconv.Atoi() failed on rawRangeValues[0] %v: %w", rawRangeValues, err)
		}
		parsedRange[1], err = strconv.Atoi(rawRangeValues[1])
		if err != nil {
			return nil, fmt.Errorf("strconv.Atoi() failed on rawRangeValues[1] %v: %w", rawRangeValues, err)
		}

		parsedRanges = append(parsedRanges, parsedRange)
	}

	return parsedRanges, nil
}
