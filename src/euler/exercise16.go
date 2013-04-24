/*

2^15 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 2^1000?

*/

package euler

import (
    "strconv"
    "math/big"
)

func Exercise16() int {
    num := big.NewInt(0)
    base := big.NewInt(2)
    exponent := big.NewInt(1000)
    num.Set(new(big.Int).Exp(base,exponent,nil))
    
    runes := []rune(num.String())
    total := 0
    val := 0
    for _, r := range runes {
        val, _ = strconv.Atoi(string(r)) 
        total += val
    }
    return total
}
