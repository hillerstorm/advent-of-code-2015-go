package day06

import (
	"testing"

	"github.com/hillerstorm/advent-of-code-2015-go/utils"
)

type Expect struct {
	input  []string
	result int
	part   utils.Part
}

func TestDay06(t *testing.T) {
	input := utils.ReadLines("06.in")

	expects := []Expect{
		Expect{[]string{"turn on 0,0 through 999,999"}, 1000 * 1000, utils.Part1},
		Expect{[]string{"toggle 0,0 through 999,0"}, 1000, utils.Part1},
		Expect{[]string{
			"turn on 0,0 through 999,999",
			"turn off 499,499 through 500,500",
		}, 999996, utils.Part1},
		Expect{[]string{"turn on 0,0 through 0,0"}, 1, utils.Part2},
		Expect{[]string{"toggle 0,0 through 999,999"}, 2000000, utils.Part2},
		Expect{input, 400410, utils.Part1},
		Expect{input, 15343601, utils.Part2},
	}

	for i, x := range expects {
		res := countLights(x.input, int(x.part))
		if res != x.result {
			t.Errorf("06_p%d_t%d: Expected %d, got %d", x.part, i, x.result, res)
		}
	}
}
