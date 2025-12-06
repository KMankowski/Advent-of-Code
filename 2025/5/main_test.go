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
		expPart2                   int
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
			14,
		},
		{
			"range expansion",
			`1-3
2-4
3-4

0`,
			0,
			4,
		},
		{
			"range merge low to high",
			`1-3
5-8
3-5

0`,
			0,
			8,
		},
		{
			"range merge high to low",
			`5-8
1-3
3-5

0`,
			0,
			8,
		},
		{
			"single value",
			`5-5

0`,
			0,
			1,
		},
		{
			"range encapsulates many ranges",
			`2-3
5-6
1-8

0`,
			0,
			8,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inpRawRangesAndIngredients := strings.NewReader(test.inpRawRangesAndIngredients)

			freshIngredients, part2, err := run(inpRawRangesAndIngredients)
			if err != nil {
				t.Fatalf("unexpected error: %v\n", err)
			}

			if freshIngredients != test.expFreshIngredients {
				t.Fatalf("expected freshIngredients %v but got %v", test.expFreshIngredients, freshIngredients)
			}

			if part2 != test.expPart2 {
				t.Fatalf("expected part2 %v but got %v", test.expPart2, part2)
			}
		})
	}
}
