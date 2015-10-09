package main

import (
	"testing"
)

var testSet = []struct {
	max    int
	result int
}{
	{13, 23514624000},
}

func TestFib(t *testing.T) {
	for _, tt := range testSet {
		if x := largestProduct(tt.max); x != tt.result {
			t.Errorf("The largest product factor %v should be %v, not %v", tt.max, tt.result, x)
		}
	}
}

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		largestProduct(13)
	}
}
