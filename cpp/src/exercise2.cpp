#include <utility>

namespace Euler {
    int fib_next(std::pair<int,int>& fib_pair) {
        return fib_pair.first + fib_pair.second;
    }

    int exercise2() {
        std::pair<int,int> fib_nums({1,1});
        int sum = 0;
        while(fib_nums.second <= 4e6) {
            if(fib_nums.second % 2 == 0) sum+=fib_nums.second;
            int next = fib_next(fib_nums);
            fib_nums.first = fib_nums.second;
            fib_nums.second = next;
        }
        return sum;
    }
}
