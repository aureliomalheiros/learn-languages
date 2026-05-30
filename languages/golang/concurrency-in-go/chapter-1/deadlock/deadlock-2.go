package main

func main() {
	ch := make(chan int)

	ch <- 42 // This will cause a deadlock because there is no goroutine receiving from the channel.

	// To fix this, we can start a goroutine that receives from the channel before sending a value to it.
	// go func() {
	//     fmt.Println(<-ch) // This will receive the value sent to the channel and print it.
	// }()
}
