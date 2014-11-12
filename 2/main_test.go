package main

import (
	"testing"
)

var testSet = []struct {
	max    int
	result int
}{
	{4000000, 4613732},
}

func TestFib(t *testing.T) {
	for _, tt := range testSet {
		if x := fib(tt.max); x != tt.result {
			t.Errorf("Fibonacci sequence under %v should be %v, not %v", tt.max, tt.result, x)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(4000000)
	}
}
