package main

import (
	"flag"
	"fmt"
)

var nth int

func init() {
	flag.IntVar(&nth, "nth", 10001, "return the nth prime number")
}

//Is a number prime
func isPrime(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

// Calculate the Largest Prime Factor of a given number
func nthPrime(number int) int {
	for current, count := 2, 0; ; current++ {
		if isPrime(current) {
			count++
			if count >= number {
				return current
			}
		}
	}
}

// https://projecteuler.net/problem=7
func main() {
	flag.Parse()
	fmt.Printf("By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.\n\n")

	result := nthPrime(nth)
	fmt.Printf("The %vth prime is %v.\n\n", nth, result)
}
