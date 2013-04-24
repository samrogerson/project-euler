package euler

import (
    "fmt"
    "euler/web"
    "strings"
    "sort"
)

var LetterLookup = map[rune] int {
    []rune("A")[0] : 1,
    []rune("B")[0] : 2,
    []rune("C")[0] : 3,
    []rune("D")[0] : 4,
    []rune("E")[0] : 5,
    []rune("F")[0] : 6,
    []rune("G")[0] : 7,
    []rune("H")[0] : 8,
    []rune("I")[0] : 9,
    []rune("J")[0] : 10,
    []rune("K")[0] : 11,
    []rune("L")[0] : 12,
    []rune("M")[0] : 13,
    []rune("N")[0] : 14,
    []rune("O")[0] : 15,
    []rune("P")[0] : 16,
    []rune("Q")[0] : 17,
    []rune("R")[0] : 18,
    []rune("S")[0] : 19,
    []rune("T")[0] : 20,
    []rune("U")[0] : 21,
    []rune("V")[0] : 22,
    []rune("W")[0] : 23,
    []rune("X")[0] : 24,
    []rune("Y")[0] : 25,
    []rune("Z")[0] : 26,
}

func AlphabetValue(name string) int {
    //runes := []rune(name)
    return 1
}

func Exercise22() int {
    data := web.FetchText("http://projecteuler.net/project/names.txt")
    lines := strings.Split(data,",")
    sort.Strings(lines)
    fmt.Println(lines)
    return 0
}
