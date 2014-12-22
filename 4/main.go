package main

import (
	"flag"
	"fmt"
	"strconv"
)

// Digits is the number of digits in the factors
var Digits int

func init() {
	flag.IntVar(&Digits, "digits", 3, "number of digits")
}

func numeral(digits int) (int, error) {
	var numeral string
	for i := 1; i <= digits; i++ {
		numeral += "9"
	}

	return strconv.Atoi(numeral)
}

func checkRange(numeral int) int {
	digits := len(strconv.Itoa(numeral))

	var str string
	for i := 1; i <= digits-1; i++ {
		str += "9"
	}

	retval, _ := strconv.Atoi(str)

	return retval
}

func palindrome(numeral int) int {
	offset := checkRange(numeral)

	for left := numeral; left > numeral-offset; left-- {
		for right := numeral; right > numeral-offset; right-- {
			if isPalindrome(left * right) {
				return left * right
			}

		}
	}

	return 0
}

func isPalindrome(candidate int) bool {
	str := strconv.Itoa(candidate)
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}

	return true
}

// https://projecteuler.net/problem=3
func main() {
	flag.Parse()
	fmt.Printf("A palindromic number reads the same both ways.  The largest palidrome made from the product of two 2-digit numbers is 9009 = 91 x 99.\n\nFind the largest palindrome made from the product of two 3-digit numbers.")

	num, _ := numeral(Digits)
	fmt.Println("The largest palindrome made from a", Digits, "-digit number is", palindrome(num))
}
