/* f(n) = f(n-1) + f(n-2); consider f(n) < 4e6; sum all even f(n) */

package main

import "fmt"

func fibo(n1, n2 int) (next int) {
	next = n1 + n2
	return
}

func main() {
	n_t := 0
	n_1 := 1
	n := 2
	total := 2
	for n < 4e6 {
		n_t = n
		n = fibo(n, n_1)
		n_1 = n_t
		if n%2 == 0 {
			total += n
		}
	}
	fmt.Println(total)

}
