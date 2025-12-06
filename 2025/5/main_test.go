package main

import (
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name                       string
		inpRawRangesAndIngredients string
		expFreshIngredients        int
	}{
		{
			"example",
			`3-5
10-14
16-20
12-18

1
5
8
11
17
32`,
			3,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inpRawRangesAndIngredients := strings.NewReader(test.inpRawRangesAndIngredients)

			freshIngredients, err := run(inpRawRangesAndIngredients)
			if err != nil {
				t.Fatalf("unexpected error: %v\n", err)
			}

			if freshIngredients != test.expFreshIngredients {
				t.Fatalf("expected freshIngredients %v but got %v", test.expFreshIngredients, freshIngredients)
			}
		})
	}
}
