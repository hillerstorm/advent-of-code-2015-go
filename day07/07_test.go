package day07

import (
	"testing"

	"github.com/hillerstorm/advent-of-code-2015-go/utils"
)

type Expect struct {
	input       []string
	wireToCheck string
	result      uint16
	part        utils.Part
}

func TestDay07(t *testing.T) {
	input := utils.ReadLines("07.in")

	circuit := []string{
		"123 -> x",
		"456 -> y",
		"x AND y -> d",
		"x OR y -> e",
		"x LSHIFT 2 -> f",
		"y RSHIFT 2 -> g",
		"NOT x -> h",
		"NOT y -> i",
	}
	expects := []Expect{
		Expect{circuit, "d", 72, utils.Part1},
		Expect{circuit, "e", 507, utils.Part1},
		Expect{circuit, "f", 492, utils.Part1},
		Expect{circuit, "g", 114, utils.Part1},
		Expect{circuit, "h", 65412, utils.Part1},
		Expect{circuit, "i", 65079, utils.Part1},
		Expect{circuit, "x", 123, utils.Part1},
		Expect{circuit, "y", 456, utils.Part1},
		Expect{input, "a", 46065, utils.Part1},
		Expect{input, "a", 14134, utils.Part2},
	}

	for i, x := range expects {
		p1, p2 := reduce(x.input, x.wireToCheck)
		var res uint16
		if x.part == utils.Part1 {
			res = p1
		} else {
			res = p2
		}
		if res != x.result {
			t.Errorf("07_p%d_t%d: Expected %d, got %d", x.part, i, x.result, res)
		}
	}
}
