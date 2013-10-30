#include "euler-tools.h"

// sum a + b + c = 1000
// a^2 + b^2 = c^2

uint64_t Euler::exercise9() {
    std::vector<uint64_t> result = pythagorean_triple(1000);
    return result[0]*result[1]*result[2];
}
