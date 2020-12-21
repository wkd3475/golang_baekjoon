package main

import "fmt"

func main() {
	var num, x, y int
	fmt.Scanf("%d", &num)
	for i:=1; i<=num; i++ {
		fmt.Scanf("%d %d", &x, &y)
		fmt.Printf("Case #%d: %d\n", i, x+y)
	}
}
