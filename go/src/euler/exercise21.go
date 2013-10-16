/*
Let d(n) be defined as the sum of proper divisors of n (numbers less than n
which divide evenly into n).
If d(a) = b and d(b) = a, where a  b, then a and b are an amicable pair and each
of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55
and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and
142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/


package euler

import (
    "euler/primes"
    "euler/utils"
)

var lookup = make([]uint64,10001)

func Exercise21() uint64 {
    sum := uint64(0)
    for num := uint64(1); num<uint64(10000); num++ {
        factors := primes.AllFactors(num)
        factors = factors[0:len(factors)-1] // shave off "self"
        dn := utils.AccumulateUint64(factors)
        if dn < uint64(10000) {
            if lookup[dn] == num {
                sum += num + dn
            }
            lookup[num] = dn
        }
    }
    return sum
}
