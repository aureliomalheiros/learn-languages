package main

import (
	"fmt"
	"time"
)


func  processarPedido(id int) {
	fmt.Printf("Pedido %d processando...\n", id)
	time.Sleep(2 * time.Second) // Simulate time
	fmt.Printf("Pedido %d processado!\n", id)
}

//With goroutines, we can process multiple requests concurrently. Each request can be handled in its own goroutine, allowing them to run simultaneously without blocking each other. This means that while one request is being processed, others can also be processed at the same time, improving the overall responsiveness of the application.

func main() {
	go processarPedido(1)
	go processarPedido(2)
	go processarPedido(3)

	time.Sleep(3 * time.Second) // Wait for all goroutines to finish
}

//Without goroutines, the code executes sequentially, meaning that each request is processed one after the other. This can lead to delays if one request takes a long time to process, as subsequent requests will have to wait until the current one finishes.

// func main() {
// 	processarPedido(1)
// 	processarPedido(2)
// 	processarPedido(3)
// }
