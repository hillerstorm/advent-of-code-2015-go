package day08

import (
	"testing"

	"github.com/hillerstorm/advent-of-code-2015-go/utils"
)

type Expect struct {
	input  []string
	result int
	part   utils.Part
}

func TestDay08(t *testing.T) {
	input := utils.ReadLines("08.in")

	lines := []string{
		"\"\"",
		"\"abc\"",
		"\"aaa\\\"aaa\"",
		"\"\\x27\"",
	}
	expects := []Expect{
		Expect{lines, 12, utils.Part1},
		Expect{lines, 19, utils.Part2},
		Expect{input, 1342, utils.Part1},
		Expect{input, 2074, utils.Part2},
	}

	for i, x := range expects {
		p1, p2 := count(x.input)
		var res int
		if x.part == utils.Part1 {
			res = p1
		} else {
			res = p2
		}
		if res != x.result {
			t.Errorf("08_p%d_t%d: Expected %d, got %d", x.part, i, x.result, res)
		}
	}
}
