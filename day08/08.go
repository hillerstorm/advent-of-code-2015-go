package day08

import (
	"strconv"
)

func count(lines []string) (int, int) {
	var p1Sum, p2Sum int
	for _, line := range lines {
		unquoted, _ := strconv.Unquote(line)
		p1Sum += (len(line) - len(unquoted))
		p2Sum += (len(strconv.Quote(line)) - len(line))
	}
	return p1Sum, p2Sum
}
