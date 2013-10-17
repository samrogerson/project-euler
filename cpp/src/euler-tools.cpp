#include "euler-tools.h"

Euler::FunctionTime Euler::time_exercise(Exercise& e) {
    clock_t start, end;
    start = clock();
    int result = e();
    end = clock();
    return {result, (double)(end-start)/CLOCKS_PER_SEC};
}

int Euler::fib_next(std::pair<int,int>& fib_pair) {
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
