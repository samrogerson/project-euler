package main

import (
    "euler"
    "reflect"
    "fmt"
)



func main() {
    funcs := []interface{}{euler.Exercise1, euler.Exercise2, euler.Exercise3,
    euler.Exercise4, euler.Exercise5, euler.Exercise6, euler.Exercise7,
    euler.Exercise8, euler.Exercise9, euler.Exercise10, euler.Exercise11,
    euler.Exercise12, euler.Exercise13, euler.Exercise14, euler.Exercise15,
    euler.Exercise16, euler.Exercise17, euler.Exercise18, euler.Exercise19,
    euler.Exercise20, euler.Exercise21 } 
    //euler.Exercise67 }

    nfuncs := len(funcs)
    for i, fi := range funcs[nfuncs-1:nfuncs] {
        f := reflect.ValueOf(fi)
        result := f.Call([]reflect.Value{})[0]
        k := result.Kind()
        if k >= reflect.Int && k <= reflect.Int64 {
            fmt.Println(i+1, "=>", result.Int())
        } else if k >= reflect.Uint && k <= reflect.Uint64 {
            fmt.Println(i+1, "=>", result.Uint())
        }
    }
}
