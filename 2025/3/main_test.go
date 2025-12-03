package main

import (
	"log/slog"
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name        string
		inpRawBanks string
		expJoltage1 int
		expJoltage2 int
	}{
		{
			"example",
			`987654321111111
811111111111119
234234234234278
818181911112111`,
			357,
			3121910778619,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inpRawBanks := strings.NewReader(test.inpRawBanks)
			logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))

			joltage1, joltage2, err := run(inpRawBanks, logger)
			if err != nil {
				t.Fatalf("unexpected error: %v\n", err)
			}

			if joltage1 != test.expJoltage1 {
				t.Fatalf("expected joltage1 %v but got %v", test.expJoltage1, joltage1)
			}

			if joltage2 != test.expJoltage2 {
				t.Fatalf("expected joltage2 %v but got %v", test.expJoltage2, joltage2)
			}
		})
	}
}
