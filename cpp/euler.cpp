#include <iostream>
#include <cstdlib>

#include "euler-tools.h"

Euler::ExerciseCollection exercises = {
    Euler::exercise1, Euler::exercise2, Euler::exercise3, Euler::exercise4, Euler::exercise5,
    Euler::exercise6, Euler::exercise7, Euler::exercise8, Euler::exercise9,
};

int main(int argc, char* argv[]) {
    using Euler::FunctionTime;
    using Euler::ExerciseCollection;
    using Euler::time_exercise;

    if(argc > 1) {
        int exnum = atoi(argv[1]) - 1;
        FunctionTime f = time_exercise(exercises[exnum]);
        std::cout << "Exercise " << exnum+1 << ": " << f.first << " in " << f.second <<
            " seconds" << std::endl;
    } else {
        for(ExerciseCollection::size_type i = 0; i < exercises.size(); ++i) {
            FunctionTime f = time_exercise(exercises[i]);
            std::cout << "Exercise " << i+1 << ": " << f.first << " in " << f.second <<
                " seconds" << std::endl;
        }
    }
    return 0;
}
