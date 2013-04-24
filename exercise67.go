
package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "euler/triangles"
)

func main() {
    triweb, _ := http.Get("http://projecteuler.net/project/triangle.txt")
    defer triweb.Body.Close()
    tritext, _ := ioutil.ReadAll(triweb.Body)
    tri := string(tritext)
    lines := triangles.Splittext(tri)
    nlines := len(lines)
    max, _ := triangles.Maxroute(lines[1:nlines])
    fmt.Println(max)

}
