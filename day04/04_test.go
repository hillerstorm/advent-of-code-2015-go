package day04

import (
	"strings"
	"testing"

	"github.com/hillerstorm/advent-of-code-2015-go/utils"
)

type Expect struct {
	input  string
	result int
	part   utils.Part
}

func TestDayTwo(t *testing.T) {
	input := "ckczppom"

	expects := []Expect{
		Expect{"abcdef", 609043, utils.Part1},
		Expect{"pqrstuv", 1048970, utils.Part1},
		Expect{input, 117946, utils.Part1},
		Expect{input, 3938038, utils.Part2},
	}

	for i, x := range expects {
		res := firstMatch(x.input, strings.Repeat("0", 4+int(x.part)))
		if res != x.result {
			t.Errorf("04_p%d_t%d: Expected %d, got %d", x.part, i, x.result, res)
		}
	}
}
