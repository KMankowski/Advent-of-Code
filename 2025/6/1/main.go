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

const INPUT_FILE = "./2025/6/input.txt"

var VALID_OPERATORS = []string{"*", "+"}

func main() {
	l := slog.New(slog.NewTextHandler(os.Stdout, nil))

	input, err := os.Open(INPUT_FILE)
	if err != nil {
		l.Error("os.Open() failed", "error", err)
		os.Exit(1)
	}

	total, err := run(input)
	if err != nil {
		l.Error("run() failed", "error", err)
	}

	l.Info("success", slog.Int("total", total))
}

func run(r io.Reader) (int, error) {
	operands, operators, err := parseProblems(r)
	if err != nil {
		return 0, fmt.Errorf("parseProblems() failed: %w", err)
	}

	solutions := make([]int, len(operands[0]))
	copy(solutions, operands[0])
	for row := 1; row < len(operands); row++ {
		for col := 0; col < len(operands[row]); col++ {
			switch operators[col] {
			case "*":
				solutions[col] *= operands[row][col]
			case "+":
				solutions[col] += operands[row][col]
			}
		}
	}

	total := 0
	for _, solution := range solutions {
		total += solution
	}

	return total, nil
}

func parseProblems(r io.Reader) ([][]int, []string, error) {
	operands := make([][]int, 0)
	var operators []string

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		columns := strings.Fields(s.Text())
		if doesContainString(VALID_OPERATORS, columns[0]) {
			operators = columns
			break
		}

		nextOperands := make([]int, 0)
		for _, rawOperand := range columns {
			parsedOperand, err := strconv.Atoi(rawOperand)
			if err != nil {
				return nil, nil, fmt.Errorf("strconv.Atoi() on row %v failed: %w", columns, err)
			}
			nextOperands = append(nextOperands, parsedOperand)
		}
		operands = append(operands, nextOperands)
	}

	return operands, operators, nil
}

func doesContainString(list []string, s string) bool {
	for _, curr := range list {
		if curr == s {
			return true
		}
	}
	return false
}
