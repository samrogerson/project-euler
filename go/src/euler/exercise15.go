/*
Starting in the top left corner of a 22 grid, and only being able to move to the
right and down, there are exactly 6 routes to the bottom right corner.
    _______
    |  |  |
    |  |  |
    -------
    |  |  |
    |  |  |
    -------

How many such routes are there through a 20 x 20 grid?

(rotate at 45 and count possible routes to each one: follow binomial rule)

(1,1), (1,2,1), (1,3,3,1), ...

*/

package euler

var ngrid = 20

func expandseq(seq []int64) (next []int64) {
    next = make([]int64, len(seq)+1)
    next[0] = 1
    next[len(next)-1] = 1
    for i :=  1; i < len(next) -1; i++ {
        next[i] = seq[i-1] + seq[i]
    }
    return
}

func reduceseq(seq []int64) (next []int64) {
    next = make([]int64, len(seq)-1)
    for i := 0; i < len(next); i++ {
        next[i] = seq[i] + seq[i+1]
    }
    return
}

func Exercise15() int64 {
    start := []int64{1}
    seq := start
    for ; len(seq) < ngrid+1; {
        seq = expandseq(seq)
    }
    for ; len(seq) > 1; {
        seq = reduceseq(seq)
    }
    return seq[0]
}
