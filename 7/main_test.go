package main

import (
	"testing"
)

var testSet = []struct {
	max    int
	result int
}{
	{10001, 104743},
}

func TestFib(t *testing.T) {
	for _, tt := range testSet {
		if x := nthPrime(tt.max); x != tt.result {
			t.Errorf("The largest prime factor of the number %v should be %v, not %v", tt.max, tt.result, x)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		nthPrime(10001)
	}
}
