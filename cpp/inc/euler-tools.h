#ifndef H_EulerTools
#define H_EulerTools

#include <cstdint>
#include <utility>
#include <ctime>

#include "euler-defs.h"

namespace Euler {
    FunctionTime time_exercise(Exercise& e);
    int fib_next(std::pair<int,int>&);
    bool is_palindrome(uint64_t);
}

#endif /* H_EulerTools */
