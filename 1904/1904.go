package main

import (
	"fmt"
)

func main() {
	var mem = make([]int, 1000000)
	var N int
	fmt.Scanln(&N)

	mem[1] = 1
	mem[2] = 2
	mem[3] = 3
	var result int = F(mem, N)
	fmt.Println(result)
}

func F(mem []int, n int) int {
	if mem[n] == 0 {
		mem[n] = (F(mem, n-2) + F(mem, n-1)) % 15746
		return mem[n]
	} else {
		return mem[n] % 15746
	}
}
