package main

import (
	"fmt"
)

var mem1 [11]int
var mem2 [11][11]int
var N int

func main() {
	for i := 1; i < 10; i++ {
		temp := i
		j := 0
		for {
			mem2[i][j] = temp
			temp = (temp * i) % 10
			j++
			if temp == mem2[i][1] {
				break
			}
		}
		mem1[i] = j - 1
	}

	fmt.Scan(&N)
	var result []int = make([]int, N)
	i := 0
	for N > 0 {
		var a, b int

		fmt.Scan(&a)
		fmt.Scan(&b)

		a = a % 10

		if a == 0 {
			result[i] = 10
		} else {
			result[i] = mem2[a][(b-1)%mem1[a]] % 10
		}
		i++
		N--
	}

	for j := 0; j < i; j++ {
		fmt.Println(result[j])
	}
}

func Pow(a int, b int) int {
	result := 1
	for b > 0 {
		result *= a
		b--
	}
	return result
}
