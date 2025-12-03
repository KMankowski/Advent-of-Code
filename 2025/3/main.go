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

const INPUT_FILE = "./2025/3/input.txt"

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	input, err := os.Open(INPUT_FILE)
	if err != nil {
		logger.Error("os.Open() failed", "error", err)
		os.Exit(1)
	}

	joltagePart1, joltagePart2, err := run(input, logger)
	if err != nil {
		logger.Error("run() failed", "error", err)
		os.Exit(1)
	}

	logger.Info("success", slog.Int("part1", joltagePart1), slog.Int("part2", joltagePart2))
}

func run(r io.Reader, l *slog.Logger) (int, int, error) {
	banks := parseBanks(r)

	// Part 1
	joltagePart1 := 0
	for _, bank := range banks {
		currJoltage, err := getJoltagePart1(bank)
		if err != nil {
			return 0, 0, fmt.Errorf("getJoltagePart1() error: %w", err)
		}

		l.Debug("bank processed", slog.Int("part", 1), slog.Int("joltage", currJoltage))
		joltagePart1 += currJoltage
	}

	// Part 2
	joltagePart2 := 0
	for _, bank := range banks {
		currJoltage, err := getJoltagePart2(bank)
		if err != nil {
			return 0, 0, fmt.Errorf("getJoltagePart2() error: %w", err)
		}

		l.Debug("bank processed", slog.Int("part", 2), slog.Int("joltage", currJoltage))
		joltagePart2 += currJoltage
	}

	return joltagePart1, joltagePart2, nil
}

func getJoltagePart1(bank string) (int, error) {
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

func getJoltagePart2(bank string) (int, error) {
	var joltageBuffer strings.Builder
	leftIndex := 0
	for i := range 12 {
		newDigit, newIndex := findMax(bank[leftIndex : len(bank)-11+i])
		joltageBuffer.WriteString(strconv.Itoa(newDigit))

		// += since findMax() returns index relative to the slice passed in
		leftIndex += newIndex + 1
	}

	joltage, err := strconv.Atoi(joltageBuffer.String())
	if err != nil {
		return 0, fmt.Errorf("strconv.Atoi() failed: %w", err)
	}

	return joltage, nil
}

func findMax(digits string) (int, int) {
	max := int(digits[0] - byte('0'))
	maxIndex := 0
	for i, char := range digits {
		digit := int(char - rune('0'))
		if digit > max {
			max = digit
			maxIndex = i
		}
	}
	return max, maxIndex
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
