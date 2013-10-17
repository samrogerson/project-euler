#ifndef H_EulerTools
#define H_EulerTools

#include <cstdint>
#include <utility>
#include <ctime>

#include "euler-defs.h"

namespace Euler {
    FunctionTime time_exercise(Exercise& e);
    uint64_t fib_next(std::pair<uint64_t,uint64_t>&);
    bool is_palindrome(uint64_t);
    uint64_t gcd(uint64_t,uint64_t);
    uint64_t lcm(uint64_t,uint64_t);
}

#endif /* H_EulerTools */
