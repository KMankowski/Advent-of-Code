package main

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name           string
		inpRawManifold string
		expTimelines   int
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
			40,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inpRawManifold := strings.NewReader(test.inpRawManifold)

			timelines := run(inpRawManifold)

			if timelines != test.expTimelines {
				t.Fatalf("expected timelines %v but got %v", test.expTimelines, timelines)
			}
		})
	}
}
