package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	candidates := make([]int, n)
	for i := 0; i < n; i++ {
		candidates[i] = i + 1
	}
	eliminatedIndex := 1
	for len(candidates) > 1 {
		candidates = append(candidates[:eliminatedIndex], candidates[eliminatedIndex+1:]...)
		eliminatedIndex = (eliminatedIndex + 1) % len(candidates)
	}
	fmt.Println(candidates[0])
}
