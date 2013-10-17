#ifndef H_EulerDefs
#define H_EulerDefs

#include <functional>
#include <vector>
#include <utility>


namespace Euler{
    typedef std::function<int()> Exercise;
    typedef std::vector<Exercise> ExerciseCollection;
    typedef std::pair<int,double> FunctionTime;

    int exercise1();
    int exercise2();
    int exercise3();
    int exercise4();

}

#endif /* H_EulerDefs */
