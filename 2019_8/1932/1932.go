package main

import (
	"bufio"
	"fmt"
	"os"
)

var in = bufio.NewScanner(os.Stdin)

func nextInt() int {
	in.Scan()
	r := 0
	for _, c := range in.Bytes() {
		r *= 10
		r += int(c - '0')
	}
	return r
}

func main() {
	in.Split(bufio.ScanWords)
	var N, max int
	N = nextInt()
	var memory [550][550]int

	memory[0][0] = 0
	memory[0][1] = 0

	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			memory[i][j] = nextInt()
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
