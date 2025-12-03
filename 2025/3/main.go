package main

import (
	"bufio"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strconv"
)

const INPUT_FILE = "./2025/3/input.txt"

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	input, err := os.Open(INPUT_FILE)
	if err != nil {
		logger.Error("os.Open() failed", "error", err)
		os.Exit(1)
	}

	joltage, err := run(input, logger)
	if err != nil {
		logger.Error("run() failed", "error", err)
		os.Exit(1)
	}

	logger.Info("success", slog.Int("joltage", joltage))
}

func run(r io.Reader, l *slog.Logger) (int, error) {
	banks := parseBanks(r)
	joltage := 0
	for _, bank := range banks {
		currJoltage, err := getJoltage(bank)
		if err != nil {
			return 0, fmt.Errorf("getJoltage() error: %w", err)
		}

		l.Debug("bank processed", slog.Int("joltage", currJoltage))
		joltage += currJoltage
	}
	return joltage, nil
}

func getJoltage(bank string) (int, error) {
	max := int(bank[0] - byte('0'))
	maxIndex := 0
	for i, runeCurr := range bank[:len(bank)-1] {
		intCurr := int(runeCurr - rune('0'))
		if intCurr > max {
			max = intCurr
			maxIndex = i
		}
	}

	secondMax := int(bank[maxIndex+1] - byte('0'))
	for _, runeCurr := range bank[maxIndex+1:] {
		intCurr := int(runeCurr - rune('0'))
		if intCurr > secondMax {
			secondMax = intCurr
		}
	}

	joltage, err := strconv.Atoi(strconv.Itoa(max) + strconv.Itoa(secondMax))
	if err != nil {
		return 0, fmt.Errorf("strconv.Atoi() error: %w", err)
	}

	return joltage, nil
}

func parseBanks(r io.Reader) []string {
	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)

	banks := make([]string, 0)
	for s.Scan() {
		banks = append(banks, s.Text())
	}

	return banks
}
