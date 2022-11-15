package main

import "fmt"

type Person struct {
	name     string
	phno     int
	location string
}

func (p Person) set_per() { //p ==> per
	fmt.Scan(&p.name)
	fmt.Scan(&p.phno)
	fmt.Scan(&p.location)
} //Here, p is a local variable. And it is out of scope for per

func (p Person) get_per() { //p ==> per
	fmt.Println("Name : ", p.name)
	fmt.Println("Phone Number : ", p.phno)
	fmt.Println("Location : ", p.location)
}

func main() {
	per := Person{}
	per.set_per()
	per.get_per()
}
