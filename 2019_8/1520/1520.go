package main

import (
	"bufio"
	"fmt"
	"os"
)

var N, M int
var mapArray [][]int
var dp [][]int
var a = [4]int{1, -1, 0, 0}
var b = [4]int{0, 0, 1, -1}

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
	N = nextInt()
	M = nextInt()
	mapArray = make([][]int, N+2)
	dp = make([][]int, N+2)

	for i := 0; i <= N; i++ {
		dp[i] = make([]int, M+2)
		mapArray[i] = make([]int, M+2)
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			mapArray[i][j] = nextInt()
		}
	}

	for i := 0; i <= N; i++ {
		for j := 0; j <= M; j++ {
			dp[i][j] = -1
		}
	}

	fmt.Println(dfs(N, M))
}

func dfs(x int, y int) int {
	if dp[x][y] != -1 {
		return dp[x][y]
	}
	if x == 1 && y == 1 {
		return 1
	}

	dp[x][y] = 0

	for i := 0; i < 4; i++ {
		nx := x + a[i]
		ny := y + b[i]

		if nx > N || nx < 1 || ny > M || ny < 1 {
			continue
		}

		if mapArray[x][y] < mapArray[nx][ny] {
			dp[x][y] += dfs(nx, ny)
		}
	}
	return dp[x][y]
}
