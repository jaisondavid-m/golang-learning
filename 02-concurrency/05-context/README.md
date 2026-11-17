# Context in Go

## Overview

This lesson introduces Go's `context` package for managing cancellation, timeouts, and deadlines in concurrent programs.

## Examples

- `eg-1`: Creates a basic context with `context.Background()`.
- `eg-2`: Cancels a context using `context.WithCancel()`.
- `eg-3`: Stops a worker goroutine when its context is cancelled.
- `eg-4`: Cancels an operation when it exceeds a timeout.
- `eg-5`: Completes an operation before its timeout.
- `eg-6`: Cancels an operation after a deadline.

## Important Points

- Use `context.Background()` as a starting point for a context.
- Use `context.WithCancel()` to manually cancel an operation.
- Use `context.WithTimeout()` to limit how long an operation can run.
- Use `context.WithDeadline()` to stop an operation at a specific time.
- Listen to `ctx.Done()` to respond to cancellation.
- Always call the returned `cancel` function when it is no longer needed.

## Key Takeaway

**Contexts help control the lifetime of operations and coordinate cancellation across goroutines.**
