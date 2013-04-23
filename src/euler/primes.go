package primes

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

func Factors(n int64) (factors []int64) {
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
