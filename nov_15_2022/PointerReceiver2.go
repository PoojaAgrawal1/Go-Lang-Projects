package main

import "fmt"

type Person struct {
	name     string
	phno     int
	location string
}

func (p *Person) set_per() { //p ==> &per
	fmt.Scan(&p.name)
	fmt.Scan(&p.phno)
	fmt.Scan(&p.location)
} //Here, p will store values at per location since it contains address of per

func (p Person) get_per() { //p ==> per
	fmt.Println("------------------ Output of per ----------------------------")
	fmt.Println("Name : ", p.name)
	fmt.Println("Phone Number : ", p.phno)
	fmt.Println("Location : ", p.location)
}

func main() {
	per := Person{}
	per.set_per()
	per.get_per()
}
