package day05

import (
	"regexp"
	"strings"

	"github.com/dlclark/regexp2"
)

var vowelsPattern = regexp.MustCompile("[aeiou]")
var doublePattern = regexp2.MustCompile("(\\w)\\1", 0)
var forbidden = []string{"ab", "cd", "pq", "xy"}
var doublePattern2 = regexp2.MustCompile("(\\w)\\w\\1", 0)
var pairsPattern = regexp2.MustCompile("(\\w)(\\w).*\\1\\2", 0)

var rulesets = map[int][](func(string) bool){
	1: [](func(string) bool){
		func(line string) bool {
			vowelCount := vowelsPattern.FindAllString(line, 3)
			return vowelCount != nil && len(vowelCount) >= 3
		},
		func(line string) bool {
			matchesDouble, _ := doublePattern.MatchString(line)
			return matchesDouble
		},
		func(line string) bool {
			for _, x := range forbidden {
				if strings.Contains(line, x) {
					return false
				}
			}
			return true
		},
	},
	2: [](func(string) bool){
		func(line string) bool {
			matchesDouble, _ := doublePattern2.MatchString(line)
			return matchesDouble
		},
		func(line string) bool {
			matchesPairs, _ := pairsPattern.MatchString(line)
			return matchesPairs
		},
	},
}

func niceStrings(input []string, ruleset int) []string {
	var nice []string
	rules := rulesets[ruleset]
	for _, line := range input {
		isNice := true
		for _, rule := range rules {
			if !rule(line) {
				isNice = false
				break
			}
		}
		if isNice {
			nice = append(nice, line)
		}
	}
	return nice
}
