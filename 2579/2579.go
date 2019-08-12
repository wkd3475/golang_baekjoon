package main

import (
	"fmt"
)

func main() {
	var N, input int
	fmt.Scanln(&N)
	var memory [305][2]int

	memory[0][0] = 0
	memory[0][1] = 0
	for i := 1; i <= N; i++ {
		fmt.Scanln(&input)
		if i == 1 {
			memory[i][0] = input
			memory[i][1] = 0
		} else if i == 2 {
			memory[i][0] = memory[i-1][0] + input
			memory[i][1] = input
		} else {
			memory[i][0] = memory[i-1][1] + input
			memory[i][1] = Max(memory[i-2][0], memory[i-2][1]) + input
		}
	}

	fmt.Println(Max(memory[N][0], memory[N][1]))
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
