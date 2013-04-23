/*
A Pythagorean triplet is a set of three natural numbers, a  b  c, for which,

a^2 + b^2 = c^2
For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

/*

a^2 + b^2 - c^2 = 0
a + b +  = 1000c  = 1000

*/


package main

import (
    "fmt"
)

func pythag(total int) (a,b,c int) {
    for a= 1; a < total - 5; a++ {
        for b=a+1; b + a < total - 3; b++ {
            for c= b+1; c < total - 3;  c++ {
                if a+b+c == 1000 && (a*a + b*b == c*c) {
                    return a,b,c
                }
            }
        }
    }
    return 0,0,0
}



func main() {
    a,b,c := pythag(1000)
    fmt.Println(a,b,c)
    fmt.Println(a*b*c)
}
