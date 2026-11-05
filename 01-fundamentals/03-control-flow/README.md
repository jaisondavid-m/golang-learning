# 03 - Control Flow in Go

## Overview

This lesson covers control-flow statements in Go. Control flow determines how a program makes decisions, executes different blocks of code, and repeats operations.

## Topics Covered

### 1. If / Else

Used to execute code based on conditions.

```go
age := 19

if age >= 18 {
    fmt.Println("Eligible to vote")
} else {
    fmt.Println("Not eligible to vote")
}
```

### 2. Switch Statement

Used to execute different blocks based on matching values.

```go
status := "SUCCESS"

switch status {
case "PENDING", "PROCESSING":
    fmt.Println("Order is being processed")
case "SUCCESS":
    fmt.Println("Order completed")
case "FAILED":
    fmt.Println("Order failed")
default:
    fmt.Println("Unknown status")
}
```

### 3. Expressionless Switch

A switch without an expression can evaluate multiple conditions.

```go
marks := 85

switch {
case marks >= 90:
    fmt.Println("Grade A")
case marks >= 75:
    fmt.Println("Grade B")
default:
    fmt.Println("Grade C or below")
}
```

### 4. For Loops

Go uses `for` as its only looping keyword.

**Traditional for loop**

```go
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}
```

**Condition-only loop**

```go
i := 1

for i <= 5 {
    fmt.Println(i)
    i++
}
```

**Infinite loop**

```go
for {
    fmt.Println("Running...")
}
```

### 5. Range

Used to iterate over arrays, slices, maps, strings, and other iterable data.

```go
numbers := []int{10, 20, 30, 40}

for index, value := range numbers {
    fmt.Println(index, value)
}
```

Use `_` when you don't need a returned value:

```go
for _, value := range numbers {
    fmt.Println(value)
}
```

### 6. Break and Continue

- `break`: Exits the nearest loop or switch.
- `continue`: Skips the current loop iteration and moves to the next one.

## Practical Project: Go ATM

Built a simple command-line ATM application using:

- Initial balance of ₹5,000
- Balance checking
- Money deposits
- Money withdrawals
- Invalid amount validation
- Insufficient balance validation
- Menu-driven interaction using `for` and `switch`

### Concepts Applied

| Concept | Application |
|---|---|
| `for` | Repeated ATM menu |
| `switch` | Menu selection |
| `if/else` | Amount validation |
| `fmt.Scan()` | User input |
| `break` / `return` | Control program execution |

## Key Learnings

- Go uses `for` for all looping operations.
- Parentheses around `if` and `switch` conditions are not required.
- Go switch statements do not automatically fall through.
- `range` makes collection iteration easier.
- `break` and `continue` control loop execution.

## Next Lesson