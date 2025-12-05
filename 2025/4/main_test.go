package main

import (
	"log/slog"
	"os"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name       string
		inpRawGrid string
		expRolls   int
	}{
		{
			"example",
			`..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.`,
			13,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			l := slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{Level: slog.LevelDebug}))
			inpRawGrid := strings.NewReader(test.inpRawGrid)

			rolls, err := run(inpRawGrid, l)
			if err != nil {
				t.Fatalf("unexpected error: %v\n", err)
			}

			if rolls != test.expRolls {
				t.Fatalf("expected rolls %v but got %v", test.expRolls, rolls)
			}
		})
	}
}
