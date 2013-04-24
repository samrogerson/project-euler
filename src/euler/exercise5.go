//2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.
//What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

package euler

func gcd(a, b int64) (t int64) {
    for b != 0 {
        t = b
        b = a % b
        a = t
    }
    return
}

func lcm(a, b int64) int64 {
    return (a * b / gcd(a, b))
}

func Exercise5() int64 {
    l := int64(1)
    for i:=2; i < 21; i++ {
        l = lcm(l,int64(i)) 
    }
    return l
}
