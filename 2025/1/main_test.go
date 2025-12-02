package main

import (
	"log/slog"
	"os"
	"slices"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name         string
		inpRotations string
		expMatches   int
	}{
		{
			"example",
			`L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`,
			3,
		},
		{
			"large numbers",
			`L50
R200
L50
L200
R50`,
			3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
			inpRotations := strings.NewReader(test.inpRotations)

			matches, err := run(inpRotations, logger)
			if err != nil {
				t.Fatalf("unexpected error: %v\n", err)
			}

			if matches != test.expMatches {
				t.Errorf("expMatches is %v but got %v", matches, test.expMatches)
			}
		})
	}
}

func TestParseRotations(t *testing.T) {
	tests := []struct {
		name               string
		inpRawRotations    string
		expParsedRotations []int
	}{
		{
			"example",
			`L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`,
			[]int{-68, -30, 48, -5, 60, -55, -1, -99, 14, -82},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inpRawRotations := strings.NewReader(test.inpRawRotations)

			parsedRotations, err := parseRotations(inpRawRotations)
			if err != nil {
				t.Fatalf("unexpected error: %v\n", err)
			}

			if !slices.Equal(parsedRotations, test.expParsedRotations) {
				t.Errorf("expParsedRotations is %v but got %v", test.expParsedRotations, parsedRotations)
			}
		})
	}
}
