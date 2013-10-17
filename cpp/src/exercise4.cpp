#include "euler-tools.h" 

#include <iostream>
namespace Euler {
    int exercise4() {
        int max_palindrome = 0;
        for(int i=100; i<1e3; ++i) {
            for(int j=100; j<1e3; ++j) {
                int n = i*j;
                if(is_palindrome(n) && n > max_palindrome) max_palindrome = n;
            }
        }
        return max_palindrome;
    }
}
