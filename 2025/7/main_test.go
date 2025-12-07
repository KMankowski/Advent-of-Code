package main

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name           string
		inpRawManifold string
		expSplits      int
	}{
		{
			"example",
			`.......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............`,
			21,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inpRawManifold := strings.NewReader(test.inpRawManifold)

			splits := run(inpRawManifold)

			if splits != test.expSplits {
				t.Fatalf("expected splits %v but got %v", test.expSplits, splits)
			}
		})
	}
}
