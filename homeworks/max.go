package main

import (
	"fmt"
)

func main() {
	var k int
	fmt.Scanf("%d", &k)
	inputs := make([]int, k)
	for i := 0; i < k; i++ {
		fmt.Scanf("%d", &inputs[i])
	}
	max := inputs[0]
	for i := 1; i < len(inputs); i++ {
		if inputs[i] > max {
			max = inputs[i]
		}
	}
	fmt.Println(max)

}
