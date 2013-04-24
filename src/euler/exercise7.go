/*

Find the 10001st prime number

*/

package euler

import (
	"euler/primes"
)

func Exercise7() int64 {
	return primes.NthPrime(10001)
}
