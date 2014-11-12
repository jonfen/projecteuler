package main

import (
	"flag"
	"fmt"
)

var max int

func init() {
	flag.IntVar(&max, "max", 600851475143, "calculate largest prime factor for this value")
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
func primeFactor(number int) int {
	largest := number

	for i := 2; i <= number/2; i++ {
		if largest%i == 0 {
			tmp := largest / i
			if isPrime(tmp) {
				return tmp
			}
		}
	}

	return largest
}

// https://projecteuler.net/problem=3
func main() {
	flag.Parse()
	fmt.Printf("The prime factors of 13195 are 5, 7, 13 and 29.\n\nWhat is the largest prime factor of the number %v ?\n\n", max)

	result := primeFactor(max)
	fmt.Println("The largest prime factor for", max, "is", result)
}
