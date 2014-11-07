package main

import (
	"testing"
)

var testSet = []struct {
	max    int
	result int
}{
	{10, 23},
	{1000, 233168},
}

func TestOneLoop(t *testing.T) {
	for _, tt := range testSet {
		if x := oneLoop(tt.max); x != tt.result {
			t.Errorf("Mulitples under %v should be %v, not %v", tt.max, tt.result, x)
		}
	}
}

func TestPatternLoop(t *testing.T) {
	for _, tt := range testSet {
		if x := patternLoop(tt.max); x != tt.result {
			t.Errorf("Mulitples under %v should be %v, not %v", tt.max, tt.result, x)
		}
	}
}

func BenchmarkOneLoop(b *testing.B) {
	for n := 0; n < b.N; n++ {
		oneLoop(1000)
	}
}

func BenchmarkPatternLoop(b *testing.B) {
	for n := 0; n < b.N; n++ {
		patternLoop(1000)
	}
}
