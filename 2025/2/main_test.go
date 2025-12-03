package main

import (
	"log/slog"
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name      string
		inpRanges string
		expSum    int
	}{
		{
			"example",
			"11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124",
			1227775554,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			logger := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
			inpRanges := strings.NewReader(test.inpRanges)

			sum, err := run(inpRanges, logger)
			if err != nil {
				t.Fatalf("unexpected error: %v\n", err)
			}

			if sum != test.expSum {
				t.Fatalf("expected sum %v but got %v", test.expSum, sum)
			}
		})
	}
}
