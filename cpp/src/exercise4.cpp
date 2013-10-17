#include "euler-tools.h" 

#include <iostream>
uint64_t Euler::exercise4() {
    uint64_t max_palindrome = 0;
    for(int i=100; i<1e3; ++i) {
        for(int j=100; j<1e3; ++j) {
            uint64_t n = i*j;
            if(is_palindrome(n) && n > max_palindrome) max_palindrome = n;
        }
    }
    return max_palindrome;
}
