/*
The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural
numbers and the square of the sum is 3025  385 = 2640.

Find the difference between the sum of the squares of the first one hundred
natural numbers and the square of the sum.
*/

package main

import "fmt"

func SumOfSquares(n int) (sum int64) {
    sum = int64(0)
    for i:= int64(1); i<=int64(n); i++ {
        sum += i*i
    }
    return
}

func SquareOfSum(n int) (sum int64) {
    sum = (int64(n+1)*int64(n)) / 2
    sum *= sum
    return
}

func main() {
    fmt.Println(SquareOfSum(100)-SumOfSquares(100))
}

