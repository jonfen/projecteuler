package main

import (
	"flag"
	"fmt"
	"math"
)

type triplet struct {
	a, b, c int
}

func (t triplet) product() int {
	return t.a * t.b * t.c
}

var sum int

func init() {
	flag.IntVar(&sum, "sum", 1000, "sum of Pythagorean triplet")
}

func tripletOfSum(sum int) triplet {
	for a := 1; a < sum; a++ {
		for b := 1; b < sum; b++ {
			c := math.Sqrt(float64(a*a + b*b))
			if math.Trunc(c) == c {
				if a+b+int(c) == sum {
					return triplet{a: a, b: b, c: int(c)}
				}
			}
		}
	}
	return triplet{a: 0, b: 0, c: 0}
}

// https://projecteuler.net/problem=9
func main() {
	flag.Parse()
	t := tripletOfSum(sum)
	fmt.Printf("There exists exactly one Pythagorean triplet (a, b, c) for which %v + %v + %v = %v.  The product of abc is %v.\n\n", t.a, t.b, t.c, sum, t.product())
}
