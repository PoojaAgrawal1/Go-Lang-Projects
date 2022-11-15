package main

import "fmt"

func display(str string) {
	for i := 0; i < 6; i++ {
		fmt.Println(str)
	}
}

func main() {

	// Calling Goroutine
	go display("Hello!")

	// Calling normal function
	display("Heyyya")
}
