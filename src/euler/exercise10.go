/*
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package euler

import (
    "euler/primes"
    "euler/utils"
)

func Exercise10() int64 {
    p := primes.Eratosthenes(2e6)
    total := utils.AccumulateInt64(p)
    return total
}
