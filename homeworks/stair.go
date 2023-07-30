package main

import (
	"fmt"
)

func countWays(n int) int {
	if n <= 1 {
		return 1
	}

	ways := make([]int, n+1)
	ways[0], ways[1] = 1, 1

	for i := 2; i <= n; i++ {
		ways[i] = ways[i-1] + ways[i-2]

		if i >= 5 {
			ways[i] += ways[i-5]
		}
	}

	return ways[n]
}

func main() {
	var n int
	fmt.Scanln(&n)

	ways := countWays(n)
	fmt.Println(ways)
}
