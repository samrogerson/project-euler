/*What is the largest prime factor of 600851475143*/

package euler

import (
	"euler/primes"
)

func Exercise3() int64 {
    factors := primes.PrimeFactors(600851475143)
    return factors[len(factors)-1]
}
