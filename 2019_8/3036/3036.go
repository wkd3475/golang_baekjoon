package main

import (
	"fmt"
)

func main() {
	var T int
	var others [100]int
	var gcd [100]int
	fmt.Scanln(&T)

	var first int
	fmt.Scan(&first)

	for i := 0; i < T-1; i++ {
		fmt.Scan(&others[i])
		gcd[i] = GCD(first, others[i])
	}

	for i := 0; i < T-1; i++ {
		fmt.Printf("%d/%d\n", first/gcd[i], others[i]/gcd[i])
	}
}

func GCD(n int, r int) int {
	if r > n {
		r, n = n, r
	}

	for r != 0 {
		var temp = n
		n = r
		r = temp % r
	}

	return n
}
