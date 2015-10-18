package main

import (
	"flag"
	"fmt"
)

var max int

func init() {
	flag.IntVar(&max, "max", 2000000, "sum all prime numbers below this max")
}

// Is a number prime
func isPrime(number int) bool {
	for i := 2; i < number; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

// Sum all Prime Numbers below the max
func sumOfPrimes(max int) int {
	result := 0

	for i := 2; i < max; i++ {
		if isPrime(i) {
			result += i
		}
	}

	return result
}

// Concurrent sum all Prime Numbers below the max
func sumOfPrimes2(max int) int {
	result := make(chan int)
	for i := 2; i < max; i++ {
		val := i
		go func() {
			if isPrime(val) {
				result <- val
			} else {
				result <- 0
			}
		}()
	}
	total := 0
	for i := 2; i < max; i++ {
		total += <-result
	}

	return total
}

// https://projecteuler.net/problem=10
func main() {
	flag.Parse()
	fmt.Printf("The sum of all the primes below %v is %v.\n\n", max, sumOfPrimes2(max))
}
