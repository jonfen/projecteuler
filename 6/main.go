package main

import (
	"flag"
	"fmt"
)

var max int

func init() {
	flag.IntVar(&max, "max", 100, "max range")
}

func sumOfSquares(max int) int {
	sum := 0

	for n := 1; n <= max; n++ {
		sum += n * n
	}

	return sum
}

func squaresOfNumbers(max int) int {
	sum := 0

	for n := 1; n <= max; n++ {
		sum += n
	}

	return sum * sum
}

// https://projecteuler.net/problem=6
func main() {
	flag.Parse()
	sum := sumOfSquares(max)
	squares := squaresOfNumbers(max)

	fmt.Printf("The difference between the sum of the squares of the first %v natural numbers and the square of the sum is %v - %v = %v\n", max, squares, sum, squares-sum)

}
