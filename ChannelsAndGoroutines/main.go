package main
import "fmt"

func  add (a int, b int) int {
	return a+b
}

func main() {

	// finish := make(chan bool) // Unbuffered Channel
	// //	buffered :=make(chan bool, 2) //Buffered channel of size 2
	// go func() {  // goroutine
	// 	fmt.Println("Hi")
	// 	finish <- true
	// }()
	// fmt.Println("There")
	// <-finish 
	fmt.Println(add(2,4))

}
