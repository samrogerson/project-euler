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
