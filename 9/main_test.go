package main

import (
	"testing"
)

var testSet = []struct {
	max    int
	result int
}{
	{1000, 31875000},
}

func TestFunc(t *testing.T) {
	for _, tt := range testSet {
		if x := tripletOfSum(tt.max).product(); x != tt.result {
			t.Errorf("The product %v should be %v, not %v", tt.max, tt.result, x)
		}
	}
}

func BenchmarkFunc(b *testing.B) {
	for n := 0; n < b.N; n++ {
		tripletOfSum(1000)
	}
}
