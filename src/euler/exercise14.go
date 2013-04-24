/*
The following iterative sequence is defined for the set of positive integers:

n -> n/2 (n is even)
n -> 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13  40  20  10  5  16  8  4  2  1

It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

package euler

func collatz(start uint64) (nterms uint64)  {
    nterms = uint64(1)
    for num := start;num != 1; nterms++ {
        if num % 2 == 0 {
            num = num / 2
        } else {
            num = 3*num + 1
        }
    }
    return
}

func MaxUInt64(a, b uint64) uint64 {
    if a > b {  
        return a
    }
    return b
}

func Exercise14() uint64 {
    chainlength := uint64(0)
    tmp := uint64(0)
    longestseed := uint64(0)
    for i:=1; i<1e6; i++ {
        tmp = collatz(uint64(i))
        if tmp > chainlength {
            chainlength = tmp
            longestseed = uint64(i)
        }
    }
    return longestseed
}
