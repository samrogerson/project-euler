#include <iostream>
#include <functional>
#include <vector>
#include <utility>

#include <cstdlib>
#include <ctime>

namespace Euler{
    typedef std::function<int()> Exercise;
    typedef std::vector<Exercise> ExerciseCollection;
    typedef std::pair<int,double> FunctionTime;

    int exercise1();
    int exercise2();
    
    FunctionTime time_exercise(Exercise& e) {
        clock_t start, end;
        start = clock();
        int result = e();
        end = clock();
        return {result, (double)(end-start)/CLOCKS_PER_SEC};
    }
}

Euler::ExerciseCollection exercises = {
    Euler::exercise1, Euler::exercise2,
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
