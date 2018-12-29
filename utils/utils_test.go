package utils

import (
	"testing"
)

func TestMinInt(t *testing.T) {
	min := 0
	max := 4
	result := MinInt(min, max)
	if result != min {
		t.Errorf("Expected %d, got %d", min, result)
	}
	result = MinInt(max, min)
	if result != min {
		t.Errorf("Expected %d, got %d", min, result)
	}
	min = -5
	max = 3
	result = MinInt(max, min)
	if result != min {
		t.Errorf("Expected %d, got %d", min, result)
	}
	result = MinInt(min, max)
	if result != min {
		t.Errorf("Expected %d, got %d", min, result)
	}
}

func TestMaxInt(t *testing.T) {
	min := 0
	max := 4
	result := MaxInt(min, max)
	if result != max {
		t.Errorf("Expected %d, got %d", min, result)
	}
	result = MaxInt(max, min)
	if result != max {
		t.Errorf("Expected %d, got %d", min, result)
	}
	min = -5
	max = 3
	result = MaxInt(max, min)
	if result != max {
		t.Errorf("Expected %d, got %d", min, result)
	}
	result = MaxInt(min, max)
	if result != max {
		t.Errorf("Expected %d, got %d", min, result)
	}
}

func TestReadFile(t *testing.T) {
	expected := "0123456"
	input := ReadFile("test1.txt")
	if input != expected {
		t.Errorf("Expected %s, got %s", expected, input)
	}
}

func TestReadLines(t *testing.T) {
	expected := []string{"abc", "123"}
	input := ReadLines("test2.txt")
	if len(input) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(input))
	}
	for i := 0; i < len(input); i++ {
		if input[i] != expected[i] {
			t.Errorf("Expected %s, got %s", expected[i], input[i])
		}
	}
}

func TestReadLinesAsInt(t *testing.T) {
	expected := []int{123, 456}
	input := ReadLinesAsInt("test3.txt")
	if len(input) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(input))
	}
	for i := 0; i < len(input); i++ {
		if input[i] != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], input[i])
		}
	}
}
