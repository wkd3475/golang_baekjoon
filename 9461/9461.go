package main

import (
	"fmt"
)

func main() {
	var mem = make([]int, 105)
	var T, N int
	fmt.Scanln(&T)

	for T > 0 {
		fmt.Scanln(&N)

		mem[1] = 1
		mem[2] = 1
		mem[3] = 1
		mem[4] = 2
		mem[5] = 2
		var result int = F(mem, N)
		fmt.Println(result)
		T--
	}
}

func F(mem []int, n int) int {
	if mem[n] == 0 {
		mem[n] = F(mem, n-1) + F(mem, n-5)
		return mem[n]
	} else {
		return mem[n]
	}
}
