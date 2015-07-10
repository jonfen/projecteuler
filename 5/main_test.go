package main

import (
	"testing"
)

var testSet = []struct {
	max    int
	result int
}{
	{10, 2520},
	{20, 232792560},
}

func TestSmallestDivisible(t *testing.T) {
	for _, tt := range testSet {
		if x := smallestDivisible(tt.max); x != tt.result {
			t.Errorf("Result for %v should be %v, not %v", tt.max, tt.result, x)
		}
	}
}

func TestSmallestDivisible2(t *testing.T) {
	for _, tt := range testSet {
		if x := smallestDivisible2(tt.max); x != tt.result {
			t.Errorf("Result for %v should be %v, not %v", tt.max, tt.result, x)
		}
	}
}

func BenchmarkSmallestDivisible(b *testing.B) {
	for n := 0; n < b.N; n++ {
		smallestDivisible(20)
	}
}

func BenchmarkSmallestDivisible2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		smallestDivisible2(20)
	}
}
