package main

import (
	"fmt"
	"time"
)

func Channel1(C1 chan string) {

	time.Sleep(3 * time.Second)
	C1 <- "This is channel 1"
}

func Channel2(C2 chan string) {

	time.Sleep(9 * time.Second)
	C2 <- "This is channel 2"
}

func main() {

	C1 := make(chan string)
	C2 := make(chan string)

	for i := 0; i < 10; i++ {
		go Channel1(C1)
		go Channel2(C2)
		select {

		// case 1 for Channel 1
		case op1 := <-C1:
			fmt.Println(op1)

		// case 2 for Channel 2
		case op2 := <-C2:
			fmt.Println(op2)
		}
	}

}
