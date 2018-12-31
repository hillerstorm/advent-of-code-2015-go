package day01

import (
	"testing"

	"github.com/hillerstorm/advent-of-code-2015-go/utils"
)

type Expect struct {
	input  string
	result int
	part   utils.Part
}

func TestDay01(t *testing.T) {
	input := utils.ReadFile("01.in")

	expects := []Expect{
		Expect{"(())", 0, utils.Part1},
		Expect{"()()", 0, utils.Part1},
		Expect{"(((", 3, utils.Part1},
		Expect{"(()(()(", 3, utils.Part1},
		Expect{"))(((((", 3, utils.Part1},
		Expect{"())", -1, utils.Part1},
		Expect{"))(", -1, utils.Part1},
		Expect{")))", -3, utils.Part1},
		Expect{")())())", -3, utils.Part1},
		Expect{")", 1, utils.Part2},
		Expect{"()())", 5, utils.Part2},
		Expect{input, 232, utils.Part1},
		Expect{input, 1783, utils.Part2},
	}

	for i, x := range expects {
		p1, p2 := traverseInput(x.input)
		var res int
		if x.part == utils.Part1 {
			res = p1
		} else {
			res = p2
		}
		if res != x.result {
			t.Errorf("01_p%d_t%d: Expected %d, got %d", x.part, i, x.result, res)
		}
	}
}
