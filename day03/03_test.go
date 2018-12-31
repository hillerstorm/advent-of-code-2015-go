package day03

import (
	"testing"

	"github.com/hillerstorm/advent-of-code-2015-go/utils"
)

type Expect struct {
	input  string
	result int
	part   utils.Part
}

func TestDay03(t *testing.T) {
	input := utils.ReadFile("03.in")

	expects := []Expect{
		Expect{">", 2, utils.Part1},
		Expect{"^>v<", 4, utils.Part1},
		Expect{"^v^v^v^v^v", 2, utils.Part1},
		Expect{"^v", 3, utils.Part2},
		Expect{"^>v<", 3, utils.Part2},
		Expect{"^v^v^v^v^v", 11, utils.Part2},
		Expect{input, 2572, utils.Part1},
		Expect{input, 2631, utils.Part2},
	}

	for i, x := range expects {
		res := deliverPresents(x.input, int(x.part))
		if res != x.result {
			t.Errorf("03_p%d_t%d: Expected %d, got %d", x.part, i, x.result, res)
		}
	}
}
