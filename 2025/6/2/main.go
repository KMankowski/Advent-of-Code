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
	lines := make([][]rune, 0)

	s := bufio.NewScanner(r)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		lines = append(lines, []rune(s.Text()))
	}

	operators := parseOperators(lines)

	var solutions []int
	var currProblemOperands [][]rune
	for col := 0; col < len(lines[0]); col++ {
		if isColumnAllSpaces(lines, col) {
			solution, err := solveProblem(currProblemOperands, operators[len(solutions)])
			if err != nil {
				return 0, fmt.Errorf("solveProblems() failed: %w", err)
			}
			solutions = append(solutions, solution)
			currProblemOperands = nil

			continue
		}

		nextOperand := make([]rune, 0)
		for row := 0; row < len(lines)-1; row++ {
			if lines[row][col] == ' ' {
				continue
			}
			nextOperand = append(nextOperand, lines[row][col])
		}
		currProblemOperands = append(currProblemOperands, nextOperand)
	}

	// For loop terminates before isColumnAllSpaces() is run against last problem
	solution, err := solveProblem(currProblemOperands, operators[len(solutions)])
	if err != nil {
		return 0, fmt.Errorf("solveProblems() failed: %w", err)
	}
	solutions = append(solutions, solution)

	total := 0
	for _, solution := range solutions {
		total += solution
	}

	return total, nil
}

func solveProblem(operands [][]rune, operator string) (int, error) {
	solution := 0
	for _, rawOperand := range operands {
		operand, err := strconv.Atoi(string(rawOperand))
		if err != nil {
			return 0, fmt.Errorf("strconv.Atoi() on operands %v failed: %w", operands, err)
		}

		if solution == 0 {
			solution = operand
		} else if operator == "*" {
			solution *= operand
		} else if operator == "+" {
			solution += operand
		}
	}
	return solution, nil
}

func isColumnAllSpaces(lines [][]rune, col int) bool {
	for row := 0; row < len(lines); row++ {
		if lines[row][col] != ' ' {
			return false
		}
	}
	return true
}

func parseOperators(lines [][]rune) []string {
	lastLine := lines[len(lines)-1]
	return strings.Fields(string(lastLine))
}
