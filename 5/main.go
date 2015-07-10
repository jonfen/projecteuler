package main

import (
	"flag"
	"fmt"
)

var max int

func init() {
	flag.IntVar(&max, "max", 20, "max range")
}

func smallestDivisible(max int) int {
	candidate := 3

	for n := max; n != 1; n-- {
		if candidate%n != 0 {
			candidate++
			n = max // reset
		}
	}
	return candidate
}

func smallestDivisible2(max int) int {
	candidate := 3

	for n := 2; n != max+1; n++ {
		if candidate%n != 0 {
			candidate++
			n = 2 // reset
		}
	}
	return candidate
}

// https://projecteuler.net/problem=5
func main() {
	flag.Parse()

	fmt.Printf("%v is the smallest positive number that is evenly divisible by all of the numbers from 2 to %v.\n", smallestDivisible(max), max)

}
