package main

import "fmt"

func main() {
	var N, K int
	fmt.Scan(&N)
	fmt.Scan(&K)
	var V []int = make([]int, N+1)
	var dp []int = make([]int, 100001)

	for i := 1; i <= N; i++ {
		fmt.Scan(&V[i])
	}

	dp[0] = 1
	//각 동전 별로 가능한 경우가 생기면 이전 값에 +1 을 해주면서 테이블을 업데이트해준다.
	for i := 1; i <= N; i++ {
		for j := 1; j <= K; j++ {
			if j >= V[i] {
				dp[j] += dp[j-V[i]]
			}
		}
	}

	fmt.Println(dp[K])
}
