package main

import "fmt"

func main() {

	var age int32 = 23
	const isCool = true
	var size float32 = 2.3

	name, email := "Shashank", "shashank.tiwari@cilalabs.com"
	fmt.Println(name, age, isCool, email)
	fmt.Printf("%T\n", size)
}
