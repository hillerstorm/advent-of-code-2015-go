package day02

import (
	"testing"

	"github.com/hillerstorm/advent-of-code-2015-go/utils"
)

type Expect struct {
	input  []string
	result int
	part   utils.Part
}

func TestDay02(t *testing.T) {
	input := utils.ReadLines("02.in")

	expects := []Expect{
		Expect{[]string{"2x3x4"}, 58, utils.Part1},
		Expect{[]string{"1x1x10"}, 43, utils.Part1},
		Expect{[]string{"2x3x4"}, 34, utils.Part2},
		Expect{[]string{"1x1x10"}, 14, utils.Part2},
		Expect{input, 1586300, utils.Part1},
		Expect{input, 3737498, utils.Part2},
	}

	for i, x := range expects {
		var res int
		if x.part == utils.Part1 {
			res = wrappingPaper(x.input)
		} else {
			res = ribbon(x.input)
		}
		if res != x.result {
			t.Errorf("02_p%d_t%d: Expected %d, got %d", x.part, i, x.result, res)
		}
	}
}
