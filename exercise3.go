/*What is the largest prime factor of 600851475143*/

package main

import (
    "fmt"
)

func factors(n int64) (factors []int64) {
    for n & 1 == 0 {
        factors = append(factors, int64(2))
        n = n / 2.
    }
    f := int64(3)
    for n >= f * f {
        if (n % f == 0) {
            factors = append(factors, int64(f))
            n = n / f
        } else {
            f += int64(1)
        }
    }

    if n != 1 {
        factors = append(factors,n)
    }

    return 
}

func main() {
    fmt.Println(factors(600851475143))
}
