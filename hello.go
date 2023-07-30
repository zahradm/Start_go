package main

import "fmt"

type MyStruct struct {
	printy string
}

func main() {
	fmt.Println("Hello!")
	output := MyStruct{}
	fmt.Printf("%+v\n", output)

}
