## About the Code

This program demonstrates basic concepts of goroutines in Go, including:

- Creation of a goroutine
- Memory analysis before/after
- Runtime information
- An example of what **NOT to do** in production code

## Problematic Function: getGID()

```go
func getGID() uint64 {
    buf := make([]byte, 64)
    buf = buf[:runtime.Stack(buf, false)]
    var id uint64
    fmt.Sscanf(string(buf), "goroutine %d ", &id)
    return id
}
```

### Why is this function problematic?

1. **Undocumented**: Not part of Go's public API
2. **Fragile**: Depends on the string format returned by runtime.Stack()
3. **Can break**: Implementation may change in future Go versions
4. **Goes against language philosophy**: Go deliberately doesn't expose goroutine IDs by design

## Recommended Alternatives

To identify goroutines, use:

1. **Contexts with values**:
```go
ctx := context.WithValue(context.Background(), "goroutine-name", "worker-1")
```

2. **Custom data structures**:
```go
type Worker struct {
    Name string
    ID   int
}
```

3. **Channels for communication** (Go standard pattern):
```go
results := make(chan Result, 10)
go worker(results, "worker-1")
```

## What This Code Demonstrates (valid usage)

| Feature | Valid Usage | Example in Code |
|---------|-----------|-------------------|
| runtime.NumGoroutine() | YES Documented | Counts active goroutines |
| runtime.ReadMemStats() | YES Documented | Memory usage analysis |
| runtime.GOMAXPROCS() | YES Documented | Number of system threads |
| runtime.Version() | YES Documented | Go version in use |
| runtime.Stack() | DEBUGGING ONLY | Usage in getGID() is a hack |

## Running the Code

```bash
# Run normally
go run goroutine_study.go

# For detailed traces
GODEBUG=schedtrace=1000 go run goroutine_study.go
```

## Expected Output

```
[Main] Goroutine ID: 1
Number of goroutines: 1
Number of threads: 8
Memory Alloc before: 91504 bytes
Starting goroutine...
[Goroutine ID: 6]
[Goroutine] Number of CPUs: 8
[Goroutine] Pointer size: 8 bytes
[Goroutine] Go version: go1.21.0
[Goroutine] OS/Arch: linux/amd64
[Goroutine] Number of Goroutines: 2
Goroutine completed.
After goroutine execution:
Number of goroutines: 1
Memory Alloc after: 95280 bytes
Memory Alloc difference: 3776 bytes
Number of goroutines before exit: 1
Main function completed.
```

## Study Objectives

1. Understand that goroutines are lightweight (~2KB initial)
2. Verify that goroutines are managed by the Go runtime
3. Understand that the main function is also a goroutine
4. Analyze memory consumption during concurrency

## Official References

### Recommended Documentation:

- **Effective Go - Concurrency**: https://go.dev/doc/effective_go#concurrency
- **Go Concurrency Patterns**: https://go.dev/blog/concurrency-patterns
- **pkg runtime**: https://pkg.go.dev/runtime
- **FAQ: Why no goroutine ID?**: https://go.dev/doc/faq#no_goroutine_id

### Important Articles:

- Go Proverbs: "The bigger the interface, the weaker the abstraction": https://go-proverbs.github.io/
- Dave Cheney: "Why you shouldn't use goroutine IDs": https://dave.cheney.net/2015/11/18/goroutine-ids-do-not-exist
- Rob Pike: "Concurrency is not Parallelism": https://www.youtube.com/watch?v=oV9rvDllKEg

### Tools for Real Debugging:

```go
import _ "net/http/pprof"  // For profiling
import "runtime/debug"     // For debug statistics
```

## Suggested Improvements

Replace the getGID() function with:

```go
// Safe version to identify goroutines
func logGoroutineInfo(name string) {
    fmt.Printf("[%s] Goroutine info:\n", name)
    fmt.Printf("  CPUs: %d\n", runtime.NumCPU())
    fmt.Printf("  Active goroutines: %d\n", runtime.NumGoroutine())
    // Add semantic identifiers
}
```
