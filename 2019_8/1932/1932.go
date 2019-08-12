package main

import (
	"fmt"
)

func main() {
	var N, max int
	fmt.Scanln(&N)
	var memory [550][550]int

	memory[0][0] = 0
	memory[0][1] = 0

	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			fmt.Scan(&memory[i][j])
			if j == 1 {
				memory[i][j] = memory[i-1][j] + memory[i][j]
			} else if j == i {
				memory[i][j] = memory[i-1][j-1] + memory[i][j]
			} else {
				memory[i][j] = Max(memory[i-1][j-1], memory[i-1][j]) + memory[i][j]
			}

			if memory[i][j] > max {
				max = memory[i][j]
			}
		}
	}

	fmt.Println(max)
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
