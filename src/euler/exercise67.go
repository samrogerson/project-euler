
package euler

import (
    "euler/triangles"
    "euler/utils"
)

func Exercise67() int {
    tri := utils.ReadFileOrURL("docs/exercise67.dat", "http://projecteuler.net/project/triangle.txt")
    lines := triangles.Splittext(tri)
    nlines := len(lines)
    max, _ := triangles.Maxroute(lines[1:nlines])
    return max
}
