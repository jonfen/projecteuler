package main

import (
	"flag"
	"fmt"
)

var max int

func init() {
	flag.IntVar(&max, "max", 1000, "max value for multiples")
}

// Loop through each number from 1 to max.
func oneLoop(max int) int {
	total := 0
	values := []int{5, 3}

	for i := 1; i < max; i++ {
		for _, m := range values {
			if i%m == 0 {
				total += i
				break
			}
		}
	}
	return total
}

// Loop automatically skips non-multiples using a sequence.
func patternLoop(max int) int {
	total := 0

	// I compared the results of multiples and noticed this incremental pattern
	pattern := []int{3, 2, 1, 3, 1, 2, 3}

	for i := 0; i < max; {
		for _, v := range pattern {
			i += v
			if i >= max {
				break
			}
			total += i
		}
	}

	return total
}

// https://projecteuler.net/problem=1
func main() {
	flag.Parse()
	fmt.Printf("If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.  \nFind the sum of all the multiples of 3 or 5 below %v.\n\n", max)

	//total := oneLoop(max)
	//fmt.Printf("Total of multiples of 3 or 5 below %v: %v\n", max, total)

	total := patternLoop(max)
	fmt.Printf("Total of multiples of 3 or 5 below %v: %v\n", max, total)
}
