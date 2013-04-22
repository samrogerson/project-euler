//A palindromic number reads the same both ways. The largest palindrome made from
//the product of two 2-digit numbers is 9009 = 91 99.

//Find the largest palindrome made from the product of two 3-digit numbers.

package main

import  (
    "fmt"
    "strconv"
)

func palindrome(s string) bool {
    return s == reverse(s)
}

func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes) -1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func main() {
    best := 0
    for i := 999; i > 99; i-- {
        for j := 999; j > 99; j-- {
            if palindrome(strconv.Itoa(j * i)) {
                if j*i > best {
                    best = j * i
                }
            }
        }
    }
    fmt.Println(best)
}
