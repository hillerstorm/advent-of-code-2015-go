package day09

import (
	"math"
	"regexp"
	"strconv"
)

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}

var pattern = regexp.MustCompile("^(\\w+) to (\\w+) = (\\d+)$")

func findPaths(lines []string) (int, int) {
	paths := make(map[string]int, len(lines)*2)
	cities := make(map[string]struct{})
	for _, line := range lines {
		matches := pattern.FindStringSubmatch(line)
		dist, _ := strconv.Atoi(matches[3])
		paths[matches[1]+"->"+matches[2]] = dist
		paths[matches[2]+"->"+matches[1]] = dist
		cities[matches[1]] = struct{}{}
		cities[matches[2]] = struct{}{}
	}

	citiesSlice := make([]string, len(cities))
	indices := make([]int, len(cities))

	i := 0
	for k := range cities {
		citiesSlice[i] = k
		indices[i] = i
		i++
	}
	shortest := math.MaxInt32
	longest := 0
	for p := make([]int, len(indices)); p[0] < len(p); nextPerm(p) {
		permutation := getPerm(indices, p)
		distance := 0
		for ix, idx := range permutation {
			if ix != len(permutation)-1 {
				distance += paths[citiesSlice[idx]+"->"+citiesSlice[permutation[ix+1]]]
			}
		}
		if distance < shortest {
			shortest = distance
		}
		if distance > longest {
			longest = distance
		}
	}
	return shortest, longest
}
