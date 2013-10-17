#include <cstdint>
#include <iostream>

#include "euler-tools.h"

uint64_t Euler::exercise3() {
    uint64_t n = 600851475143;
    while(n % 2 == 0) n /= 2;
    uint64_t max_fact = 0; 
    for(uint64_t fact = 3; n > fact * fact; fact++) {
        while(n % fact == 0)  {
            n /= fact;
            max_fact = fact;
        }
    }
    if(n != 1) max_fact = n;
    return max_fact;
}
