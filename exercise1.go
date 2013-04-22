/*Find the sum of the mutliple of 3 or 5 below 1000*/

package main

import (
    "fmt"
)

func SumDivisors(divisors []int, max int) (total int) {
    for i := 0; i < max; i+=1 {
        for _, div := range divisors {
            if i % div == 0 {
                total += i
                break
            }
        }
    }
    return
}

func main() {
    divisors := make([]int,2)
    divisors[0] = 3
    divisors[1] = 5
    fmt.Println(divisors)

    total := SumDivisors(divisors, 1000)
    fmt.Println(total)
}
