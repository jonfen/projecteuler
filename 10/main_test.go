package main

import (
	"testing"
)

var testSet = []struct {
	max    int
	result int
}{
	{200000, 1709600813},
	{2000000, 142913828922},
}

func TestFunc(t *testing.T) {
	for _, tt := range testSet {
		if x := sumOfPrimes(tt.max); x != tt.result {
			t.Errorf("The largest prime factor of the number %v should be %v, not %v", tt.max, tt.result, x)
		}
	}
}

func TestFunc2(t *testing.T) {
	for _, tt := range testSet {
		if x := sumOfPrimes2(tt.max); x != tt.result {
			t.Errorf("The largest prime factor of the number %v should be %v, not %v", tt.max, tt.result, x)
		}
	}
}

func BenchmarkFunc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sumOfPrimes(2000000)
	}
}

func BenchmarkFunc2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		sumOfPrimes2(2000000)
	}
}
