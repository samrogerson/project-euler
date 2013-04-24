/*What is the largest prime factor of 600851475143*/

package euler

import (
	"euler/primes"
)

func Exercise3() []int64 {
	return primes.Factors(600851475143)
}
