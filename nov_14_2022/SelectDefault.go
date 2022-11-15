package main

import "fmt"

func main() {

	mychannel := make(chan int)
	select {
	case <-mychannel:

	default:
		//Since Channel is not yet assigned
		fmt.Println("Not found")
	}
}
