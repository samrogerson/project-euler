package triangles

import (
    "strings"
    "strconv"
)
func Splittext(s string) (linearrays [][]int) {
    lines := strings.Split(s,"\n")
    nlines := len(lines)
    linearrays = make([][]int,nlines)

    for lnum, line := range lines {
        linesplit := strings.Split(line," ")
        linearrays[nlines-1-lnum] = make([]int,len(linesplit))
        for decnum, dec := range linesplit {
            linearrays[nlines-1-lnum][decnum], _ = strconv.Atoi(dec)
        }
    }
    return
}

func MaxInt(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func Maxroute(linearrays [][]int) (int, [][]int) {
    nlines := len(linearrays)
    totals := make([][]int,nlines)
    totals[0] = linearrays[0]

    for lnum := 1; lnum < nlines; lnum++ {
        linelength := len(linearrays[lnum])
        totals[lnum] = make([]int, linelength)
        for intnum := range totals[lnum] {
            totals[lnum][intnum] = linearrays[lnum][intnum] +
                MaxInt(totals[lnum-1][intnum],totals[lnum-1][intnum+1])
        }
    }

    return totals[nlines-1][0], totals
}
