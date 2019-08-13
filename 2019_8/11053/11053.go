package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scanln(&N)
	var dp []int = make([]int, N+1)
	var input []int = make([]int, N+1)

	for i := 1; i <= N; i++ {
		fmt.Scan(&input[i])
	}

	dp[1] = 1
	for i := 2; i <= N; i++ {
		dp[i] = 1
		for j := 1; j < i; j++ {
			if input[i] == input[j] {
				dp[i] = dp[j]
			} else if input[i] > input[j] && dp[j] >= dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}

	max := 0
	for i := 1; i <= N; i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}

	fmt.Println(max)
}
