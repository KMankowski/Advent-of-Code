package main

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name           string
		inpRawProblems string
		expTotal       int
	}{
		{
			"example",
			`123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +  `,
			4277556,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inpRawProblems := strings.NewReader(test.inpRawProblems)

			total, err := run(inpRawProblems)
			if err != nil {
				t.Fatalf("unexpected error: %v\n", err)
			}

			if total != test.expTotal {
				t.Fatalf("expected total %v but got %v", test.expTotal, total)
			}
		})
	}
}
