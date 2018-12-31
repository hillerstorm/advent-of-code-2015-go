package day09

import (
	"testing"

	"github.com/hillerstorm/advent-of-code-2015-go/utils"
)

type Expect struct {
	input  []string
	result int
	part   utils.Part
}

func TestDayTwo(t *testing.T) {
	input := utils.ReadLines("09.in")

	lines := []string{
		"London to Dublin = 464",
		"London to Belfast = 518",
		"Dublin to Belfast = 141",
	}
	expects := []Expect{
		Expect{lines, 605, utils.Part1},
		Expect{lines, 982, utils.Part2},
		Expect{input, 207, utils.Part1},
		Expect{input, 804, utils.Part2},
	}

	for i, x := range expects {
		shortest, longest := findPaths(x.input)
		var res int
		if x.part == utils.Part1 {
			res = shortest
		} else {
			res = longest
		}
		if res != x.result {
			t.Errorf("09_p%d_t%d: Expected %d, got %d", x.part, i, x.result, res)
		}
	}
}
