package main

import (
	"flag"
	"fmt"
)

var max int

func init() {
	flag.IntVar(&max, "max", 4000000, "max value for Fibonacci sequence")
}

//Sum of even Fibonacci sequence numbers under a max value
func fib(max int) int {
	total := 0

	first := 0
	second := 1

	for {
		tmp := second
		second += first
		first = tmp
		if first > max || second > max {
			break
		} else if second%2 == 0 {
			total += second
		}
	}

	return total
}

// https://projecteuler.net/problem=2
func main() {
	flag.Parse()
	fmt.Printf("Each new term in the Fibonacci sequence is generated by adding the previous two terms. By starting with 1 and 2, the first 10 terms will be:\n\n1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...\n\nBy considering the terms in the Fibonacci sequence whose values do not exceed four million (%v), find the sum of the even-valued terms.\n\n", max)

	result := fib(max)
	fmt.Println("The sum of even-valued Fibonacci terms under", max, "is", result)
}