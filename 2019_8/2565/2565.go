package main

import (
	"fmt"
	"sort"
)

type Pair struct {
	first  int
	second int
}

func main() {
	var N, max int
	fmt.Scanln(&N)
	var dp []int = make([]int, N)
	var input []Pair = make([]Pair, N)

	for i := 0; i < N; i++ {
		fmt.Scanf("%d %d", &input[i].first, &input[i].second)
	}

	sort.Slice(input, func(i, j int) bool {
		return input[i].first < input[j].first
	})

	for i := 0; i < N; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if input[i].second > input[j].second {
				if dp[j] == dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
		max = Max(max, dp[i])
	}

	fmt.Println(N - max)
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
