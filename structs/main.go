package main

import (
	"fmt"
	"strconv"
)

// Define person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value reciever)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer reciever)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer reciever)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {

	person1 := Person{firstName: "Shashank", lastName: "Tiwari", city: "Shillong", gender: "m", age: 23}
	// Alternative
	person2 := Person{"Bob", "Johnson", "New York", "m", 30}
	person1.hasBirthday()

	fmt.Println(person2.greet())
}
