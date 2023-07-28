package main

import (
	"fmt"
	"strconv"
)

func maghloob() {
	var input int
	fmt.Scanln(&input)
	strInput := strconv.Itoa(input)
	strFinal := ""
	for i := len(strInput) - 1; i > -1; i-- {
		strFinal = strFinal + string(strInput[i])

	}
	if strFinal == strInput {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func main() {
	maghloob()
}
