# Deadlocks in Go

## Overview

This lesson demonstrates deadlocks in Go. A deadlock occurs when a goroutine waits indefinitely for an operation that can never complete.

## Examples

- `eg-1`: An unbuffered channel send blocks because no goroutine is ready to receive the value.
- `eg-2`: A mutex is locked twice by the same goroutine before the first lock is released.

## Important Points

- An unbuffered channel send waits for a receiver.
- A `sync.Mutex` must be unlocked before it can be locked again.
- The Go runtime detects deadlocks when all goroutines are blocked.
- Use a receiver goroutine, a buffered channel, or correct lock/unlock coordination to avoid deadlocks.

## Key Takeaway

**Careful coordination of channels and mutexes is essential to prevent goroutines from waiting forever.**
