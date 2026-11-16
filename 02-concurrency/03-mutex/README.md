# Mutexes in Go

## Overview

This lesson demonstrates how mutexes protect shared data when multiple goroutines access it concurrently.

## Examples

- `eg-1`: Multiple goroutines update a shared balance without locking the mutex.
- `eg-2`: Multiple goroutines safely increment a shared counter using `sync.Mutex`.

## Important Points

- A mutex provides exclusive access to shared data.
- Call `Lock()` before accessing shared data.
- Call `Unlock()` after the operation is complete.
- Use `defer` or careful control flow to ensure a mutex is released.
- A `sync.WaitGroup` waits for all worker goroutines to finish.

## Key Takeaway

**Use mutexes to prevent data races when multiple goroutines read or modify shared data.**
