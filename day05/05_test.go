package day05

import (
	"testing"

	"github.com/hillerstorm/advent-of-code-2015-go/utils"
)

type Expect struct {
	input  []string
	result int
	part   utils.Part
}

func TestDay05(t *testing.T) {
	input := utils.ReadLines("05.in")

	expects := []Expect{
		Expect{[]string{"ugknbfddgicrmopn"}, 1, utils.Part1},
		Expect{[]string{"aaa"}, 1, utils.Part1},
		Expect{[]string{"jchzalrnumimnmhp"}, 0, utils.Part1},
		Expect{[]string{"haegwjzuvuyypxyu"}, 0, utils.Part1},
		Expect{[]string{"dvszwmarrgswjxmb"}, 0, utils.Part1},
		Expect{[]string{"qjhvhtzxzqqjkmpb"}, 1, utils.Part2},
		Expect{[]string{"xxyxx"}, 1, utils.Part2},
		Expect{[]string{"uurcxstgmygtbstg"}, 0, utils.Part2},
		Expect{[]string{"ieodomkazucvgmuy"}, 0, utils.Part2},
		Expect{input, 255, utils.Part1},
		Expect{input, 55, utils.Part2},
	}

	for i, x := range expects {
		res := niceStrings(x.input, int(x.part))
		if len(res) != x.result {
			t.Errorf("05_p%d_t%d: Expected %d, got %d", x.part, i, x.result, len(res))
		}
	}
}
