package main

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name           string
		inpRawProblems []string
		expTotal       int
	}{
		{
			"example",
			[]string{
				"123 328  51 64 ",
				" 45 64  387 23 ",
				"  6 98  215 314",
				"*   +   *   +  ",
			},
			3263827,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Input is []string instead of multiline literal to avoid IDE trimming trailing spaces
			inpRawProblems := ""
			for _, s := range test.inpRawProblems {
				inpRawProblems += s
				inpRawProblems += "\n"
			}

			total, err := run(strings.NewReader(inpRawProblems))
			if err != nil {
				t.Fatalf("unexpected error: %v\n", err)
			}

			if total != test.expTotal {
				t.Fatalf("expected total %v but got %v", test.expTotal, total)
			}
		})
	}
}
