#include <utility>

#include "euler-tools.h"

uint64_t Euler::exercise2() {
    std::pair<uint64_t,uint64_t> fib_nums({1,1});
    uint64_t sum = 0;
    while(fib_nums.second <= 4e6) {
        if(fib_nums.second % 2 == 0) sum+=fib_nums.second;
        uint64_t next = fib_next(fib_nums);
        fib_nums.first = fib_nums.second;
        fib_nums.second = next;
    }
    return sum;
}
