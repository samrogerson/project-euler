package main

import (
    "euler"
    "reflect"
    "fmt"
    "flag"
)

var ex = flag.Int("ex", 1, "The exercise you want to run")

var funcs = map[int] interface{}{
    1:  euler.Exercise1,
    2:  euler.Exercise2,
    3:  euler.Exercise3,
    4:  euler.Exercise4,
    5:  euler.Exercise5,
    6:  euler.Exercise6,
    7:  euler.Exercise7,
    8:  euler.Exercise8,
    9:  euler.Exercise9,
    10: euler.Exercise10,
    11: euler.Exercise11,
    12: euler.Exercise12,
    13: euler.Exercise13,
    14: euler.Exercise14,
    15: euler.Exercise15,
    16: euler.Exercise16,
    17: euler.Exercise17,
    18: euler.Exercise18,
    19: euler.Exercise19,
    20: euler.Exercise20,
    21: euler.Exercise21,
    22: euler.Exercise22,
    67: euler.Exercise67,
} 

func main() {
    flag.Parse()

    _, check := funcs[*ex]
    if check {
        fmt.Print("Running Exercise ", *ex, "... ")
        f := reflect.ValueOf(funcs[*ex])
        result := f.Call([]reflect.Value{})[0]
        k := result.Kind()
        if k >= reflect.Int && k <= reflect.Int64 {
            fmt.Println(result.Int())
        } else if k >= reflect.Uint && k <= reflect.Uint64 {
            fmt.Println(result.Uint())
        }
    } else {
        fmt.Println("Specified invalid Exercise number", *ex)
    }
}
