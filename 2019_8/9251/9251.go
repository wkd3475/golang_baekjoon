package main

import (
	"fmt"
)

func main() {
	var first, second string
	var dp [1002][1002]int
	fmt.Scanln(&first)
	fmt.Scanln(&second)

	for i := 1; i <= len(first); i++ {
		for j := 1; j <= len(second); j++ {
			if first[i-1] == second[j-1] {
				//이 경우에는 매칭되는 문자가 있어서 최대 길이가 1 증가한다.
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				//i와 j만 변경되어 현재 상태가 될 수 있으며, 이 경우에는 서로 매칭이 안된 경우라서
				//최대 길이가 증가하지 않고 저장된 값을 이용해 최대값을 얻어 유추할 수 있다.
				dp[i][j] = Max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	fmt.Println(dp[len(first)][len(second)])
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
