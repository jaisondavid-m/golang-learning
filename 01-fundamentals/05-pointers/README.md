# 05 - Pointers in Go

## Overview

Learned how pointers work in Go, how to access memory addresses, modify original variables, and pass pointers to functions.

## Topics Covered

- Pointer declaration (`*int`)
- Getting memory addresses (`&`)
- Dereferencing pointers (`*`)
- Modifying variables using pointers
- Passing pointers to functions
- Nil pointers and nil checks

## Practical Examples

### 1. Update Age

Created a function that modifies the original age variable using a pointer.

### 2. Update Bank Balance

Created a function to add money to an existing balance using a pointer.

```go
func updateBalance(balance *int, amount int) {
    *balance += amount
}
```

### 3. Nil Pointer Check

Learned to check whether a pointer is `nil` before dereferencing it.

## Key Learnings

- A pointer stores the memory address of a variable.
- `&` gets the address of a variable.
- `*` accesses or modifies the value at that address.
- Pointers allow functions to modify the caller's original data.
- Dereferencing a nil pointer causes a runtime panic.
- Go supports pointers but not pointer arithmetic.

