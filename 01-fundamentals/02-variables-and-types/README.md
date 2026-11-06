# 02 - Variables, Constants & Data Types

## Overview

This lesson covers how Go handles variables, constants, data types, type inference, and zero values.

## 1. Constants

Constants are values that cannot be modified after declaration.

```go
const appName = "AppName fixed"
```

Go also supports grouped constant declarations:

```go
const (
    statusPending = "PENDING"
    statusSuccess = "SUCCESS"
    statusFailed  = "FAILED"
)
```

**Key points:**
- Declared using the `const` keyword.
- Cannot be reassigned.
- Useful for fixed application values, status codes, and configuration defaults.

## 2. Variable Declaration

Go provides multiple ways to declare variables.

### Method 1: Explicit Type Declaration

```go
var name string = "jaison"
var age int = 18
var salary float64 = 25000.00
```

The programmer explicitly specifies the variable's data type.

### Method 2: Type Inference

```go
var name = "jaison"
var age = 18
```

Go automatically determines the variable type based on the assigned value.

For example:
- `"jaison"` → string
- `18` → int

### Method 3: Short Variable Declaration

```go
name := "jaison"
```

**Important:**
- Uses `:=` instead of `var`.
- Automatically infers the type.
- Can only be used inside functions.
- Commonly used for local variables.

## 3. Zero Values

When a variable is declared without an initial value, Go automatically assigns its zero value.

Example:

```go
var name string
var age int
var salary float64
var isActive bool
```

| Data Type | Zero Value |
|---|---|
| string | `""` |
| int | `0` |
| float64 | `0` |
| bool | `false` |
| Pointer | `nil` |
| Slice | `nil` |
| Map | `nil` |

Go initializes variables automatically, so they do not contain unpredictable garbage values.

## 4. Unused Variables

Go does not allow declared local variables to remain unused.

Example:

```go
var age int = 18
```

If `age` is never used, Go reports a compilation error.

One way to explicitly discard a value is the blank identifier:

```go
_ = age
```

The blank identifier tells Go that the value is intentionally discarded.

**Note:** It is better to actually use a variable when possible rather than use `_` simply to silence compiler errors.

## 5. Data Types Covered

| Type | Example | Description |
|---|---|---|
| string | `"jaison"` | Text |
| int | `18` | Integer |
| float64 | `25000.00` | Floating-point number |
| bool | `true` | Boolean value |

## 6. Key Takeaways

- `var` declares variables.
- `:=` provides short variable declaration inside functions.
- Go supports automatic type inference.
- `const` declares immutable values.
- Every Go variable receives a zero value if no initial value is provided.
- Go requires local variables to be used.
- The blank identifier `_` discards values intentionally.

## Practical Application

These concepts are fundamental to backend development.

For example, backend applications use variables and constants to manage:

- Application configuration.
- User information.
- Request data.
- Authentication status.
- API response statuses.
- Database values.
