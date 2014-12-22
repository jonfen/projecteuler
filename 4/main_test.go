package main

import (
	"testing"
)

var testSet = []struct {
	digits int
	result int
}{
	{2, 9009},
	{3, 906609},
}

/*
func TestPal(t *testing.T) {
	for _, tt := range testSet {
		num, _ := numeral(tt.digits)
		if x := palindrome(num); x != tt.result {
			t.Errorf("The largest palindrome of %v-digits should be %v, not %v", tt.digits, tt.result, x)
		}
	}
}
*/
func TestIsPalindrome(t *testing.T) {
	var samples = []struct {
		num    int
		result bool
	}{
		{9009, true},
		{906609, true},
		{9008, false},
	}

	for _, tt := range samples {
		if x := isPalindrome(tt.num); x != tt.result {
			t.Errorf("It is %v that %v is a palindrome, not %v", tt.result, tt.num, x)
		}
	}
}

func BenchmarkPal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		num, _ := numeral(3)
		palindrome(num)
	}
}
