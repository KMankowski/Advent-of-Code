package main

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name          string
		inpRawGrid    string
		expRollsPart1 int
		expRollsPart2 int
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
			43,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inpRawGrid := strings.NewReader(test.inpRawGrid)

			rollsPart1, rollsPart2, err := run(inpRawGrid)
			if err != nil {
				t.Fatalf("unexpected error: %v\n", err)
			}

			if rollsPart1 != test.expRollsPart1 {
				t.Fatalf("expected rollsPart1 %v but got %v", test.expRollsPart1, rollsPart1)
			}

			if rollsPart2 != test.expRollsPart2 {
				t.Fatalf("expected rollsPart2 %v but got %v", test.expRollsPart2, rollsPart2)
			}
		})
	}
}
