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

var words = map[string] int {
    "1": 3, "2": 3, "3": 5, "4": 4, "5": 4, "6": 3, "7": 5, "8": 5, "9": 4,
    "10": 3, "11": 6, "12": 6, "13": 8, "14": 8, "15": 7, "16": 7, "18": 8,
    "19": 4, "20": 6, "30": 6, "40": 5, "50": 5, "60": 5, "70": 7, "80": 6, "90": 6,
    "hundred": 7,
    "thousand": 8,
}

func countnumber(num int) (letters int) {
    runes := []rune(strconv.Itoa(num)) 

    nrunes := len(runes)
    ngroups := nrunes / 3

    groups := [][]rune{[]rune{0,0,0}}
    for i := 0; i<ngroups; i++ {
        fmt.Println(i, nrunes)
        groups[i] = []rune{runes[nrunes-3*i-3],runes[nrunes-3*i-2],runes[nrunes-3*i-1]}
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
    for i, g := range groups {
        fmt.Println(i, string(g))
    }
    return
}

func main() {
    countnumber(55544433)
}



    //sum := 0
    //val := 0
    //for i:=1; i<=1000; i++ {
        //val = countnumber(i)
        //fmt.Println(i, val)
        //sum += val
    //}
    //fmt.Println(sum)


    //runes := []rune(strconv.Itoa(num))
    //// reverse so we can work from the bottom up
	//for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		//runes[i], runes[j] = runes[j], runes[i]
	//} 
    //// do first 2 then last 1 or 2
    //val := ""
    //modifier := ""
    //letters = 0
    //for _, r  := range runes {
        //val = string(r)
        //if val != "0" {
            //if len(val+modifier) < 3 { // is a normal number
                //letters += words[val+modifier]
            //} else if len(val+modifier) < 4 {
                //letters += words[val]
            //} else {
                //letters += words[val]
            //}
        //} 
        //modifier += "0"
    //}
    //if len(runes) >= 3 {
        //if string(runes[0]) != "0" || string(runes[1]) != "0" {
            //letters += 3
        //}
        //if len(runes) >= 4 {
            //letters += words["thousand"]
        //} else {
            //letters += words["hundred"]
        //}
    //}
    //return letters
