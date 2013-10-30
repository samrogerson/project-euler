#include "euler-tools.h"

Euler::FunctionTime Euler::time_exercise(Exercise& e) {
    clock_t start, end;
    start = clock();
    int result = e();
    end = clock();
    return {result, (double)(end-start)/CLOCKS_PER_SEC};
}

uint64_t Euler::fib_next(std::pair<uint64_t,uint64_t>& fib_pair) {
    return fib_pair.first + fib_pair.second;
}

uint64_t Euler::max_factor(uint64_t n) {
    while(n % 2 == 0) n /= 2;
    uint64_t max_fact = 0; 
    for(uint64_t fact = 3; n >= fact * fact; fact++) {
        while(n % fact == 0)  {
            n /= fact;
            max_fact = fact;
        }
    }
    if(n != 1) max_fact = n;
    return max_fact;
}

bool Euler::is_palindrome(uint64_t n) {
    uint64_t num = n;
    uint64_t rev = 0;
    while(num > 0) {
        uint64_t dig = num % 10;
        rev = rev * 10 + dig;
        num /= 10;
    }
    return n == rev;
}

uint64_t Euler::gcd(uint64_t a, uint64_t b) {
    uint64_t t=0;
    while(b != 0) {
        t = b;
        b = a % b;
        a = t;
    }
    return t;
}

uint64_t Euler::lcm(uint64_t a, uint64_t b) {
    return (a * b / gcd(a, b));
}

uint64_t Euler::sum_squares(uint64_t n) {
    uint64_t total = 0;
    for(uint64_t i=1; i<=n; ++i) {
        total += i*i;
    }
    return total;
}

uint64_t Euler::square_sum(uint64_t n) {
    uint64_t total = (n+1) * n / 2;
    return total*total;
}

bool Euler::is_prime(uint64_t n) {
    return Euler::max_factor(n) == n;
}

std::vector<uint64_t> Euler::pythagorean_triple(uint64_t total) {
    for(uint64_t a = 1; a < total - 5; ++a) {
        for(uint64_t b=a+1; b + a < total - 3; ++b) {
            for(uint64_t c = b+1; c < total - 3;  ++c) {
                if((a+b+c == total) && (a*a + b*b == c*c)) {
                    return std::vector<uint64_t>{a,b,c};
                }
            }
        }
    }
    return std::vector<uint64_t>{0,0,0};
}
