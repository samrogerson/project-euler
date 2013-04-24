/*

Find the 10001st prime number

*/

package main

import (
	"euler/primes"
	"fmt"
)

func main() {
	fmt.Println("10001st prime is ", primes.NthPrime(10001))
}
