package main

import (
	"fmt"
	"time"
)

// function 1
func Channel1(C1 chan string) {

	time.Sleep(3 * time.Second)
	C1 <- "This is channel 1"
}

// function 2
func Channel2(C2 chan string) {

	time.Sleep(9 * time.Second)
	C2 <- "This is channel 2"
}

// main function
func main() {

	// Declaring channels
	C1 := make(chan string)
	C2 := make(chan string)

	// calling function 1 and
	// function 2 in goroutine

	for i := 0; i < 10; i++ {
		go Channel1(C1)
		go Channel2(C2)
		select {

		// case 1 for portal 1
		case op1 := <-C1:
			fmt.Println(op1)

		// case 2 for portal 2
		case op2 := <-C2:
			fmt.Println(op2)
		}
	}

}
