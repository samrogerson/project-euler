#include "euler-tools.h"

uint64_t Euler::exercise7() {
    // find 10,001st prime
    uint64_t n_primes = 1;
    uint64_t current = 3;
    for(uint64_t i=3; n_primes < 10001; i+=2) {
        if(is_prime(i)) {
            current = i;
            n_primes++;
        }
    }
    return current;
}
