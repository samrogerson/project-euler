#include "euler-tools.h"

uint64_t Euler::exercise5() {
    uint64_t l = 1;
    for(int i=2; i < 21; ++i) {
        l = lcm(l,i);
    }
    return l;
}
