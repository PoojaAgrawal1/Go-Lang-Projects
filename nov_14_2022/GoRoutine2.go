package main

import (
	"fmt"
	"time"
)

func Person(str string) {
	for i := 0; i < 6; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}

func main() {

	// Calling Goroutine
	go Person("Pooja")

	// Calling normal function
	Person("Aman")
}
