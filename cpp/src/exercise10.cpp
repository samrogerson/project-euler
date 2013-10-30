#include <iostream>

#include "euler-tools.h"

uint64_t Euler::exercise10() {
    std::vector<bool> primes = eratosthenes(2000000);
    uint64_t total = 0;
    for(std::vector<bool>::size_type i = 2; i < primes.size(); ++i) {
        if(primes[i]) {
            total += i;
        }
    }
    std::cout << std::endl;
    return total;
}
