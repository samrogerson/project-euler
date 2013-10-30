#ifndef H_EulerTools
#define H_EulerTools

#include <cstdint>
#include <utility>
#include <ctime>
#include <vector>

#include "euler-defs.h"

namespace Euler {
    FunctionTime time_exercise(Exercise& e);
    uint64_t fib_next(std::pair<uint64_t,uint64_t>&);
    uint64_t max_factor(uint64_t n);
    bool is_palindrome(uint64_t);
    uint64_t gcd(uint64_t,uint64_t);
    uint64_t lcm(uint64_t,uint64_t);
    uint64_t sum_squares(uint64_t);
    uint64_t square_sum(uint64_t);
    bool is_prime(uint64_t);
    std::vector<uint64_t> pythagorean_triple(uint64_t);
}

#endif /* H_EulerTools */
