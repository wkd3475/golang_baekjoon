package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N)
	fmt.Scan(&K)
	var W []int = make([]int, N+1)
	var V []int = make([]int, N+1)
	var dp [][]int = make([][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([]int, 100001)
	}

	for i := 1; i <= N; i++ {
		fmt.Scan(&W[i])
		fmt.Scan(&V[i])
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= K; j++ {
			if j >= W[i] {
				dp[i][j] = Max(dp[i-1][j], dp[i-1][j-W[i]]+V[i])
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	fmt.Println(dp[N][K])
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
