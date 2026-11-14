# Goroutines in Go

## Overview

This lesson introduces concurrency in Go using goroutines. Goroutines are lightweight functions managed by the Go runtime that allow multiple tasks to execute independently.

## Concepts Covered

- Understanding concurrency
- Creating goroutines using the `go` keyword
- Main goroutine and program execution
- Understanding goroutine lifecycle
- Using `time.Sleep()` for basic experimentation
- Understanding `sync.WaitGroup`
- `WaitGroup.Add()`, `WaitGroup.Done()`, and `WaitGroup.Wait()`
- Understanding `defer`
- Running multiple goroutines concurrently


## Important Points

### 1. Goroutines

- Use the `go` keyword to start a goroutine.
- Goroutines are managed by the Go runtime.
- The main function does not automatically wait for other goroutines.
- When `main()` returns, the entire program exits.

### 2. WaitGroup

A WaitGroup helps coordinate multiple goroutines.

| Method | Purpose |
|---|---|
| `Add(n)` | Increases the WaitGroup counter |
| `Done()` | Decreases the counter by one |
| `Wait()` | Blocks until the counter reaches zero |

### 3. Defer

- `defer` executes a function when the surrounding function returns.
- `defer wg.Done()` ensures the WaitGroup is notified when a worker finishes.
- Deferred calls execute in last-in, first-out order.

### 4. Concurrent Execution

- Multiple goroutines can make progress independently.
- Execution order is not guaranteed.
- Use synchronization when the main goroutine needs to wait for workers.

## Key Takeaway

**Goroutines enable concurrent execution, while WaitGroups help coordinate and wait for multiple goroutines to finish.**
