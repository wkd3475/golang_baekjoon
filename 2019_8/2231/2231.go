package main

import (
	"fmt"
)

func main() {
	var min int
	var N int
	fmt.Scanln(&N)

	i := N
	j := 0
	for i > 0 {
		i = i / 10
		j++
	}

	min = N

	for i := N - 9*j; i < N; i++ {
		result := i
		value := i
		for value > 0 {
			result += value % 10
			value /= 10
		}

		if result == N && i < min {
			min = i
		}
	}

	if min == N {
		fmt.Println(0)
	} else {
		fmt.Println(min)
	}
}
