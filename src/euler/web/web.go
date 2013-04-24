
package web

import (
    "net/http"
    "io/ioutil"
)

func FetchText(url string) string {
    resp, _ := http.Get(url)
    defer resp.Body.Close()
    readresp, _ := ioutil.ReadAll(resp.Body)
    return string(readresp)
}
