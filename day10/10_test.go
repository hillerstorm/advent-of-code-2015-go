package day10

import (
	"testing"
)

type Expect struct {
	input  string
	times  int
	result int
}

func TestDay10(t *testing.T) {
	input := "1113222113"

	expects := []Expect{
		Expect{"1", 1, 2},
		Expect{"1", 2, 2},
		Expect{"1", 3, 4},
		Expect{"1", 4, 6},
		Expect{"1", 5, 6},
		Expect{input, 40, 252594},
		Expect{input, 50, 3579328},
	}

	for i, x := range expects {
		res := expand(x.input, x.times)
		if res != x.result {
			t.Errorf("10_t%d: Expected %d, got %d", i, x.result, res)
		}
	}
}
