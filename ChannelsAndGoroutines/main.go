package main
import "fmt"

func main() {

	finish := make(chan bool, 2) // Unbuffered Channel
	//	buffered :=make(chan bool, 2) //Buffered channel of size 2
	go func() {
		fmt.Println("Hi")
		finish <- true
	}()
	fmt.Println("There")
	<-finish

}
