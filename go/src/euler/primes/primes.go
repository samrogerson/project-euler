package primes

import (
    "math"
)

func IsPrime(n int64) (prime bool) {
    return len(PrimeFactors(n)) == 1
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

func Eratosthenes(n int64) []int64 {
    sievemax := (n / 2. - 1)
    upperlim := (int64(math.Sqrt(float64(n))) -1) / 2

    // going to reverse logic because of default initialization of bool
    // false => prime, true => checked off
    primearray := make([]bool,sievemax+1)
    for i := int64(1); i <= upperlim; i++ {
        if !primearray[i] {
            for j := i * 2 * (i + 1); j <= sievemax; j += 2 * i + 1 {
                primearray[j] = true
            }
        }
    }

    primes := make([]int64,int64(float64(n) / (math.Log(float64(n)) - 1.08366)));
    primes[0] = 2
    pcount := 1
    for i := int64(1); i <= sievemax; i++ {
        if !primearray[i] {
            primes[pcount] = 2*i + 1
            pcount += 1 
        }
    }

    return primes[:pcount]
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

func PrimeFactors(n int64) (factors []int64) {
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
