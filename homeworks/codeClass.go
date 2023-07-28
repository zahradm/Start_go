package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var kInput int
	fmt.Scanln(&kInput)

	var builder strings.Builder
	for i := 0; i < 40001; i++ {
		builder.WriteString(strconv.Itoa(i))
	}

	finalString := builder.String()
	fmt.Println(string(finalString[kInput]))
}
