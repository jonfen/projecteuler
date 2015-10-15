package main

import (
	"flag"
	"fmt"
)

var max int

func init() {
	flag.IntVar(&max, "max", 2000000, "sum all prime numbers below this max")
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

// Sum all Prime Numnbers below the max
func sumOfPrimes(max int) int {
	result := 0

	for i := 2; i < max; i++ {
		if isPrime(i) {
			result += i
		}
	}

	return result
}

// https://projecteuler.net/problem=7
func main() {
	flag.Parse()
	fmt.Printf("The sum of all the primes below %v is %v.\n\n", max, sumOfPrimes(max))
}
