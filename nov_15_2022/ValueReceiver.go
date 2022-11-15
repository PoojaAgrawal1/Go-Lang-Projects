package main

import "fmt"

type Person struct {
	name     string
	phno     int
	location string
}

//This has receiver hence it is a method
func (p Person) get_per() { //p ==> per
	fmt.Println("Name : ", p.name)
	fmt.Println("Phone Number : ", p.phno)
	fmt.Println("Location : ", p.location)
}

func main() {
	per := Person{"Samuel", 9899998933, "Chennai"}
	per.get_per()
}
