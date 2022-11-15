package main

import "fmt"

func main() {
	var a int = 10
	var b *int
	b = &a
	fmt.Println("The value of Pointer variable b ")
	fmt.Println(b)
	fmt.Println("----------------")
	fmt.Println("The value of another variable whose address b contains")
	fmt.Println(*b)
}
