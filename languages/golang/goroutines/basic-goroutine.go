package main

import (
	"fmt"
	"time"
	"runtime"
	"unsafe"
)


func myFunction() {
	fmt.Printf("[Goroutine ID: %d]\n", getGID())
	fmt.Printf("[Goroutine] Number of CPUs: %d\n", runtime.NumCPU())
	fmt.Printf("[Goroutine] Pointer size: %d bytes\n", unsafe.Sizeof(uintptr(0)))
	fmt.Printf("[Goroutine] Go version: %s\n", runtime.Version())
	fmt.Printf("[Goroutine] OS/Arch: %s/%s\n", runtime.GOOS, runtime.GOARCH)
	fmt.Printf("[Goroutine] Number of Goroutines: %d\n", runtime.NumGoroutine())
	fmt.Println("Goroutine completed.")
}

func getGID() uint64 {

	buf := make([]byte, 64)
	buf  = buf[:runtime.Stack(buf, false)]
	var id uint64
	fmt.Sscanf(string(buf), "goroutine %d ", &id)
	return id

}

func main() {
	fmt.Printf("[Main] Goroutine ID: %d\n", getGID())
    fmt.Printf("Número de goroutines: %d\n", runtime.NumGoroutine())
    fmt.Printf("Número de threads: %d\n", runtime.GOMAXPROCS(0))

	var memBefore runtime.MemStats
	runtime.ReadMemStats(&memBefore)
	fmt.Printf("Memory Alloc before: %d bytes\n", memBefore.Alloc)

	fmt.Println("Starting goroutine...")
	go myFunction()

	time.Sleep(20 * time.Second)

	fmt.Println("After goroutine execution:")
	fmt.Printf("Number of goroutines: %d\n", runtime.NumGoroutine())

	var memAfter runtime.MemStats
	runtime.ReadMemStats(&memAfter)
	fmt.Printf("Memory Alloc after: %d bytes\n", memAfter.Alloc)
	fmt.Printf("Memory Alloc difference: %d bytes\n", memAfter.Alloc - memBefore.Alloc)

	time.Sleep(5 * time.Second)
	
	fmt.Printf("Number of goroutines before exit: %d\n", runtime.NumGoroutine())
	fmt.Println("Main function completed.")

}
