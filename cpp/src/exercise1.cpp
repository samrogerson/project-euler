#include "euler-tools.h"

uint64_t Euler::exercise1() {
    int sum(0);
    for(int i=3; i<1000; ++i) {
        if(i%3==0 || i%5==0) sum+=i;
    }
    return sum;
}
