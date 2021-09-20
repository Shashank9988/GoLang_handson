package main

import "fmt"

func main() {
	
	emails := map[string]string{"Shashank": "tiwarishashank325@gmail.com", "Rahuk": "rahul@gmail.com"}

	emails["James"] = "james@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["James"])

	// Delete from map
	delete(emails, "James")
	fmt.Println(emails)

	// MAP using make
	mp := make(map[string]int)
	


}
