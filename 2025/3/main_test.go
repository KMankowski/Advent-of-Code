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
		expJoltage  int
	}{
		{
			"example",
			`987654321111111
811111111111119
234234234234278
818181911112111`,
			357,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inpRawBanks := strings.NewReader(test.inpRawBanks)
			logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))

			joltage, err := run(inpRawBanks, logger)
			if err != nil {
				t.Fatalf("unexpected error: %v\n", err)
			}

			if joltage != test.expJoltage {
				t.Fatalf("expected joltage %v but got %v", test.expJoltage, joltage)
			}
		})
	}
}
