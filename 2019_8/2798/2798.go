package main

import (
	"fmt"
)

func main() {
	var card [200]int
	var N, M int
	fmt.Scanln(&N, &M)

	for i := 0; i < N; i++ {
		var temp int
		fmt.Scan(&temp)
		card[i] = temp
	}

	var max int = 0
	for i := 0; i < N-2; i++ {
		for j := i + 1; j < N-1; j++ {
			for z := j + 1; z < N; z++ {
				var sum int = card[i] + card[j] + card[z]
				if sum <= M && sum > max {
					max = sum
				}
			}
		}
	}
	fmt.Println(max)
}
