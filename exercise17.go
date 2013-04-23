/*
If the numbers 1 to 5 are written out in words: one, two, three, four, five,
then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in
words, how many letters would be used?
*/

package main

import (
    "fmt"
    "strconv"
    //"strings"
)

var hundred = 7

var words = map[string] int {
    "1": 3, "2": 3, "3": 5, "4": 4, "5": 4, "6": 3, "7": 5, "8": 5, "9": 4,
    "10": 3, "11": 6, "12": 6, "13": 8, "14": 8, "15": 7, "16": 7, "17": 9, "18": 8,
    "19": 8, "20": 6, "30": 6, "40": 5, "50": 5, "60": 5, "70": 7, "80": 6, "90": 6,
}

var orders = map[int] int {
    1 : 0,
    2 : 8,  // len("thousand")
    3 : 7,  // len("million")
    4 : 7,  // len("billion")
}

func SeparateByThrees(num int) [][]rune {
    runes := []rune(strconv.Itoa(num)) 

    nrunes := len(runes)
    ngroups := nrunes / 3

    groups := make([][]rune,ngroups)

    offset :=0
    for i := 0; i<ngroups; i++ {
        offset = 3*(i+1)
        groups[i] = []rune{runes[nrunes-offset],runes[nrunes-offset+1],runes[nrunes-offset+2]}
    }

    if ngroups * 3 != nrunes {
        remaining := nrunes % 3 
        extra := []rune{}
        for i:=0; i<remaining; i++ {
            extra = append(extra,runes[i])
        }
        groups = append(groups, extra)
    }
    /* 
        { hhh } hundreds
        { ttt } hundreds of thousands 
        { mmm } hundreds of millions
    */

    return groups
}

func NumberAsLetterCount(groups [][]rune) int {
    total := 0

    for i := 1; i <= len(groups); i++ { // adding millions / billions
        total += orders[i]
    }

    s := ""
    for _, g := range groups {
        s = string(g[0])
        if len(g) == 3 {
            if s != "0" {
                total += words[s] + hundred  // X hundred
            }
            if string(g[1]) != "0" || string(g[2]) != "0" {
                total += 3 //and
            }
        }
        if len(g) >=2 {
            // get the last 2 digits
            numpair := []rune{ g[len(g)-2], g[len(g)-1] } // e.g. {X27} -> {27}
            if string(numpair[0]) ==  "0" { //less than ten
                total += words[string(numpair[1])]
            } else if string(numpair[0]) == "1" { // teen
                total += words[string(numpair)]
            } else { //2X, 3X, 4X...
                total += words[string(numpair[0])+"0"] + words[string(numpair[1])]
            }
        } else {
            total += words[s]
        }
        
    }
    return total
}


func main() {
    totalletters :=0
    temp := 0
    for num := 1; num<=1000; num++ {
        groups := SeparateByThrees(num)
        temp = NumberAsLetterCount(groups)
        totalletters += temp
        fmt.Println(num, temp)
    }
    fmt.Println(totalletters)
}
