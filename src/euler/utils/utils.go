package utils

import (
    "io/ioutil"
    "net/http"
    "fmt"
)

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
