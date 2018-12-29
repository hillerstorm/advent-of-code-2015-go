package day02

import (
	"regexp"
	"sort"
	"strconv"

	"github.com/hillerstorm/advent-of-code-2015-go/utils"
)

type present struct {
	l, w, h int
}

var digits = regexp.MustCompile("\\d+")

func parsePresents(input []string) []present {
	var presents []present
	for _, x := range input {
		res := digits.FindAllString(x, -1)
		l, _ := strconv.Atoi(res[0])
		w, _ := strconv.Atoi(res[1])
		h, _ := strconv.Atoi(res[2])
		presents = append(presents, present{l, w, h})
	}
	return presents
}

func wrappingPaper(input []string) int {
	presents := parsePresents(input)
	sqft := 0
	for _, x := range presents {
		sqft += (2 * x.l * x.w) + (2 * x.w * x.h) + (2 * x.h * x.l) + utils.MinInt(utils.MinInt(x.l*x.w, x.l*x.h), x.w*x.h)
	}
	return sqft
}

func ribbon(input []string) int {
	presents := parsePresents(input)
	ft := 0
	for _, x := range presents {
		values := []int{x.l, x.w, x.h}
		sort.Ints(values)
		min1, min2 := values[0], values[1]
		ft += min1 + min1 + min2 + min2 + (x.l * x.w * x.h)
	}
	return ft
}
