#ifndef H_EulerDefs
#define H_EulerDefs

#include <functional>
#include <vector>
#include <utility>
#include <cstdint>


namespace Euler{
    typedef std::function<uint64_t()> Exercise;
    typedef std::vector<Exercise> ExerciseCollection;
    typedef std::pair<uint64_t,double> FunctionTime;

    uint64_t exercise1();
    uint64_t exercise2();
    uint64_t exercise3();
    uint64_t exercise4();
    uint64_t exercise5();

}

#endif /* H_EulerDefs */
