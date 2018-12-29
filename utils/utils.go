package utils

import (
	"bufio"
	"os"
	"strconv"
)

// Part represents a puzzle part
type Part int

const (
	// Part1 represents part 1 of a puzzle
	Part1 Part = 1
	// Part2 represents part 2 of a puzzle
	Part2 Part = 2
)

// MinInt returns the minimum of two integers
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt returns the maximum of two integers
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// ReadFile returns each line of a file as a string
func ReadFile(path string) string {
	file, _ := os.Open(path)
	defer file.Close()

	var input string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input += scanner.Text()
	}

	return input
}

// ReadLines returns each line of a file as a string
func ReadLines(path string) []string {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}

// ReadLinesAsInt returns each line of a file as an integer
func ReadLinesAsInt(path string) []int {
	file, _ := os.Open(path)
	defer file.Close()

	var lines []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, line)
	}

	return lines
}
