package main

import "fmt"

func main() {
	var input int
	fmt.Scanln(&input)
	var evenList []int
	for i := 0; i < input; i++ {
		var number int
		fmt.Scanln(&number)
		if number%2 == 0 {
			evenList = append(evenList, number)
		}
	}
	fmt.Println(evenList)
}
