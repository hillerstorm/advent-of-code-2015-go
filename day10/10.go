package day10

import (
	"strconv"
)

func expand(startSequence string, times int) int {
	sequence := make([]int, len(startSequence))
	for i, x := range startSequence {
		sequence[i], _ = strconv.Atoi(string(x))
	}
	for ; times > 0; times-- {
		var newSequence []int
		var curDigit, digitCount int
		maxIdx := len(sequence) - 1
		for i, x := range sequence {
			if i == 0 {
				curDigit = x
				digitCount = 1
			} else if x != curDigit {
				newSequence = append(newSequence, digitCount, curDigit)
				curDigit = x
				digitCount = 1
			} else {
				digitCount++
			}
			if i == maxIdx {
				newSequence = append(newSequence, digitCount, curDigit)
			}
		}

		sequence = newSequence
	}
	return len(sequence)
}
