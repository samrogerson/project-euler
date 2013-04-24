/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package main

import (
    "fmt"
    "euler/primes"
)

func main() {
    p := primes.PrimesBelow(2e6)
    total := int64(0)
    for _, pv := range p {
        total += pv
    }
    fmt.Println(total)
}
