package utils

import (
    "io/ioutil"
    "net/http"
    "fmt"
)


// has to be a better way
func AccumulateUint64(seq []uint64) (total uint64) {
    for _, v := range seq {
        total += v
    }
    return
}

func AccumulateInt64(seq []int64) (total int64) {
    for _, v := range seq {
        total += v
    }
    return
}



func MaxIntOfSeq(vals []int) int {
    max := vals[0]
    for _, v := range vals {
        if v > max {
            max = v
        }
    }
    return max
}

func MaxInt64(a, b int64) int64 {
    if a > b {
        return a
    } else {
        return b
    }
    return 0
}

func MaxUInt64(a, b uint64) uint64 {
    if a > b {  
        return a
    }
    return b
}


func FetchText(url string) string {
    resp, _ := http.Get(url)
    defer resp.Body.Close()
    readresp, _ := ioutil.ReadAll(resp.Body)
    return string(readresp)
}


func ReadFileOrURL(fname, url string) string {
    filetext := ""
    b, err := ioutil.ReadFile(fname)
    if err != nil {
        fmt.Print("fetching file")
        resp, _ := http.Get(url)
        b, _ := ioutil.ReadAll(resp.Body)
        ioutil.WriteFile(fname, b, 0644)
    }

    filetext = string(b)

    return filetext
}
