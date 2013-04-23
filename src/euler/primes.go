package primes

import "math"

func IsPrime(n int64) (prime bool) {
    return len(Factors(n)) == 1
}

func NthPrime(n int) (p int64) {
    found := 1
    for i := int64(3); found<=n; i+= int64(2) {
        if IsPrime(int64(i)) {
            found += 1
            p = int64(i)
        }
    }
    return
}

func PrimesBelow(n int64) (primes []int64) {
    primes = make([]int64,1)
    primes[0] = 2
    for i := int64(3); i<n; i += 2 {
        if IsPrime(i) {
            primes = append(primes,i)
        }
    }
    return primes
}

func Factors(n int64) (factors []int64) {
    if n <= int64(0) {
        return
    }
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

func AllFactors(n uint64) (factors []uint64) {
    for i:=uint64(1); i<=n; i++ {
        if n % i == 0 {
            factors = append(factors,i)
        }
    }
    return 
}

func NumFactors(n uint64) (nfac uint64) {
    nfac = uint64(0)
    for i:=uint64(2); i<=uint64(math.Sqrt(float64(n))); i++ {
        if n % i == 0 {
            nfac++
        }
    }
    nfac *=2
    return 
}
