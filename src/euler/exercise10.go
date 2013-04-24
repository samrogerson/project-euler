/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package euler

import (
    "euler/primes"
)

func Exercise10() int64 {
    p := primes.PrimesBelow(2e6)
    total := int64(0)
    for _, pv := range p {
        total += pv
    }
    return total
}
