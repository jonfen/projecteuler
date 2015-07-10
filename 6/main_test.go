package main

import (
	"testing"
)

var testSet = []struct {
	max     int
	sum     int
	squares int
	result  int
}{
	{10, 385, 3025, 2640},
	{100, 338350, 25502500, 25164150},
}

func TestSumOfSquares(t *testing.T) {
	for _, tt := range testSet {
		if x := sumOfSquares(tt.max); x != tt.sum {
			t.Errorf("Result for %v should be %v, not %v", tt.max, tt.sum, x)
		}
	}
}

func TestSquaresOfNumbers(t *testing.T) {
	for _, tt := range testSet {
		if x := squaresOfNumbers(tt.max); x != tt.squares {
			t.Errorf("Result for %v should be %v, not %v", tt.max, tt.squares, x)
		}
	}
}

func BenchmarkSumOfSquares(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sumOfSquares(100)
	}
}

func BenchmarkSmallestDivisible2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		squaresOfNumbers(100)
	}
}
