package day06

import (
	"strconv"
	"strings"

	"github.com/hillerstorm/advent-of-code-2015-go/utils"
)

func parseRange(prefix string, line string) (int, int, int, int) {
	parts := strings.Split(strings.TrimPrefix(line, prefix), " through ")
	startStrings := strings.Split(parts[0], ",")
	start := make([]int, len(startStrings))
	for i, s := range startStrings {
		start[i], _ = strconv.Atoi(s)
	}
	endStrings := strings.Split(parts[1], ",")
	end := make([]int, len(endStrings))
	for i, s := range endStrings {
		end[i], _ = strconv.Atoi(s)
	}
	return start[0], start[1], end[0], end[1]
}

var rulesets = map[int]map[string](func(int) int){
	1: map[string](func(int) int){
		"turn on": func(_ int) int {
			return 1
		},
		"turn off": func(_ int) int {
			return 0
		},
		"toggle": func(value int) int {
			if value == 1 {
				return 0
			}
			return 1
		},
	},
	2: map[string](func(int) int){
		"turn on": func(value int) int {
			return value + 1
		},
		"turn off": func(value int) int {
			return utils.MaxInt(0, value-1)
		},
		"toggle": func(value int) int {
			return value + 2
		},
	},
}

func countLights(input []string, ruleset int) int {
	grid := make([]int, 1000*1000)

	rules := rulesets[ruleset]
	for _, line := range input {
		if strings.HasPrefix(line, "turn") {
			var x1, y1, x2, y2 int
			var valueFun func(int) int
			if strings.HasPrefix(line, "turn on") {
				x1, y1, x2, y2 = parseRange("turn on ", line)
				valueFun = rules["turn on"]
			} else {
				x1, y1, x2, y2 = parseRange("turn off ", line)
				valueFun = rules["turn off"]
			}
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					idx := 1000*y + x
					grid[idx] = valueFun(grid[idx])
				}
			}
		} else {
			x1, y1, x2, y2 := parseRange("toggle ", line)
			valueFun := rules["toggle"]
			for x := x1; x <= x2; x++ {
				for y := y1; y <= y2; y++ {
					idx := 1000*y + x
					grid[idx] = valueFun(grid[idx])
				}
			}
		}
	}

	sum := 0
	for _, x := range grid {
		sum += x
	}
	return sum
}
