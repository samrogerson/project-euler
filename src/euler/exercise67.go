
package euler

import (
    "euler/triangles"
    "euler/web"
)

func Exercise67() int {
    tri := web.FetchText("http://projecteuler.net/project/triangle.txt")
    lines := triangles.Splittext(tri)
    nlines := len(lines)
    max, _ := triangles.Maxroute(lines[1:nlines])
    return max
}
