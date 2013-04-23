//A palindromic number reads the same both ways. The largest palindrome made from
//the product of two 2-digit numbers is 9009 = 91 99.

//Find the largest palindrome made from the product of two 3-digit numbers.

package main

import (
	"fmt"
	"math"
)

var ndigits = 3

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	imin := int(math.Pow(10, float64(ndigits-1)))
	imax := int(math.Pow(10, float64(ndigits)) - 1)

	imin *= imin
	imax *= imax

	s := string("")
	r := string("")
	for i := imax; i > imin; i-- {
		s = fmt.Sprintf("%d", i)
		r = reverse(s)
		if s == r {
			break
		}
	}
	fmt.Println(s)
}
