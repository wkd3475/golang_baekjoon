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
	check := 1
	for _, c := range in.Bytes() {
		if c == '-' {
			check = -1
		} else {
			r *= 10
			r += int(c - '0')
		}
	}
	return r * check
}

func main() {
	in.Split(bufio.ScanWords)
	N := nextInt()
	var r []int = make([]int, N+1)
	var c []int = make([]int, N+1)
	var dp [][]int = make([][]int, N+1)

	for i := 1; i <= N; i++ {
		r[i] = nextInt()
		c[i] = nextInt()
		dp[i] = make([]int, N+1)
	}

	for d := 1; d < N; d++ {
		for i := 1; i+d <= N; i++ {
			dp[i][i+d] = int(^uint(0) >> 1)
			for mid := i; mid < i+d; mid++ {
				dp[i][i+d] = Min(dp[i][i+d], dp[i][mid]+dp[mid+1][i+d]+r[i]*c[mid]*c[i+d])
			}
		}
	}

	fmt.Println(dp[1][N])
}

func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}
