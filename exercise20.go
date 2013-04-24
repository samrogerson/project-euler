/*
n! means n x (n - 1) x  ... x 3 x 2 x 1

For example, 10! = 10  9  ...  3  2  1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

package main

import (
    "fmt"
    "math/big"
    "strconv"
)

func Factorial(n int64) (total big.Int) {
    total.MulRange(2,n)
    return total
}

func main() {
    result := Factorial(100)
    rstring := result.String()
    runes := []rune(rstring)
    total := 0
    for _, r := range runes {
        temp, _ := strconv.Atoi(string(r))
        total += temp
    }
    fmt.Println(total)
}
